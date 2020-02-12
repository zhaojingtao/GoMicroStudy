package main

import (
	"github.com/micro/micro/cmd"

	// go-micro plugins
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	_ "github.com/micro/go-plugins/transport/nats"
)

func main() {
	cmd.Init()
}
