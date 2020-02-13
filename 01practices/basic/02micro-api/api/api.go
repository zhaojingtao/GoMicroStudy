package main

import (
	proto "GoMicroStudy/01practices/basic/02micro-api/api/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
	"strings"
)

type Example struct {
}

func (e Example) Call(ctx context.Context, request *api.Request, response *api.Response) error {
	log.Log("Example.Call接口收到请求")
	name, ok := request.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "参数不正确")
	}
	// 打印请求头
	for k, v := range request.Header {
		log.Log("请求头信息，", k, " : ", v)
	}
	response.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": "我们已经收到你的请求，" + strings.Join(name.Values, " "),
	})
	// 设置返回值
	response.Body = string(b)
	return nil
}

type Foo struct {
}

func (f Foo) Bar(ctx context.Context, request *api.Request, response *api.Response) error {
	log.Logf("Foo.Bar接口收到请求")
	if request.Method != "POST" {
		return errors.BadRequest("go.micro.api.example", "require post")
	}
	ct, ok := request.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "need content-type")
	}
	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.example", "expect application/json")
	}
	var body map[string]interface{}
	json.Unmarshal([]byte(request.Body), &body)
	// 设置返回值
	response.Body = "收到消息：" + string([]byte(request.Body))
	return nil
}

func main() {
	var err error

	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)
	service.Init()

	// 注册 example handler
	err = proto.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 注册 foo handler
	err = proto.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = service.Run(); err != nil {
		fmt.Println(err)
		return
	}

}
