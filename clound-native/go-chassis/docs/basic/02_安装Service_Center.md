# 安装服务中心Service Center


## 可执行文件安装
## Docker安装

### 拉取镜像
```sh
docker pull servicecomb/service-center
```
### 启动容器
```sh
docker run -p 30100:30100   --name servicecenter -d servicecomb/service-center
```

**注意：**

Service Center发行版使用了精巧的etcd，如果要使用etcd的单独实例，则可以单独部署etcd并在此处配置etcd ip
```sh
vi conf/app.conf

## Edit this file
# registry address
# 1. if registry_plugin equals to 'embeded_etcd'
# manager_name = "sc-0"
# manager_addr = "http://127.0.0.1:2380"
# manager_cluster = "sc-0=http://127.0.0.1:2380"
# 2. if registry_plugin equals to 'etcd'
# manager_cluster = "127.0.0.1:2379"
manager_cluster = "127.0.0.1:2379"
```

Service Center默认的地址和端口是`127.0.0.1:2379`，可以在`app.conf`文件中修改IP和端口
```sh
vi conf/app.conf

httpaddr = 127.0.0.1
httpport = 30100
```
### 查看容器是否启动成功
```sh
hblock@hblock:~$ docker ps                                                                                                                                                
CONTAINER ID        IMAGE                        COMMAND                  CREATED             STATUS              PORTS                      NAMES                        
ad9f5d8588d4        servicecomb/service-center   "/opt/service-center…"   5 seconds ago       Up 3 seconds        0.0.0.0:30100->30100/tcp   servicecenter   
```
