package main

import (
	proto "GoMicroStudy/01practices/basic/01micro-service/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
)

func main() {
	register := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"106.12.118.76:2379"}
	})
	service := micro.NewService(
		micro.Name("com.client.greeter"),
		micro.Version("1.0.0"),
		micro.Registry(register),
	)
	service.Init()
	// 调用服务
	greeter := proto.NewGreeterService("com.service.greeter", service.Client())
	// 调用greeter服务
	rsp, err := greeter.Hello(context.Background(), &proto.HelloRequest{Name: "袁国仁"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
