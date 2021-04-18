# Kafka安装


## 安装
个人是直接安装的kafaka，使用了自带的zookeeper。

安装地址：http://kafka.apache.org/downloads.html

![](https://img-blog.csdn.net/20180301183305251)

## 配置修改
`config` 目录下修改两个配置文件 `server.properties`： 
### server.properties

主要修改:
#### `listeners=PLAINTEXT://:9092` 
监听地址，端口号就默认9092即可


#### `advertised.listeners=PLAINTEXT://your.host.name:9092`

如果该值没有配置，则使用 `listeners`配置的值；当然如果配置了，就直接使用自己配置的值了


#### `log.dirs=/tmp/kafka-log`
日志存放目录。修改为自定义的目录即可。如Windows下可以修改为：`log.dirs=E:/mydata/kafka/log`


#### `zookeeper.connect=localhost:2181`
zookeeper的地址，如果就是单机，个人电脑使用，就不用改了，直接使用默认值 `localhost:2181`

如果是线上环境，有多个集群，那么需要指定zookeeper的连接地址。如: "zookeeper.connect=127.0.0.1:3000,127.0.0.1:3001,127.0.0.1:3002"

如果有多个地址，记得使用 `,` 隔开

#### `zookeeper.connection.timeout.ms=18000`
连接到Zookeeper的超时时间(以毫秒为单位)

### zookeeper.properties 
#### `dataDir=tmp/kafka-data` 
存储快照的目录，一般用来存储Kafka的数据。

可以改成自定义的目录，如Windows下改成：`dataDir=E:/mydata/kafka/`


## 启动

### 在kafka安装的目录下打开终端
```sh
$/dev-software/Kafkakafka_2.13-2.7.0
```
### 先启动zookeeper：
```sh
$ bin/windows/zookeeper-server-start.bat config/zookeeper.properties
```
### 后启动kafaka：
```sh
$ bin/windows/kafka-server-start.bat config/server.properties

```
没报错的话，就成功启动了

参考：

[Kafka安装教程（详细过程）,包含了每个配置项的解释](https://blog.csdn.net/Poppy_Evan/article/details/79415460  )
