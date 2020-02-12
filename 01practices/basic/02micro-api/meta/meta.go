package main

import (
	proto "GoMicroStudy/01practices/basic/02micro-api/meta/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/api"
	rapi "github.com/micro/go-micro/api/handler/api"
	"github.com/micro/go-micro/api/handler/rpc"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"log"
)

type Example struct {
}

func (e Example) Call(ctx context.Context, request *proto.CallRequest, response *proto.CallResponse) error {
	log.Print("Meta Example.Call接口收到请求")
	if len(request.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}
	response.Message = "Meta已经收到你的请求，" + request.Name
	return nil
}

type Foo struct {
}

func (f Foo) Bar(ctx context.Context, request *proto.EmptyRequest, response *proto.EmptyResponse) error {
	log.Print("Meta Foo.Bar接口收到请求")
	// noop
	return nil
}

func main() {
	var err error
	register := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"106.12.118.76:2379"}
	})
	service := micro.NewService(
		// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
		micro.Name("com.service.api.example"),
		micro.Version("1.0.0"),
		micro.Registry(register),
	)
	// 初始化服务
	service.Init()
	// 注册Example接口处理器
	err = proto.RegisterExampleHandler(service.Server(), new(Example), api.WithEndpoint(&api.Endpoint{
		Name:    "Example.Call",
		Handler: rpc.Handler,
		Method:  []string{"POST", "GET"},
		Path:    []string{"/example"},
	}))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 注册Foo接口处理器
	err = proto.RegisterFooHandler(service.Server(), new(Foo), api.WithEndpoint(&api.Endpoint{
		Name:    "Foo.Bar",
		Handler: rapi.Handler,
		Method:  []string{"POST", "GET", "DELETE"},
		Path:    []string{"/foo/bar"},
	}))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = service.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
