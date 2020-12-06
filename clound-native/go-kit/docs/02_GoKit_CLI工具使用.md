# GoKit CLI 是基于go-kit的工具

[github](https://github.com/GrantZheng/kit)

这个工具可以根据我们的需求自动生成`service`，`transport`和`endpoint`模版，以及生成供调用方的使用的client library，节省我们大量的时间，提高我们的生产效率。



## 安装
```
git clone https://github.com/GrantZheng/kit.git
cd kit
go install 
```

## 使用

### 查看帮助文档
```
kit help
```

### 创建一个新服务

`kit`默认情况下使用`go module`来管理依赖项，请确保go版本> = 1.3，或已将`GO111MODULE`设置为打开。 

如果要指定模块名称，则应使用`--module` 参数，否则`go.mod`文件中的模块名称默认为项目名称。

**以User服务为例：**
```sh
## 查看`new service`帮助文档
kit new service --help 
kit new service app

```
或者使用别名，更简单的命令，效果是一致的：
```
## or using aliases
kit n s user
```

生成的目如下：
```sh
app/
├── go.mod
└── pkg
    └── service
        └── service.go
```

其中：
`go.mod`的内容如下：

```go
module user

go 1.15
```

这里没有使用`--module`来指定模块名称，所以默认是`user`

`service.go`的内容如下：
```go
package service

// UserService describes the service.
type UserService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
}
```

