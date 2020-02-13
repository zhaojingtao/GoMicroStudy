package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"io"
	"net/http"

	proto "GoMicroStudy/01practices/basic/01micro-service/stream/file/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

var fileService proto.FileService

var c client.Client

func UploadFile(rsp http.ResponseWriter, req *http.Request) {
	if err := req.ParseMultipartForm(10 << 20); err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	// 取到文件对象
	files, ok := req.MultipartForm.File["file"]
	if !ok {
		rsp.WriteHeader(400)
		_, _ = rsp.Write([]byte("请选择文件上传"))
		return
	}
	// 将文件通过流式传输到srv
	file, err := files[0].Open()
	if err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	// 建立链接
	// 因为这里是用的临时文件储存的方式,如果因为负载均衡算法导致下一次节点切换,另外一个节点是无法通过,文件名来获取到文件数据的
	// 使用这种方法来固定一个节点
	next, _ := c.Options().Selector.Select("file.service")
	node, _ := next()
	stream, err := fileService.File(context.Background(), func(options *client.CallOptions) {
		// 指定节点
		options.Address = []string{node.Address}
	})
	if err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	for {
		buff := make([]byte, 1024*1024) // 缓冲1MB,每次发送1MB的内容,注意不能超过rpc的限制(grpc默认为4MB)
		sendLen, err := file.Read(buff)
		if err != nil {
			if err == io.EOF {
				//全部读取完成,发送一个完成标识,跳出
				err = stream.Send(&proto.FileByte{
					Byte: nil,
					Len:  -1,
				})
				if err != nil {
					rsp.WriteHeader(500)
					_, _ = rsp.Write([]byte(err.Error()))
					return
				}
				break
			}
			rsp.WriteHeader(500)
			_, _ = rsp.Write([]byte(err.Error()))
			return
		}
		err = stream.Send(&proto.FileByte{
			Byte: buff[:sendLen],
			Len:  int64(sendLen),
		})
		if err != nil {
			rsp.WriteHeader(500)
			_, _ = rsp.Write([]byte(err.Error()))
			return
		}
	}
	// 接收收到的消息之后就可以关闭了
	fileMsg := &proto.FileMsg{}
	if err := stream.RecvMsg(fileMsg); err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	_ = stream.Close()

	// 调用文件处理的rpc
	ret, err := fileService.DealFile(context.Background(), &proto.DealFileRequest{
		FileName: fileMsg.FileName,
		Param:    "a param",
	}, func(options *client.CallOptions) {
		// 指定节点
		options.Address = []string{node.Address}
	})
	if err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}

	_, _ = rsp.Write([]byte(ret.Md5))
}

func main() {
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(micro.Name("file.client"))
	service.Init()

	// 创建客户端
	c = service.Client()
	fileService = proto.NewFileService("file.service", c)

	// 一个文件上传的api
	http.HandleFunc("/upload", UploadFile)
	_ = http.ListenAndServe(":8080", nil)
}
