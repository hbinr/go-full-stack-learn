# grpc服务端和客户端新增SSL证书调用

## 一.生成证书
[只看证书生成步骤即可](https://www.jianshu.com/p/5938432e2130)

如果不想去鼓捣这些证书的生成，可以直接使用[cert](../code/demo_product/cert/)目录下的相关文件，可直接在本地使用，域名是:`localhost`

## 二.生成服务端和客户端证书，双向验证
具体代码：[cert.go](../code/demo_product/util/cert.go)

```go
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
```

## 三.引入证书

### 1.服务端引入
主要变动，在创建grpc server的时候：
```go

// 1.创建一个 rpc server，通过CA和SSL证书来进行TLS验证
rpcServer := grpc.NewServer(grpc.Creds(util.GetServerCert()))
```
### 2.客户端引入
主要变动，在创建grpc客户端连接的时候：

```go
// 1.创建连接套接字，端口和服务端一致，并开启 TLS 验证
conn, err := grpc.Dial(":8092", grpc.WithTransportCredentials(util.GetClientCert()))
```