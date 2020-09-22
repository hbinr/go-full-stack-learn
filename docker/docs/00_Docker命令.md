# Docker 命令

### 1、Docker 容器信息

```sh
##查看docker容器版本
docker version

##查看docker容器信息
docker info

##查看docker容器帮助
docker --help
```

### 2、镜像操作

**提示：对于镜像的操作可使用镜像名、镜像长 ID 和短 ID。**

#### 2.1、镜像查看

```sh
##列出本地 images
docker images
##含中间映像层
docker images -a
```

![](https://img2018.cnblogs.com/blog/1659331/201905/1659331-20190521104721523-485290950.png)

```sh
##只显示镜像 ID
docker images -q
##含中间映像层
docker images -qa
```

![](https://img2018.cnblogs.com/blog/1659331/201905/1659331-20190521104927909-600452122.png)

```sh
##显示镜像摘要信息(DIGEST 列)
docker images --digests
##显示镜像完整信息
docker images --no-trunc
```

![](https://img2018.cnblogs.com/blog/1659331/201905/1659331-20190521105114405-1780655005.png)

```sh
##显示指定镜像的历史创建；参数：-H 镜像大小和日期，默认为 true；--no-trunc 显示完整的提交记录；-q 仅列出提交记录 ID
docker history -H redis
```

#### 2.2、镜像搜索

```sh
##搜索仓库MySQL镜像
docker search mysql

## --filter=stars=600：只显示 starts>=600 的镜像
docker search --filter=stars=600 mysql

## --no-trunc 显示镜像完整 DESCRIPTION 描述
docker search --no-trunc mysql

## --automated ：只列出 AUTOMATED=OK 的镜像
docker search  --automated mysql
```

![](https://img2018.cnblogs.com/blog/1659331/201905/1659331-20190521110514156-691788920.png)

#### 2.3、镜像下载

```sh
##下载Redis官方最新镜像，相当于：docker pull redis:latest
docker pull redis
##下载仓库所有Redis镜像
docker pull -a redis
##下载私人仓库镜像
docker pull bitnami/redis
```

![](https://img2018.cnblogs.com/blog/1659331/201905/1659331-20190521112716615-10141164.png)

#### 2.4、镜像删除

```sh
##单个镜像删除，相当于：docker rmi redis:latest
docker rmi redis
##强制删除(针对基于镜像有运行的容器进程)
docker rmi -f redis
##多个镜像删除，不同镜像间以空格间隔
docker rmi -f redis tomcat nginx
##删除本地全部镜像
docker rmi -f $(docker images -q)
```

#### 2.5、镜像构建

```sh
##（1）编写dockerfile
cd /docker/dockerfile
## 编写相关配置
vim mycentos

##（2）构建docker镜像
docker build -f /docker/dockerfile/mycentos -t mycentos:1.1
```

### 3、容器操作

**提示：对于容器的操作可使用 CONTAINER ID 或 NAMES。**

#### 3.1、容器启动

```sh
##新建并启动容器，参数：-i 以交互模式运行容器；-t 为容器重新分配一个伪输入终端；--name 为容器指定一个名称
docker run -i -t --name mycentos
##后台启动容器，参数：-d 已守护方式启动容器
docker run -d mycentos
```

**注意：**

此时使用`docker ps -a`会发现容器已经退出。

这是 docker 的机制：要使 Docker 容器后台运行，就必须有一个前台进程。

解决方案：将你要运行的程序以前台进程的形式运行。

```sh
##启动一个或多个已经被停止的容器
docker start redis

##重启容器
docker restart redis
```

#### 3.2、容器进程

```sh
##top支持 ps 命令参数，格式：docker top [OPTIONS] CONTAINER [ps OPTIONS]
##列出redis容器中运行进程
docker top redis

##查看所有运行容器的进程信息
for i in  `docker ps |grep Up|awk '{print $1}'`;do echo \ &&docker top $i; done
```

#### 3.3、容器日志

```sh
##查看redis容器日志，默认参数
docker logs rabbitmq
##查看redis容器日志，参数：-f  跟踪日志输出；-t   显示时间戳；--tail  仅列出最新N条容器日志；
docker logs -f -t --tail=20 redis
##查看容器redis从2019年05月21日后的最新10条日志。
docker logs --since="2019-05-21" --tail=10 redis
```

#### 3.4、容器的进入与退出

```sh
##使用run方式在创建时进入
docker run -it redis /bin/bash
##关闭容器并退出
exit
##仅退出容器，不关闭
快捷键：Ctrl + P + Q
##直接进入redis容器启动命令的终端，不会启动新进程，多个attach连接共享容器屏幕，参数：--sig-proxy=false  确保CTRL-D或CTRL-C不会关闭容器
docker attach --sig-proxy=false redis
##在 redis 容器中打开新的交互模式终端，可以启动新进程，参数：-i  即使没有附加也保持STDIN 打开；-t  分配一个伪终端
docker exec -i -t  redis /bin/bash
##以交互模式在容器中执行命令，结果返回到当前终端屏幕
docker exec -i -t redis ls -l /tmp
##以分离模式在容器中执行命令，程序后台运行，结果不会反馈到当前终端
docker exec -d redis  touch cache.txt
```

#### 3.5、查看容器

```sh
##查看正在运行的容器
docker ps
##查看正在运行的容器的ID
docker ps -q
##查看正在运行+历史运行过的容器
docker ps -a
```

![](https://img2018.cnblogs.com/blog/1659331/201905/1659331-20190521132255698-500560462.png)

```sh
##显示运行容器总文件大小
docker ps -s
```

![](https://img2018.cnblogs.com/blog/1659331/201905/1659331-20190521133039811-1994116017.png)

```sh
##显示最近创建容器
docker ps -l
##显示最近创建的3个容器
docker ps -n 3
##不截断输出
docker ps --no-trunc
```

![](https://img2018.cnblogs.com/blog/1659331/201905/1659331-20190521132741451-294716433.png)

```sh
##获取镜像redis的元信息
docker inspect redis
##获取正在运行的容器redis的 IP
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' redis
```

#### 3.6、容器的停止与删除

```sh
##停止一个运行中的容器
docker stop redis
##杀掉一个运行中的容器
docker kill redis
##删除一个已停止的容器
docker rm redis
##删除一个运行中的容器
docker rm -f redis
##删除多个容器
docker rm -f $(docker ps -a -q)
docker ps -a -q | xargs docker rm
## -l 移除容器间的网络连接，连接名为 db
docker rm -l db
## -v 删除容器，并删除容器挂载的数据卷
docker rm -v redis
```

#### 3.7、生成镜像

```sh
##基于当前redis容器创建一个新的镜像；
##参数：-a 提交的镜像作者；-c 使用Dockerfile指令来创建镜像；-m :提交时的说明文字；-p :在commit时，将容器暂停
docker commit -a="DeepInThought" -m="my redis" [redis容器ID]  myredis:v1.1
```

#### 3.8、容器与主机间的数据拷贝

```sh
##将rabbitmq容器中的文件copy至本地路径
docker cp rabbitmq:/[container_path] [local_path]
##将主机文件copy至rabbitmq容器
docker cp [local_path] rabbitmq:/[container_path]/
##将主机文件copy至rabbitmq容器，目录重命名为[container_path]（注意与非重命名copy的区别）
docker cp [local_path] rabbitmq:/[container_path]
```

参考：
https://www.cnblogs.com/DeepInThought/p/10896790.html
