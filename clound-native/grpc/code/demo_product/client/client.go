package main

import (
	"context"
	"fmt"
	"go-full-stack-learn/clound-native/grpc/code/demo_product/service"
	"go-full-stack-learn/clound-native/grpc/code/demo_product/util"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 1.创建连接套接字，端口和服务端一致，并开启 TLS 验证
	conn, err := grpc.Dial(":8092", grpc.WithTransportCredentials(util.GetClientCert()))
	if err != nil {
		log.Fatal("客户端连接服务端异常，err:", err)
	}
	defer conn.Close()

	// 2.创建grpc客户端服务，调用 proto 生成的代码中的方法.
	prodClient := service.NewProdServiceClient(conn)

	// 3.调用远程服务，并处理响应
	prodResp, err := prodClient.GetProdName(context.Background(), &service.ProdRequest{
		ProdID:   40,
		ProdArea: 0,
	})
	if err != nil {
		log.Fatal("GetProdName 获取商品名称异常，err:", err)
	}
	prodList, err := prodClient.GetProdNameList(context.Background(), &service.QueryRequest{PageSize: 5})
	if err != nil {
		log.Fatal("GetProdNameList 获取商品名称列表异常，err:", err)
	}

	prodInfo, err := prodClient.GetProdInfo(context.Background(), &service.ProdRequest{
		ProdID:   40,
		ProdArea: 0,
	})
	if err != nil {
		log.Fatal("GetProdInfo 获取商品信息异常，err:", err)
	}

	fmt.Println("调用 GetProdName 接口成功，商品名称为：", prodResp.ProdName)
	fmt.Println("调用 GetProdNameList 接口成功，商品名称为：", prodList)
	fmt.Println("调用 GetProdInfo 接口成功，商品信息为：", prodInfo)
}
