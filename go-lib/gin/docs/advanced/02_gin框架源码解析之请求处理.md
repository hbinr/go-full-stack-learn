gin 框架 Engine 实现 ServeHTTP 接口代码：

```go
// ServeHTTP conforms to the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    // engine.pool 通过对象池去减少每次申请创建临时对象、内存申请、垃圾回收操作的消耗
	c := engine.pool.Get().(*Context)
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}

```
