package utils

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/gzorm/commons/core/logx"
	"github.com/gzorm/commons/core/stores/redis"
	"github.com/gzorm/commons/ent"
	"strings"
	"sync"
)

type entIntercept struct {
	redis    *redis.Redis
	key      string
	duration int
	lock     sync.Mutex
	enable   bool
}

func WithPrefixKey(key string) cacheFuncOpt {
	return func(intercept *entIntercept) {
		intercept.key = key
	}
}

func WithEnableCache() cacheFuncOpt {
	return func(intercept *entIntercept) {
		intercept.enable = true
	}
}

func WithCacheValidatedPeriod(duration int) cacheFuncOpt {
	return func(intercept *entIntercept) {
		intercept.duration = duration
	}
}

const (
	cacheKey = ":zero-core:table-cache"
)

type cacheFuncOpt func(intercept *entIntercept)

var CacheInterceptor *entIntercept
var once sync.Once

func NewCacheIntercept(redis *redis.Redis, opts ...cacheFuncOpt) *entIntercept {
	once.Do(func() {
		CacheInterceptor = &entIntercept{
			redis:    redis,
			duration: 60 * 10,
		}
		for _, opt := range opts {
			opt(CacheInterceptor)
		}
	})
	return CacheInterceptor
}

func (e *entIntercept) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
		d, ok := ctx.Value("disable").(bool)
		if !CacheInterceptor.enable || (ok && d) {
			return next.Query(ctx, query)
		}
		var k = ""

		if q, ok := query.(*ent.WinConfigQuery); ok {
			sql, args := q.SqlQuery(ctx)
			ret, ok := getFromRedis[*ent.WinConfig](sql, args, ctx, &k)
			if ok {
				return ret, nil
			}
		} else if q, ok := query.(*ent.WinGameSlotQuery); ok {
			sql, args := q.SqlQuery(ctx)
			ret, ok := getFromRedis[*ent.WinGameSlot](sql, args, ctx, &k)
			if ok {
				return ret, nil
			}
		} else if q, ok := query.(*ent.WinPlatListQuery); ok {
			sql, args := q.SqlQuery(ctx)
			ret, ok := getFromRedis[*ent.WinPlatList](sql, args, ctx, &k)
			if ok {
				return ret, nil
			}
		}
		CacheInterceptor.lock.Lock()
		if q, ok := query.(*ent.WinConfigQuery); ok {
			sql, args := q.SqlQuery(ctx)
			ret, ok := getFromRedis[*ent.WinConfig](sql, args, ctx, &k)
			if ok {
				CacheInterceptor.lock.Unlock()
				return ret, nil
			}
		} else if q, ok := query.(*ent.WinGameSlotQuery); ok {
			sql, args := q.SqlQuery(ctx)
			ret, ok := getFromRedis[*ent.WinGameSlot](sql, args, ctx, &k)
			if ok {
				CacheInterceptor.lock.Unlock()
				return ret, nil
			}
		} else if q, ok := query.(*ent.WinPlatListQuery); ok {
			sql, args := q.SqlQuery(ctx)
			ret, ok := getFromRedis[*ent.WinPlatList](sql, args, ctx, &k)
			if ok {
				CacheInterceptor.lock.Unlock()
				return ret, nil
			}
		}
		nq, err := next.Query(ctx, query)
		bs, err := json.Marshal(nq)
		if err != nil {
			logx.Errorf("json marshal %s err %+v\n", k, err)
		}
		_, err = CacheInterceptor.redis.SetnxExCtx(ctx, k, string(bs), CacheInterceptor.duration)
		if err != nil {
			logx.Errorf("redis set %s err %+v\n", k, err)
		}
		CacheInterceptor.lock.Unlock()
		return nq, err
	})
}

type wEnt interface {
	*ent.WinConfig | *ent.WinGameSlot | *ent.WinPlatList
}

func getFromRedis[T wEnt](sql string, args []any, ctx context.Context, k *string) (ret []T, exist bool) {
	for _, arg := range args {
		sql = fmt.Sprintf("%s:%v", sql, arg)
	}
	hash := sha256.New()
	hash.Write([]byte(sql))
	*k = fmt.Sprintf("%s:%s:%x", CacheInterceptor.key, cacheKey, hash.Sum(nil))
	val, err := CacheInterceptor.redis.GetCtx(ctx, *k)
	if err != nil {
		logx.Errorf("redis get %s err %+v\n", *k, err)
	}
	if strings.TrimSpace(val) != "" {
		err = json.Unmarshal([]byte(val), &ret)
		if err != nil {
			logx.Errorf("json unmarshal %s err %+v\n", *k, err)
		} else {
			return ret, true
		}
	}

	return nil, false
}
