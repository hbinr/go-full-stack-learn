# 创建第一个REST示例

## 安装go-chassis

直接使用 `go get github.com/go-chassis/go-chassis` 命令会默认下载`v1.8.3`版本，并不是最新的

```sh
$ go get github.com/go-chassis/go-chassis
go: downloading github.com/go-chassis/go-chassis v1.8.3
go: github.com/go-chassis/go-chassis upgrade => v1.8.3
go: downloading github.com/go-chassis/go-archaius v1.2.1
go: downloading github.com/opentracing/opentracing-go v1.0.2
go: downloading github.com/go-mesh/openlogging v1.0.1
go: downloading github.com/go-chassis/go-restful-swagger20 v1.0.2
go: downloading github.com/prometheus/client_model v0.0.0-20190115171406-56726106282f
go: downloading github.com/prometheus/procfs v0.0.0-20190117184657-bf6a532e95b1
go: downloading golang.org/x/net v0.0.0-20191004110552-13f9640d40b9
go: downloading github.com/emicklei/go-restful v2.11.1+incompatible
go: downloading github.com/go-chassis/foundation v0.1.1-0.20191113114104-2b05871e9ec4
go: downloading github.com/go-chassis/paas-lager v1.0.2-0.20190328010332-cf506050ddb2
go: downloading golang.org/x/sys v0.0.0-20190826190057-c7b8b68b1456
go: downloading k8s.io/utils v0.0.0-20191114200735-6ca3b61696b6
```

截止目前2020-12-12，最新版本为`v2.1.0`，**通过指定版本下载最新的**
```sh
go get github.com/go-chassis/go-chassis/v2@v2.1.0
```

## 启动Service Center


```sh
docker start servicecenter
```
- `servicecenter` 为我创建容器时，自己指定的容器名，此处写为自己的设置的容器名或者直接用容器ID

## 编写服务
### 请求
```go
package server

import (
	"net/http"

	"github.com/go-chassis/go-chassis/v2/server/restful"
	"github.com/go-chassis/openlog"
)

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UserController struct {
}

// PostUser .
func (r *UserController) PostUser(ctx *restful.Context) {
	var (
		req LoginRequest
		err error
	)
	if err = ctx.ReadEntity(&req); err != nil {
		openlog.Error("server:LoginRequest ReadEntity failed" + err.Error())
		return
	}

	ctx.WriteJSON(&req, "application/json") // 既然已经是WriteJson了，为什么还要指定Content-Type类型？
}

// URLPatterns 帮助响应相应的API调用，必须实现，底层会自动掉接口。类似我们平时写的在RegisterHandle中设置路由
func (u *UserController) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method: http.MethodPost, Path: "/user", ResourceFunc: u.PostUser},
	}
}
```


### 初始+注册服务
`main.go` 主要做了三件事：

1. `chassis.RegisterSchema` 注册对外暴露的协议+服务
2. `chassis.Init` 初始化
3. `chassis.Run()` 启动服务

**代码如下：**
```go
package main

import (
	"github.com/go-chassis/go-chassis/v2"
	"github.com/go-chassis/openlog"

	"hb.study/clound-native/go-chassis/code/basic/rest/server"
)

func main() {
	chassis.RegisterSchema("rest", &server.UserController{}) // 注册对外暴露的协议+服务
	if err := chassis.Init(); err != nil {
		openlog.Fatal("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
```
注意:
> `chassis.RegisterSchema`注册服务时，要保证参数为指针，即 `&server.UserController{}`，否则会报错 Parse services from config failed: input must be an ptr
## 编写配置文件
在 `conf`目录下编写配置文件

### 设置 chassis.yaml

`chassis.yaml` 可以理解为应用的配置文件，配置的是微服务的公共属性，如公共的AppId信息，使用的注册中心类型信息、地址信息，服务的协议、监听地址、注册发现地址、传输协议信息等；

```yaml
servicecomb:
  registry:  # 注册中心（ServiceCenter）的地址，默认为 127.0.0.1:30100 
      address: http://127.0.0.1:30100 
  protocols: # 传输协议，REST
    rest:    # 开发REST接口
      listenAddress: 127.0.0.1:8081  # http 服务的地址和端口
```
### 设置 microservice.yaml
`microservice.yaml` 顾名思义，针对微服务设置的文件，配置的是微服务的私有属性，包括服务名、版本等

```yaml
servicecomb:
  service:
	name: hello.user.service  # 自定义你的服务provider名
	version: 0.0.1            # 版本号
```
## 设置 CHASSIS_HOME 或者CHASSIS_CONF_DIR

这两个环境变量设置其中任意一个都可以
### CHASSIS_HOME

go-chassis会自动读取`main.go`所在目录下的`conf/chassis.yaml`和`conf/microservice.yaml`

因此`CHASSIS_HOME`的设置取决你的程序启动目录。

看下`main.go`路径：
```sh
$pwd 
/home/hblock/go/src/study/go-chassis-demo/rest
```

那么设置`CHASSIS_HOME`为该值
```sh
export CHASSIS_HOME=/home/hblock/go/src/study/go-chassis-demo/rest
```

两个命令合并一下也可以：
```sh
export CHASSIS_HOME=$(pwd) 
or 
export CHASSIS_HOME=$PWD  # 大写pwd
```
### CHASSIS_CONF_DIR
看下配置文件所在的路径 

```sh
cd conf/
$ pwd

/home/hblock/go/src/study/go-chassis-demo/rest/conf
```
然后设置配置文件的路径
```sh
export CHASSIS_CONF_DIR=/home/hblock/go/src/study/go-chassis-demo/rest/conf
or 
export CHASSIS_CONF_DIR=$PWD  # 大写pwd
```
## 启动程序

启动过程会在终端打印很多日志：init相关+自己写的服务注册等


打开Postman测试，`http://localhost:8081/user`，设定好`Body`测试。

当我回去看终端输出时，看到了下面一行日志,，而且是隔几秒就会定时答应出看来：

> {"level":"DEBUG","timestamp":"2020-12-12 23:57:40.342 +08:00","file":"servicecenter/servicecenter.go:112","msg":"heartbeat success, microServiceID/instanceID: 2f8499f9f3062ad73a058529e67a1524f4b153ea/b136eb5b3c9211eb99a80242ac110002."}

`msg`中包含了**健康检查**、**服务ID**、**实例ID**，这些值都是唯一的
- heartbeat success
- microServiceID/instanceID: 2f8499f9f3062ad73a058529e67a1524f4b153ea/b136eb5b3c9211eb99a80242ac110002.