## 优势
### 性能强
### 完美集成go原生http

## 缺陷
### 路由冲突
```go
r := httprouter.New()
r.Handler(http.MethodGet, "/get/:id", serverHandle)
r.Handler(http.MethodGet, "/get/info/:id", serverHandle)
http.ListenAndServe(":8080", r)
```

报错信息：
> panic: 'info' in new path '/get/info/:id' conflicts with existing wildcard ':id' in existing prefix '/get/:id'

>goroutine 1 [running]:



补充：

[chi](https://github.com/go-chi/chi)也完美集成原生http，但是其性能要差点。比不上httprouter