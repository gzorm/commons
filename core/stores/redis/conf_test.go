package redis

import (
	"testing"

	"github.com/gzorm/common/core/stringx"
	"github.com/stretchr/testify/assert"
)

func TestRedisConf(t *testing.T) {
	tests := []struct {
		name string
		RedisConf
		ok bool
	}{
		{
			name: "missing host",
			RedisConf: RedisConf{
				Host: "localhost:6379",
				Type: NodeType,
				Pass: "",
				Db:   1,
			},
			ok: false,
		},
		{
			name: "missing type",
			RedisConf: RedisConf{
				Host: "localhost:6379",
				Type: "",
				Pass: "",
				Db:   1,
			},
			ok: false,
		},
		{
			name: "ok",
			RedisConf: RedisConf{
				Host: "localhost:6379",
				Type: NodeType,
				Pass: "",
				Db:   1,
			},
			ok: true,
		},
		{
			name: "ok",
			RedisConf: RedisConf{
				Host: "localhost:6379",
				Type: ClusterType,
				Pass: "",
				Db:   1,
				Tls:  true,
			},
			ok: true,
		},
	}

	for _, test := range tests {
		t.Run(stringx.RandId(), func(t *testing.T) {
			if test.ok {
				assert.Nil(t, test.RedisConf.Validate())
				_, err := NewRedis(test.RedisConf)
				assert.NotNil(t, err)
			} else {
				assert.NotNil(t, test.RedisConf.Validate())
			}
		})
	}
}

func TestRedisKeyConf(t *testing.T) {
	tests := []struct {
		name string
		RedisKeyConf
		ok bool
	}{
		{
			name: "missing host",
			RedisKeyConf: RedisKeyConf{
				RedisConf: RedisConf{
					Host: "localhost:6379",
					Type: NodeType,
					Pass: "",
					Db:   1,
				},
				Key: "foo",
			},
			ok: false,
		},
		{
			name: "missing key",
			RedisKeyConf: RedisKeyConf{
				RedisConf: RedisConf{
					Host: "localhost:6379",
					Type: NodeType,
					Pass: "",
					Db:   1,
				},
				Key: "",
			},
			ok: false,
		},
		{
			name: "ok",
			RedisKeyConf: RedisKeyConf{
				RedisConf: RedisConf{
					Host: "localhost:6379",
					Type: NodeType,
					Pass: "",
					Db:   1,
				},
				Key: "foo",
			},
			ok: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.ok {
				assert.Nil(t, test.RedisKeyConf.Validate())
			} else {
				assert.NotNil(t, test.RedisKeyConf.Validate())
			}
		})
	}
}
