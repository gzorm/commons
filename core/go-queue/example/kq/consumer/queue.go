package main

import (
	"fmt"

	"github.com/gzorm/commons/core/conf"
	"github.com/gzorm/commons/core/go-queue/kq"
)

func main() {
	var c kq.KqConf
	conf.MustLoad("config.yaml", &c)

	q := kq.MustNewQueue(c, kq.WithHandle(func(k, v string) error {
		fmt.Printf("=> %s\n", v)
		return nil
	}))
	defer q.Stop()
	q.Start()
}
