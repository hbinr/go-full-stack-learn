# 

使用 `make proto` 命令来生产grpc相关代码时，如果 `google`的protobuf官方文件未放置到正确目录，执行 `make proto` 会报以下错误：
```
$ make proto
protoc --proto_path=. \
           --proto_path=E:\Workspace\go/pkg/mod/github.com/go-kratos/kratos/v2@v2.0.0-alpha3/api \
           --proto_path=E:\Workspace\go/pkg/mod/github.com/go-kratos/kratos/v2@v2.0.0-alpha3/third_party \
           --proto_path=E:\Workspace\go/src \

           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           --go-http_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. ./api/helloworld/descriptor.proto ./api/helloworld/errors/helloworld.proto ./api/helloworld/helloworld.proto ./api/helloworld/v1/greeter.proto ./internal/conf/conf.proto
google/protobuf/descriptor.proto: File not found.

kratos/api/annotations.proto:9:1: Import "google/protobuf/descriptor.proto" was not found or had errors.  #  关键错误提示


kratos/api/annotations.proto:11:8: "google.protobuf.EnumOptions" seems to be defined in "api/helloworld/descriptor.proto", which is not imported by "kratos/api/annotations.proto".  To use it here, please add the necessary import.
api/helloworld/errors/helloworld.proto:5:1: Import "kratos/api/annotations.proto" was not found or had errors.
make: *** [Makefile:17: proto] Error 1

```
重点看 `--proto_path` 相关内容，我故意用空行隔开了。

`proto_bath`：表示寻找 `.proto` 文件的路径

根据关键错误提示：
```go
Import "google/protobuf/descriptor.proto" was not found or had errors.
```

在上述所有 `--proto_path` 陈列的路径中找不到 `descriptor.proto`文件。

修改方法就是让找到即可，也就是将安装的 `protobuf` 相关目录移动到上面任意一个 `--proto_path` 路径下，要注意，官方的protobuf目录在include路径下，具体如图：
  
![grpc_include目录](../img/grpc_include目录.png)

- 如果是个人开发，推荐直接将 `google` 目录整体拷贝到你的 `GOROOT`目录下，省的以后还需要某个官方 `.proto` 文件时再拷贝。
- 如果是项目打包上线，可以将需要的 `.proto` 拷贝到项目中。
