package config

import "github.com/gzorm/common/core/go-queue/rabbitmq"

type Config struct {
	ListenerConf rabbitmq.RabbitListenerConf
}
