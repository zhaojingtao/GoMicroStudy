package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/micro/go-micro/v2"
	"io"
	"io/ioutil"
	"os"

	proto "GoMicroStudy/01practices/basic/01micro-service/stream/file/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
)

type File struct{}

// 文件接收方法
func (g *File) File(ctx context.Context, file proto.File_FileStream) error {
	//将接受到的内容储存到临时文件中
	temp, err := ioutil.TempFile("", "micro")
	if err != nil {
		return errors.InternalServerError("file.service", err.Error())
	}
	for {
		b, err := file.Recv()
		if err != nil {
			return errors.InternalServerError("file.service", err.Error())
		}
		if b.Len == -1 {
			//完成标志
			break
		}
		if _, err := temp.Write(b.Byte); err != nil {
			return errors.InternalServerError("file.service", err.Error())
		}
	}
	println(temp.Name())
	// 发送文件信息
	return file.SendMsg(&proto.FileMsg{
		FileName: temp.Name(),
	})
}

// 文件处理方法
func (g *File) DealFile(ctx context.Context, req *proto.DealFileRequest, rsp *proto.DealFileRespond) error {
	// 通过文件名获取到文件内容
	// 计算文件md5
	hash := md5.New()
	file, err := os.OpenFile(req.FileName, os.O_RDONLY, 0755)
	if err != nil {
		return errors.InternalServerError("file.service", err.Error())
	}
	_, _ = io.Copy(hash, file)
	MD5Str := hex.EncodeToString(hash.Sum(nil))
	// 加上param
	hash2 := md5.New()
	hash2.Write([]byte(MD5Str + req.Param))
	rsp.Md5 = hex.EncodeToString(hash2.Sum(nil))
	println(req.FileName + "|" + rsp.Md5)
	return nil
}

func main() {
	// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
	service := micro.NewService(
		micro.Name("file.service"),
		micro.Version("latest"),
	)
	service.Init()

	// 注册服务
	_ = proto.RegisterFileHandler(service.Server(), new(File))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
