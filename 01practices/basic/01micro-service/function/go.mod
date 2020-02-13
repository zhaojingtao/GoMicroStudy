module GoMicroStudy/01practices/basic/01micro-service/function

go 1.13

replace GoMicroStudy/01practices/basic/01micro-service/proto => /opt/goPath/src/GoMicroStudy/01practices/basic/01micro-service/proto

require (
	GoMicroStudy/01practices/basic/01micro-service/proto v0.0.0-00010101000000-000000000000
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.0.0 // indirect
)
