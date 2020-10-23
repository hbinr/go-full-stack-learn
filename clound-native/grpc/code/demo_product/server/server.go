package main

import (
	"fmt"
	"log"
	"net"

	"hb.study/clound-native/grpc/code/demo_product/service"
	"hb.study/clound-native/grpc/code/demo_product/util"

	"google.golang.org/grpc"
)

func main() {
	// 1.创建一个 rpc server，通过CA和SSL证书来进行TLS验证
	rpcServer := grpc.NewServer(grpc.Creds(util.GetServerCert()))

	// 2.注册服务，调用 proto 生成的代码中的方法.
	service.RegisterProdServiceServer(rpcServer, new(service.ProdService)) // 商品服务

	// 3.创建TCP连接 监听套接字
	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		log.Fatal("创建TCP连接异常,err:", err)
	}
	// 4.使用grpc，创建连接
	rpcServer.Serve(lis)

	fmt.Println("服务端启动成功！")
}
