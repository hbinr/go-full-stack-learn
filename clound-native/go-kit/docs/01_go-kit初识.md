[go-kit](https://github.com/go-kit/kit)是一个**微服务工具包合集**，并不是一个类似 spring-cloud 的框架，go-kit 使用灵活，方便扩展和组件化。它可以提供独立的包，通过这些包，开发者可以用来组建自己的应用程序。微服务架构意味着构建分布式系统，这带来了许多挑战，Go Kit 可以为多数业务场景下实施微服务软件架构提供指导和解决方案

## Go-kit 的三层架构

基于 Go Kit 的应用程序架构由三个主要部分组成：传输层、接口层和服务层。

### 1.Transport

传输层用于网络通信，服务通常使用 HTTP 或 gRPC 等网络传输方式，或使用 NATS 等发布订阅系统相互通信。除此之外，Go Kit 还支持使用 AMQP 和 Thrift 等多种网络通信模式。

### 2.Endpoint

定义 Request 和 Response 格式,并可以使用装饰器包装函数,以此来实现各种中间件嵌套。


Endpoint该层负责接收请求并返回响应。对于每一个服务接口，endpoint 层都使用一个抽象的 Endpoint 来表示 ，我们可以为每一个 Endpoint 装饰 Go-kit 提供的附加功能，如日志记录、限流、熔断等。
### 3.Service

服务层是具体的业务逻辑实现，包括核心业务逻辑。它不会也不应该进行 HTTP 或 gRPC 等具体网络传输，或者请求和响应消息类型的编码和解码。

提供具体的业务实现接口，endpoint 层中的 Endpoint 通过调用 service 层的接口方法处理请求。
## 使用

### 1.安装

`go get github.com/go-kit/kit`

### 2.使用

以一个简单的服务为例，如根据用户 ID 获取用户姓名。
[具体代码](../code/app/)

**第一步:创建 Service，主要是业务接口及业务类**

创建一个用户业务处理接口，并实现其方法：

```go
// IUserService 用户处理业务接口
type IUserService interface {
	GetUserName(id int) string
}

type UserService struct {
}

func (u *UserService) GetUserName(id int) string {
	if id == 101{
		return "101 httprouter_test"
	}
	return "guest"
}
```

**第二步:创建 Endpoint，定义 Request 和 Response 格式**

创建请求和响应参数，并生成 go-kit 特有的 endpoint

```go
// 请求
type UserRequest struct {
	UserID int `json:"id"`
}

// 响应
type UserResponse struct {
	UserName string `json:"userName"`
}

// 生成endpoint
func GenUserEndpoint(service service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		name := service.GetUserName(r.UserID)
		return UserResponse{
			UserName: name,
		}, nil
	}
}
```

**第三步:创建 Transport，主要负责与 HTP.gRPC. thrift 等相关的逻辑，对 request 进行解码 和对 response 进行编码**

对于响应:响应是我们发出的我们需要 encode,比较通用且简单的就是
json

```go
// DecodeUserRequest 解码请求参数
func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	// http://localhost:8080/get/1
	params := httprouter.ParamsFromContext(r.Context())
	if idStr := params.ByName("id"); idStr != "" {
		id, _ := strconv.Atoi(idStr)
		return endpoint.UserRequest{
			UserID: id,
		}, nil
	}

	return nil, errors.New("请求参数错误")
}

// EncodeUserResponse 将响应参数进行编码
func EncodeUserResponse(c context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}
```

**第四步:创建服务**

需要注意两点：

- 使用 go-kit 的 http 库创建 serverHandler，注册了 endpoint 和 transport 的编解码函数
- 使用 go 原生的 http 库启动和监听服务

```go
package main

import (
	"hb.study/go-kit/code/user/endpoint"
	"hb.study/go-kit/code/user/service"
	"hb.study/go-kit/code/user/transport"
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
	// 使用httprouter创建路由， 性能好于gorilla/mux
	r := httprouter.New()
	r.Handler(http.MethodGet, "/get/:id", serverHandle)
	// 启动服务
	http.ListenAndServe(":8080", r)
}
```
