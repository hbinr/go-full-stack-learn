package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"hb.study/clound-native/go-kit/code/app/endpoint"
	"hb.study/clound-native/go-kit/code/app/service"
	"hb.study/clound-native/go-kit/code/app/transport"

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
	r := httprouter.New()
	r.Handler(http.MethodGet, "/get/:id", serverHandle)
	http.ListenAndServe(":8080", r)
}
