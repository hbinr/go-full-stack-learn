# grpc-go安装

## 一.grpc-go 安装

如果是使用go moudle 开发项目，直接导入以下包，IDE会自动安装：
```go
import "google.golang.org/grpc"
```
或者在终端执行命令：
```sh
$ go get -u google.golang.org/grpc
```

## 二.安装PB编译器

### 1.Windows安装
[2020-8-03]目前最新3.12.4版本安装：

https://github.com/protocolbuffers/protobuf/releases/tag/v3.12.4

下载后，将bin目下的执行文件 protoc.exe 移动到 GOPATH/bin目录下

