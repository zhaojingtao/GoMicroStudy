package main

import (
	greeter "GoMicroStudy/01practices/basic/01micro-service/proto"
	proto "GoMicroStudy/01practices/basic/01micro-service/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("com.client.greeter"),
	)
	service.Init()
	// 调用服务
	greeter := greeter.NewGreeterService("com.service.greeter", service.Client())
	// 调用greeter服务
	rsp, err := greeter.Hello(context.Background(), &proto.HelloRequest{Name: "袁国仁"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
