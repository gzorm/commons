package config

import "github.com/gzorm/commons/core/go-queue/rabbitmq"

type Config struct {
	ListenerConf rabbitmq.RabbitListenerConf
}
