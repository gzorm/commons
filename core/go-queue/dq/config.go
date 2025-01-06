package dq

import "github.com/gzorm/commons/core/stores/redis"

type (
	Beanstalk struct {
		Endpoint string
		Tube     string
	}

	DqConf struct {
		Beanstalks []Beanstalk
		Redis      redis.RedisConf
	}
)
