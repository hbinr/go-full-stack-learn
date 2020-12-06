# Docker 安装 MySQL 并连接

**命令都为 root 角色下使用**

## 一.查找镜像

```sh
docker search mysql

```

也可以去官网查看镜像 tag，选择自己需要的版本，否则会下载最新版本：https://hub.docker.com/_/mysql/

## 二.下载镜像

可以指定想要的版本，如 5.7 版本，不指定则为最新版：

```sh
docker pull mysql:5.7
```

## 三.查看镜像是否下载成功

查看 docker 现有的镜像：

```sh
docker images

输出结果：
[root@localhost docker]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
mysql               5.7                 8679ced16d20        7 days ago          448MB
```

看到 mysql 镜像表示下载成功：

- `TAG`表示 mysql 版本
- `IMAGE ID` 表示镜像 ID

## 四.mysql 配置

参考：https://www.cnblogs.com/qiaoxin/p/10844492.html

1. 在 /mydata/mysql/conf 目录下新建 my.cnf 文件，填入以下内容：

```sh
 # mysql 5.7 配置
[client]
default-character-set = utf8

[mysql]
default-character-set = utf8

[mysqld]
init_connect='SET collation_connection = utf8_unicode_ci'
init_connect='SET NAMES utf8'
character-set-server=utf8
collation-server=utf8_unicode_ci
skip-character-set-client-handshake
skip-name-resolve

 # mysql 8.0+ 配置
[client]
default-character-set=utf8

[mysql]
default-character-set=utf8

[mysqld]
app=mysql
character-set-server=utf8
default_authentication_plugin=mysql_native_password
secure_file_priv=/var/lib/mysql
expire_logs_days=7
sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
max_connections=1000

```

- 设置编码为 utf8
- skip-name-resolve 跳过域名解析，解决 mysql 连接慢的问题
- secure_file_priv=/var/lib/mysql ，MYSQL 新特性 secure_file_priv 对读写文件的影响，需要在配置中加入该行内容

### 五.通过镜像创建 mysql 容器并运行

```sh
## mysql 5.7
docker run -p 3306:3306 --name mysql \
-v /mydata/mysql/conf:/etc/mysql \
-v /mydata/mysql/log:/var/log/mysql \
-v /mydata/mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
-d mysql:5.7 \


## mysql 8.0  /mydata和/home/hblock/MyData 目录是自定义的
docker run -p 3306:3306 --name mysql \
-v /home/hblock/MyData/mysql/conf:/etc/mysql \
-v /home/hblock/MyData/mysql/log:/var/log/mysql \
-v /home/hblock/MyData/mysql/data:/var/lib/mysql \
-e MYSQL_ROOT_PASSWORD=123456 \
-d mysql
```

- --name ：当前启动的容器的名字

- -p 3306:3306：mysql 容器内部的默认端口为 3306 ，为了保持通信，需要将容器的 3306 端口映射到原 Linux 主机的 3306 端口

- -v /mydata/mysql/conf:/etc/mysql：将主机目录/mydata/mysql/conf 挂载到 mysql 容器的 /etc/mysql 目录--配置

- -v /mydata/mysql/log:/var/log/mysql：将主机目录/mydata/mysql/log 挂载到 mysql 容器的 /var/log/mysql 目录--日志

- -v /mydata/mysql/data:/var/lib/mysql ：将主机目录/mydata/mysql/data 挂载到 mysql 容器的 /var/lib/mysql--保存的数据文件

- -e MYSQL_ROOT_PASSWORD=123456：初始化 root 用户的密码

- -d mysql:5.7：以后台方式运行，使用 mysql:5.7 镜像启动容器

安装好的 mysql 容器可以理解为是在一个独立“Linux”环境中运行。

-v 参数都是表示挂载目录，挂载完成后，直接操作 Linux 主机对应的目录就相当于操作 mysql 容器的相关目录了。不需要每次进入 mysql 容器去操作了，直接在 Linux 主机操作即可

### 五.查看正在运行的镜像进程

```sh
docker ps
```

发现没有输出结果

**问题：**

- 执行 docker run ... 后，容器处于 exited 状态，希望能够出入 up 状态，可以 exec 进去查看

**原因：**

- docker 容器执行任务完成后就会处于 exited 状态

先使用以下命令查看所有容器，包括不运行的容器

```sh
docker ps -a

```

**输出结果：**

```
hblock@hblock:~/MyData$ docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                     PORTS               NAMES
5a6f4db70942        mysql               "docker-entrypoint.s…"   12 seconds ago      Exited (1) 9 seconds ago                       mysql
      0.0.0.0:3306->3306/tcp, 33060/tcp   mysql
```

可以看到刚刚容器名为 mysql 的容器已经存在了，并且有个 `Exited (1) 9 seconds ago` 状态，表示容器已经退出

这是 docker 的机制：要使 Docker 容器后台运行，就必须有一个前台进程。

解决方案：将你要运行的程序以前台进程的形式运行。

```sh
##启动一个或多个已经被停止的容器
docker start redis


##或者重启容器
docker restart redis
```

如果还不行，仍旧存在上述的问题，很可能是容器里的运行的代码报错了，然后容器 Exited (1) 3 seconds ago 了

`docker logs -f container_id `能看到哪里错了

目录挂载错误请看：https://www.cnblogs.com/linjiqin/p/11465804.html

我遇到的问题是需要再配置文件中加入：`secure_file_priv=/var/lib/**mysql`

### 六.进入容器

安装好的 mysql 容器可以理解为是在一个独立“Linux”环境中运行。所以我们也可以进入 mysql 容器，就像在操作 Linux 系统

运行以下命令，`mysql`表示容器的名字，也可以使用容器 ID

```sh
docker exec -it mysql /bin/bash

输出结果：
[root@localhost /]# docker exec -it mysql /bin/bash
root@cd5561897cf2:/#
```

- -i: 以交互模式运行容器，通常与 -t 同时使用；
- -t: 为容器重新分配一个伪输入终端，通常与 -i 同时使用；
- root@cd5561897cf2:/# ---》 表示 root 角色，后面的数字和字母组合表示该容器的 ID

`exit`命令可退出当前容器

## 八.正确的启动方式：

如果我们 centos 关机了，下次还想启动 mysql 容器，光使用 `docker run mysql:5.7`是不可行的，会提示你：

> Database is uninitialized and password option is not specified You need to specify one of MYSQL_ROOT_PASSWORD, MYSQL_ALLOW_EMPTY_PASSWORD and MYSQL_RANDOM_ROOT_PASSWORD

正确应该是；

```sh
[root@xxxxxx ~]# docker run --name mysql_01 -e MYSQL_ROOT_PASSWORD=123456  -d mysql
```

注意 `--name mysql_01` 这个名字可以自己指定，但是不要和以存在的容器名冲突

## 九.Kitematic docker 管理终端设置 mysql 连接

mysql 容器启动成功后，可以通过"EXEC"查看数据库启动是否成功:

```sh
mysql -u root -p 123456
```

参考：
https://www.jianshu.com/p/d297b0be4157
