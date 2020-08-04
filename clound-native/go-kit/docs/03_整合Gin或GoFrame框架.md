# 整合 web 框

go-kit 三层架构不用变，只需要在启动 server 的时候使用 web 框架即可

和 main.go 同级目录下新增一个 `server.go`(server.go 的目录可自定义，demo 演示便和 main.go 同目录了)，然后写入以下内容：

```go
// NewHTTPServer 结合gin框架启动server
func NewHTTPServer(ctx context.Context, endpoints Endpoints) *gin.Engine {
	router := gin.Default()
	router.Use(commonMiddlewar())
	disk := router.Group("/user")
	{
		v1 := disk.Group("/v1")
		{
			v1.GET("/userinfo", func(ctx *gin.Context) {
                httptrasport.NewServer(endp, transport.DecodeUserRequest, transport.EncodeUserResponse)
			})
		}
	}

	return router
}

// commonMiddlewar 中间件
func commonMiddlewar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Add("Context-Type", "application/json")
	}
}
```

修改 `main.go`中的 **newserver** 部分：

```go
func main() {

	// 创建服务结构体
	userService := service.UserService{}
	// 生成endpoint
	endp := endpoint.GenUserEndpoint(userService)
    // 创建服务，结合gin框架
    server := NewHTTPServer(ctx, eps)
	// 启动服务
	server.Run(":8080")
}


```
