package main

import (
	greeter "GoMicroStudy/01practices/basic/01micro-service/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type Greeter struct {
}

func (g Greeter) Hello(ctx context.Context, request *greeter.HelloRequest, response *greeter.HelloResponse) error {
	response.Greeting = "你好呀," + request.Name
	return nil
}

func main() {
	var (
		err error
	)
	// 创建函数请求响应
	fnc := micro.NewFunction(
		micro.Name("greeter"),
	)
	fnc.Init()
	err = fnc.Handle(new(Greeter))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fnc.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
