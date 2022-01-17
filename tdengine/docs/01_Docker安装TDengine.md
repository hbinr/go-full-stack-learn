# Docker 安装TDengine

## 通过镜像创建 tdengine 容器并运行
其中 `/home/hbinr/MyData` 目录是自定义的

```sh
docker run -d --name tdengine \
--hostname="tdengine-server" 
-v /home/hbinr/MyData/tdengine/log:/var/log/taos \
-v /home/hbinr/MyData/tdengine/data:/var/lib/taos \
-p 6030-6041:6030-6041 \
-p 6030-6041:6030-6041/udp \
tdengine/tdengine
```

- `--name` ：当前启动的容器的名字
- `-p 6030-6041:6030-6041` \    将容器的 6030 到 6041 端口映射到宿主机的 6030 到 6041 端口上
- `-p 6030-6041:6030-6041/udp` \    将容器的 6030 到 6041 端口映射到宿主机的 6030 到 6041 端口上
  - 为了支持 TDengine 客户端操作 TDengine server 服务， **TCP 和 UDP 端口都需要打开**。
- `-v /mydata/tdengine/log:/var/log/taos/mydata/tdengine/log` 挂载到 tdengine 容器的 /var/log/taos 目录--日志
- `-v /mydata/tdengine/data:/var/lib/taos/mydata/tdengine/data` 挂载到 tdengine 容器的 /var/lib/taos--保存的数据文件



安装好的 tdengine 容器可以理解为是在一个独立“Linux”环境中运行。

`-v` 参数都是表示挂载目录，挂载完成后，直接操作 Linux 主机对应的目录就相当于操作 tdengine 容器的相关目录了。不需要每次进入 tdengine 容器去操作了，
直接在 Linux 主机操作即可

**补充：在windows上使用docker安装MySQL**

```sh
docker run -d --name tdengine --hostname="tdengine-server"  \
-v E:\\mydata\\tdengine\\log:/var/log/taos \
-v E:\\mydata\\tdengine\\log:/var/lib/taos  \
-p 6030-6041:6030-6041 \
-p 6030-6041:6030-6041/udp \
tdengine/tdengine
```

## 查看正在运行的镜像进程
```sh
docker ps
CONTAINER ID   IMAGE               COMMAND   CREATED          STATUS          PORTS                                                                          NAMES
427c38cc7417   tdengine/tdengine   "taosd"   1 minutes ago   Up 1 minutes   0.0.0.0:6030-6041->6030-6041/tcp, 0.0.0.0:6030-6041->6030-6041/udp, 6042/tcp   tdengine
```
## 进入容器，验证安装是否成功
1. 进入容器
```sh
docker exec -it tdengine /bin/bash
```

2. 执行 `taos` 命令，输出一下内容表示安装成功
```sh
$ taos
Welcome to the TDengine shell from Linux, Client Version:2.4.0.4
Copyright (c) 2020 by TAOS Data, Inc. All rights reserved.
taos>
```

在 TDengine 终端中，可以通过 SQL 命令来创建/删除数据库、表、超级表等，并可以进行插入和查询操作。具体可以参考 [TAOS SQL 说明文档](https://www.taosdata.com/cn/documentation/taos-sql)

