package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
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

	// HttpRouter(serverHandle)
	HttpChi(serverHandle)
	// HttpMux(serverHandle)
	// HttpGoFrame(serverHandle)
}

func HttpRouter(serverHandle *httptrasport.Server) {
	r := httprouter.New()
	r.Handler(http.MethodGet, "/get/:id", serverHandle)
	http.ListenAndServe(":8080", r)
}

func HttpChi(serverHandle *httptrasport.Server) {
	chi := chi.NewRouter()
	chi.Method(http.MethodGet, "/get/:id", serverHandle)
	http.ListenAndServe(":8080", chi)

}

func HttpMux(serverHandle *httptrasport.Server) {
	mux := mux.NewRouter()
	mux.Handle("/get/:id", serverHandle).Methods("GET")
	http.ListenAndServe(":8080", mux)
}

func HttpGoFrame(serverHandle *httptrasport.Server) {
	// s := g.Server()

	// s.Group("/user").GET("/get", serverHandle)
	// s.Run()

	// http.ListenAndServe(":8080", s)
}

// test: wrk -t12 -c100 -d30s http://localhost:8080/get/1
