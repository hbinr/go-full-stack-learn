
先在nacos创建`dubbo.properties` 
```go
dubbo.service.com.ikurento.user.UserProvider.cluster=failback
dubbo.service.com.ikurento.user.UserProvider.protocol=myDubbo
dubbo.protocols.myDubbo.port=20000
dubbo.protocols.myDubbo.name=dubbo
```
然后在启动服务