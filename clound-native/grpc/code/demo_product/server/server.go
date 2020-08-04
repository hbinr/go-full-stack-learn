package main

import (
	"go-full-stack-learn/clound-native/grpc/code/demo_product/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 1.创建一个 rpc server
	rpcServer := grpc.NewServer()

	// 2.注册服务，调用 proto 生成的代码中的方法. 参数：server， 创建的服务
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService))

	// 3.创建TCP连接 监听套接字
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("创建TCP连接异常,err:", err)
	}
	// 4.使用grpc，创建连接
	rpcServer.Serve(lis)
}
