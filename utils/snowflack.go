package utils

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gzorm/common/core/logx"
	"github.com/gzorm/common/core/stores/redis"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type Snowflake interface {
	GetId() int64
}

func NewSnowflake(redis *redis.Redis, opts ...snowflakeOpt) Snowflake {
	s := &snowflake{}
	for _, so := range opts {
		so(s)
	}
	workId, err := s.getWorkId(redis)
	if err != nil {
		panic(err)
	}
	s.workerId = workId
	return s
}

type snowflakeOpt func(s *snowflake)

func WithKey(key string) snowflakeOpt {
	return func(s *snowflake) {
		s.key = key
	}
}

const (
	WorkIdKeyFmt = "zero-core:common" + "%s" + ":work-id"
)

func (s *snowflake) getWorkId(redis *redis.Redis) (int64, error) {
	var r int64 = 0
	var err error
	WorkIdKey := fmt.Sprintf(WorkIdKeyFmt, s.key)
	for i := 0; i < int(workerMax); i++ {
		// 递增redis WorkId
		r, err = redis.Incr(WorkIdKey)
		if err != nil {
			return 0, err
		}
		var b bool
		// 判断当前work+1 %1024 后是否被使用过
		b, err = redis.Exists(fmt.Sprintf("%s:%d", WorkIdKey, r&workerMax))
		if err != nil {
			return 0, err
		}
		if !b {
			break
			// 没有被使用过说明机器位没有满
		}
		if int64(i) == (workerMax - 1) {
			// 说明已经部署了1024 台机器
			return 0, errors.New("not enough space work")
		}
	}
	// 对当前workId 设置过期key
	_, err = redis.SetnxEx(fmt.Sprintf("%s:%d", WorkIdKey, r&workerMax), time.Now().Format(time.DateTime), 120)
	if err != nil {
		return 0, err
	}

	s.workerId = r & workerMax
	go func() {
		ticker := time.NewTicker(time.Second * 110)
		defer ticker.Stop()
		for {
			<-ticker.C
			//logx.Infof("expire key【%s】time:%v", fmt.Sprintf("%s:%d", WorkIdKey, r&workerMax), time.Now())
			err = redis.Expire(fmt.Sprintf("%s:%d", WorkIdKey, r&workerMax), 120)
			if err != nil {
				logx.Error(err)
			}
		}
	}()
	return r & workerMax, nil
}

type snowflake struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
	key       string
}

func (w *snowflake) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.timestamp != 0 {
		w.number++
		if w.number > numberMax {
			//for now <= w.timestamp {
			//	now = time.Now().UnixNano() / 1e6
			//}
			// number 需要重置为0 否则work id 会递增导致重启时候 work id 重复可能会有重复的key
			w.number = 0
			// 解决时钟回拨问题
			w.timestamp++
		}
	} else {
		w.number = 0
		now := time.Now().UnixNano() / 1e6
		w.timestamp = now
	}
	ID := int64((w.timestamp-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}
