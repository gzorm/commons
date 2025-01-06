package main

import (
	"flag"
	"fmt"

	"github.com/gzorm/common/core/conf"
	"github.com/gzorm/common/core/go-queue/example/rabbitmq/listener/config"
	"github.com/gzorm/common/core/go-queue/rabbitmq"
	"github.com/gzorm/common/core/service"
)

var configFile = flag.String("f", "listener.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	listener := rabbitmq.MustNewListener(c.ListenerConf, Handler{})
	serviceGroup := service.NewServiceGroup()
	serviceGroup.Add(listener)
	defer serviceGroup.Stop()
	serviceGroup.Start()
}

type Handler struct {
}

func (h Handler) Consume(message string) error {
	fmt.Printf("listener %s\n", message)
	return nil
}
