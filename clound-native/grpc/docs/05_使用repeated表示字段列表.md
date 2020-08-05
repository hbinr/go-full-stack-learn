# repeated 表示字段列表

在实际开发中，通常参数可能是一个list对象，在PB中是通过 repeated 来表示该字段可以重复，这种字段可以重复任意多次（包括0次）。

## 一.定义 .porto 文件
远程服务：通过pageSize参数来控制返回的商品名称数量。

**知识点：**
- 使用 repeated 来修饰响应参数，表示其可能返回的是列表
- 消息 message 可以当做字段引用，也可以被导入到其他 `.proto` 文件中使用。
```go
message ProdResponse {
  string prodName=1;
}

message QueryRequest {
  int32 pageSize=1;
}

message ProdListResponse {
  repeated ProdResponse prodList = 1;
}

service ProdService{
  rpc GetProdNameList(QueryRequest) returns(ProdListResponse);
}
```
## 二.生成代码

这里使用了一个小技巧，因为生成代码命令可能会经常使用，所以封装到了 bat 文件中(Windows下)，以后直接在终端运行该脚本文件即可。

封装内容：
```sh
cd .. && cd pbfile && protoc --go_out=plugins=grpc:../service --go_opt=paths=source_relative product.proto
```

由于是单独新建了一个 bin 目录，所以需要cd  bin 目录下，终端执行：`genGo.bat`即可

如果我们开发的时候有个多服务，难道需要写多个 bat 脚本吗？ 

当然是不需要的，将上述命令末尾 `product.proto` 改为 `*.proto`，就可以将所有的 `.proto` 文件都生成对应go代码