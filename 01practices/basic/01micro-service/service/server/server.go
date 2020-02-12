package main

import (
	greeter "GoMicroStudy/01practices/basic/01micro-service/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
)

type Greeter struct {
}

func (g Greeter) Hello(ctx context.Context, request *greeter.HelloRequest, response *greeter.HelloResponse) error {
	response.Greeting = "你好呀," + request.Name
	return nil
}

func main() {
	var err error
	register := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"106.12.118.76:2379"}
	})
	service := micro.NewService(
		// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
		micro.Name("com.service.greeter"),
		micro.Version("1.0.0"),
		micro.Registry(register),
	)
	// 初始化服务
	service.Init()
	// 注册服务到相关方
	err = greeter.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = service.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
