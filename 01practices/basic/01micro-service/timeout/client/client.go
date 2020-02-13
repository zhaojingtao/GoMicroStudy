package main

import (
	proto "GoMicroStudy/01practices/basic/01micro-service/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"time"
)

func main() {
	cli := client.NewClient(
		client.RequestTimeout(5 * time.Second),
	)
	// 定义服务，可以传入相关的其他参数
	service := micro.NewService(
		micro.Name("timeout.client"),
		micro.Client(cli),
	)
	service.Init()
	greeter := proto.NewGreeterService("timeout.service", service.Client())
	// 调用greeter服务
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "MICRO中国"})
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印响应结果
	fmt.Println(rsp.Greeting)
}
