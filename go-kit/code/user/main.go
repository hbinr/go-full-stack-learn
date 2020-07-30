package main

import (
	"go-full-stack-learn/go-kit/code/user/endpoint"
	"go-full-stack-learn/go-kit/code/user/service"
	"go-full-stack-learn/go-kit/code/user/transport"
	"net/http"

	httptrasport "github.com/go-kit/kit/transport/http"
)

func main() {

	// 创建服务结构体
	userService := service.UserService{}
	// 生成endpoint
	endp := endpoint.GenUserEndpoint(userService)
	// 创建服务，go-kit的 http库
	serverHandle := httptrasport.NewServer(endp, transport.DecodeUserRequest, transport.EncodeUserResponse)
	// 启动服务
	http.ListenAndServe(":8080", serverHandle)
}
