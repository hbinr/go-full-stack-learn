# Gin 框架源码解析

## c.Next()

我们先看下二者的源码，核心是： `c.index`，表示 handleFunc(即注册的中间件、路由处理函数) 的下标，这直接决定了该 handleFunc 是否执行

`c.Next()`源码：

```go
func (c *Context) Next() {
    c.index++
    // int8(len(c.handlers)) 的最大值为63
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c) // 执行下标命中的方法
		c.index++
	}
}
```

为什么 `int8(len(c.handlers))`的最大值为 63 呢？因为在

```go
// 示例方法
func main() {
	r := gin.Default()

	r.GET("/student", func(c *gin.Context) {
		c.String(http.StatusOK, "使用String()做出响应:%s", "GET 请求..")
    })
    r.Run()
}
```

### 源码跟踪

查看 `r.GET`的源码，发现会调用 `RouterGroup` 的 `handle` 方法

```go
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	return group.handle(http.MethodGet, relativePath, handlers)
}
```

`group.handle()` 源码：

```go
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
    // 计算路由URL，返回路由路径
    absolutePath := group.calculateAbsolutePath(relativePath)
    // 拼接handle
    handlers = group.combineHandlers(handlers)
    // 注册路由
	group.engine.addRoute(httpMethod, absolutePath, handlers)
	return group.returnObj()
}
```

`combineHandlers()`源码：

```go
func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
    finalSize := len(group.Handlers) + len(handlers)
    // 如果需要处理的请求handle超过一定数量(63个)就会退出这次请求,
	if finalSize >= int(abortIndex) {
		panic("too many handlers")
	}
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}
```

## c.Abort()

`c.Abort` 源码：

```go
// gin's context.go
const abortIndex int8 = math.MaxInt8 / 2  // 63

// 其他代码
// ...

func (c *Context) Abort() {
	c.index = abortIndex
}
```

可以看到 `c.index`的值直接赋值为最大值，那么 c.Next() c.index 极值判断

```go
c.index < int8(len(c.handlers))

```

就不满足条件了，直接退出 for 循环，也就不行后续的 handleFunc 了
