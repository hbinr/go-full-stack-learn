# docker 安装 nacos 并连接 MySQL

#### 1.安装 nacos 镜像

这边没有安装最新版，而是指定了版本号。如需要安装最新版，执行以下命令

```sh
docker pull nacos/nacos-server

```

#### 2.查看 nacos 镜像

```sh
docker images

```

#### 3.创建本地的映射文件，custom.properties，

目录自己定义

```sh
mkdir -p /data/home/hblock/MyData/nacos/init.d  /data/home/hblock/MyData/nacos/logs
touch /data/home/hblock/MyData/nacos/init.d/custom.properties
```

在 custom.properties 文件中写入以下配置

```sh
management.endpoints.web.exposure.include=*
```

#### 4.创建数据库 nacos，并初始化 nacos 的 sql

**提示：**我本机 MySQL 的版本是 8.0+。

[nacos 官网 sql](https://github.com/alibaba/nacos/blob/master/config/src/main/resources/META-INF/nacos-db.sql)

因为 nacos 的相关配置信息是保存在 MySQL 中的，所以需要初始化官网提供的 nacos-db.sql，以后我们开发时新增的配置信息也在对应的表中。

- -env MODE=standalone 表示以单机形式启动的那个，nacos 1.3.0+ 版本后都默认是以集群形式(MODE=cluster)启动

访问 localhost:8848 登陆 nacos 管理页面， 账号密码默认都是 nacos

**需要注意：**

- 单机模式中`-e` 添加的参数不同与集群模式，集群模式见文章结尾。

单机模式中/home/nacos/conf/application.properties 配置文件如下：

```sh
# spring
server.servlet.contextPath=${SERVER_SERVLET_CONTEXTPATH:/nacos}
server.contextPath=/nacos
server.port=${NACOS_SERVER_PORT:8848}
spring.datasource.platform=${SPRING_DATASOURCE_PLATFORM:""}
nacos.cmdb.dumpTaskInterval=3600
nacos.cmdb.eventTaskInterval=10
nacos.cmdb.labelTaskInterval=300
nacos.cmdb.loadDataAtStart=false
db.num=${MYSQL_DATABASE_NUM:1}
db.url.0=jdbc:mysql://${MYSQL_SERVICE_HOST}:${MYSQL_SERVICE_PORT:3306}/${MYSQL_SERVICE_DB_NAME}?characterEncoding=utf8&connectTimeout=1000&socketTimeout=3000&autoReconnect=true
db.url.1=jdbc:mysql://${MYSQL_SERVICE_HOST}:${MYSQL_SERVICE_PORT:3306}/${MYSQL_SERVICE_DB_NAME}?characterEncoding=utf8&connectTimeout=1000&socketTimeout=3000&autoReconnect=true
db.user=${MYSQL_SERVICE_USER}
db.password=${MYSQL_SERVICE_PASSWORD}
### The auth system to use, currently only 'nacos' is supported:
nacos.core.auth.system.type=${NACOS_AUTH_SYSTEM_TYPE:nacos}


### The token expiration in seconds:
nacos.core.auth.default.token.expire.seconds=${NACOS_AUTH_TOKEN_EXPIRE_SECONDS:18000}

### The default token:
nacos.core.auth.default.token.secret.key=${NACOS_AUTH_TOKEN:SecretKey012345678901234567890123456789012345678901234567890123456789}

### Turn on/off caching of auth information. By turning on this switch, the update of auth information would have a 15 seconds delay.
nacos.core.auth.caching.enabled=${NACOS_AUTH_CACHE_ENABLE:false}

server.tomcat.accesslog.enabled=${TOMCAT_ACCESSLOG_ENABLED:false}
server.tomcat.accesslog.pattern=%h %l %u %t "%r" %s %b %D
# default current work dir
server.tomcat.basedir=
## spring security config
### turn off security
nacos.security.ignore.urls=/,/error,/**/*.css,/**/*.js,/**/*.html,/**/*.map,/**/*.svg,/**/*.png,/**/*.ico,/console-fe/public/**,/v1/auth/**,/v1/console/health/**,/actuator/**,/v1/console/server/**
# metrics for elastic search
management.metrics.export.elastic.enabled=false
management.metrics.export.influx.enabled=false

nacos.naming.distro.taskDispatchThreadCount=10
nacos.naming.distro.taskDispatchPeriod=200
nacos.naming.distro.batchSyncKeyCount=1000
nacos.naming.distro.initDataRatio=0.9
nacos.naming.distro.syncRetryDelay=5000
nacos.naming.data.warmup=true

```

**有个坑：**

在按 nacos 1.2.+的时候，需要修改 nacos 配置 MySQL 链接 db.url，因为 mysql8 及其以上版本需要带时区，所以还需要修改 db.url，新增下面内容：

```sh
serverTimezone=Asia/Shanghai
```

但是在重新安装 nacos 1.3.2 时，不需要新增时区的配置了，使用默认的 application.properties 即可

#### 5.创建容器

这里的指定参数根据/home/nacos/conf/application.properties 配置设置的，主要是 mysql 配置的修改。

```sh
docker run -d -p 8848:8848  \
-e MODE=standalone \
-e PREFER_HOST_MODE=hostname \
-e MYSQL_SERVICE_HOST=127.0.0.1 \
-e MYSQL_SERVICE_PORT=3306 \
-e MYSQL_SERVICE_DB_NAME=nacos \
-e MYSQL_SERVICE_USER=root \
-e MYSQL_SERVICE_PASSWORD=123456 \
-e MYSQL_DATABASE_NUM=1 \
-v /data/home/hblock/MyData/nacos/init.d/custom.properties:/home/nacos/init.d/custom.properties \
-v /data/home/hblock/MyData/nacos/logs:/home/nacos/logs \
--restart always --name nacos nacos/nacos-server
```

**参数解析：**

- -e MODE=standalone 以单机模式运行，nacos 1.2 之后，默认以集群(cluster)方式运行

[集群配置详情见 Nacos 集群配置参数官方文档](https://nacos.io/zh-cn/docs/quick-start-docker.html)

#### 6.启动容器

```sh
docker start nacos
```

#### 7.测试

访问 http://localhost:8848/nacos/ ，账号默认 nacos、密码默认 nacos

_参考：_

https://www.it235.com/%E9%AB%98%E7%BA%A7%E6%A1%86%E6%9E%B6/SpringCloudAlibaba/nacos.html#mysql%E6%94%AF%E6%8C%81

https://www.cnblogs.com/niunafei/p/12803965.html
