package util

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc/credentials"
)

// GetServerCert 加载服务端证书
func GetServerCert() credentials.TransportCredentials {

	cert, err := tls.LoadX509KeyPair("../cert/server.pem", "../cert/server.key")
	if err != nil {
		log.Fatal("服务端生成cert异常，err:", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../cert/ca.pem")
	if err != nil {
		log.Fatal("服务端读取ca.pem异常，err:", err)
	}
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},        // 服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert, // 客户端验证（双向验证）
		ClientCAs:    certPool,                       // cert 池
	})

	return creds
}

// GetClientCert 获取客户端证书
func GetClientCert() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("../cert/client.pem", "../cert/client.key")
	if err != nil {
		log.Fatal("客户端生成cert异常，err:", err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../cert/ca.pem")
	if err != nil {
		log.Fatal("客户端读取ca.pem异常，err:", err)
	}
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, // 客户端证书
		ServerName:   "localhost",             // HTTP请求访问的域名
		RootCAs:      certPool,                // cert池
	})

	return creds
}
