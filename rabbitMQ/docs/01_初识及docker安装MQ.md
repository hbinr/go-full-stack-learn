# 初识

# docker 安装 rabbitMQ
## 安装
1. 先创建一个目录用来专门存放mq对应的数据
```sh
mkdir 

``` 
2. 获取rabbit镜像，默认下载最新版本：
```
docker pull rabbitmq:management
```
1. 执行以下命令：
```sh
docker run -d 
--hostname rabbit-server
--name myrabbit 
-v /mydata/rabbitmq:/var/lib/rabbitmq
-e RABBITMQ_DEFAULT_USER=admin 
-e RABBITMQ_DEFAULT_PASS=admin 
-p 15672:15672 -p 5672:5672 -p 25672:25672 -p 61613:61613 -p 1883:1883 
rabbitmq:management
```
**命令解释：**
- docker run -d    以后台形式运行
- --hostname rabbit-server   设置rabbitMQ的主机名
- --name myrabbit    设置当前启动的容器的名字
- -v /mydata/rabbitmq:/var/lib/rabbitmq   挂载主机目录和docker容器目录
- -e RABBITMQ_DEFAULT_USER=admin   设置用户名，不设置默认是guest
- -e RABBITMQ_DEFAULT_PASS=admin   设置密码，不设置默认是guest
- -p 5672:5672 -p 15672:15672  -p 25672:25672 -p 61613:61613 -p 1883:1883    设置端口映射:5672->用于访问API的端口；15672->用于访问rabbitMQ UI管理界面端口；25672->用于集群节点之间通信的端口；
- rabbitmq:management     容器镜像源

## 验证是否安装成功

浏览器访问： `localhost:15672`，但是却无法访问，需要进行以下设置：

1. 确保RabbitMQ的端口等配置正确，进入RabbitMQ中，开启一项配置。

例：开启RabbitMQ  　　
```
docker run -itd --name myrabbitmq -p 15672:15672 -p 5672:5672 rabbitmq
```

2. 进入RabbitMQ　　
```
docker exec -it myrabbitmq /bin/bash
```
3. 开启管理界面配置
```
rabbitmq-plugins enable rabbitmq_management
```
