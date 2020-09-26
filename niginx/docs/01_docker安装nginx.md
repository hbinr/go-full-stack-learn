# Docker 安装 nginx

## 搜索并拉取镜像

1.搜索镜像

```sh
docker serach niginx
```

也可以去 [DockerHub 官网](https://hub.docker.com/)搜索镜像，可以看到具体的帮助文档

2.拉取镜像

```sh

# 默认拉取最新版本
docker pull nginx

```

## 创建 nginx 相关目录

目录可自定义

```sh
mkdir -p /home/hblock/MyData/nginx/www /home/hblock/MyData/nginx/logs /home/hblock/MyData/nginx/conf

```

- www： 目录将映射为 nginx 容器配置的虚拟目录。
- logs： 目录将映射为 nginx 容器的日志目录。
- conf： 目录里的配置文件将映射为 nginx 容器的默认配置文件。

创建这些目录的目的是要挂载到 docker 容器内，方便从主机修改 nginx 的配置，这样就不用每次进入 docker 容器修改了

## 首先创建一个测试的 nginx，并复制相关配置文件到主机，方便配置

主配置在/etc/nginx/nginx.conf 文件中，然而平时用到的 server 配置在/etc/nginx/conf.d/ 的目录下

```sh
docker run --name test -d nginx

```

因为不能挂载文件，只能挂载文件夹，所以先在一个 test 容器中复制一份配置文件，先复制 nginx.conf

```sh
docker cp test:/etc/nginx/nginx.conf /home/hblock/MyData/nginx/conf/
```

然后复制 default.conf 文件

```sh
docker cp test:/etc/nginx/conf.d/default.conf  /home/hblock/MyData/nginx/conf.d/

```

配置文件先不需要改。

为什么要先弄配置文件？因为一般情况下 docker 启动时会进行配置，只要把配置文件的目录挂载出来就可以，简洁方便，但是 nginx 却是先加载一个主配置文件 nginx.conf，在 nginx.conf 里再加载 conf.d 目录下的子配置文件（一般最少一个 default.conf 文件）

注意 default.conf 文件里面，网页访问路径就是你容器里面 nginx 的网站路径。写成宿主机的会出错。

```sh
location / {
        root   /usr/share/nginx/html; # 这个写容器里面nginx的路径
        index  index.html index.htm;
    }
```

因为这个配置文件最终会被 docker 容器解析给 nginx 所以应该配置 docker 的 nginx 容器下的路径，然后再使用 docker 容器构建时候的目录映射功能将 docker 下的 nginx 目录映射到自己主机上操作。

```s
location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }
```

## 运行容器，并挂载目录

```sh
docker run \
    --name nginx \
    -p 80:80 \
    -v /home/hblock/MyData/nginx/www:/usr/share/nginx/html:rw \
    -v /home/hblock/MyData/nginx/conf/nginx.conf:/etc/nginx/nginx.conf:rw \
    -v /home/hblock/MyData/nginx/conf.d:/etc/nginx/conf.d:rw \
    -v /home/hblock/MyData/nginx/logs:/var/log/nginx \
    -d nginx
```

- –name nginx：将容器命名为 nginx。
- -p 80:80：将容器的 80 端口映射到主机的 80 端口。
- -v /home/hblock/MyData/nginx/www:/usr/share/nginx/html：将我们自己创建的 www 目录(项目位置)挂载到容器的 /usr/share/nginx/html 目录。
- -v /home/hblock/MyData/nginx/conf/nginx.conf:/etc/nginx/nginx.conf：将我们自己创建的 nginx.conf 文件挂载到容器的 /etc/nginx/nginx.conf 文件。
- -v /home/hblock/MyData/nginx/logs:/var/log/nginx：将我们自己创建的 logs 挂载到容器的 /var/log/nginx 目录。
- -v /home/hblock/MyData/nginx/conf.d:/etc/nginx/conf.d：将我们自己创建的 conf.d 目录挂载到容器的 /etc/nginx/conf.d 目录。
- -d nginx：以守护进程的方式启动 Nginx。
- :rw 配置权限

## 查看容器运行情况

```sh
hblock@hblock:~$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                               NAMES
721b4e76c2e3        nginx               "/docker-entrypoint.…"   About a minute ago   Up About a minute   0.0.0.0:80->80/tcp                  nginx
```

## 添加静态页面并访问

在 /home/hblock/MyData/nginx/www 目录下新建 index.html 文件，内容自己写即可，测试用

然后重启容器：

```sh
docker restart nginx

```

浏览器访问 http://localhost，显示出自己编写的 index.html 内容即表示安装成功

参考：

[docker 中部署 nginx 镜像挂载文件夹和文件并解决出错](https://blog.csdn.net/qq_42114918/article/details/85238011)

[nginx 踩坑---nginx 配置文件问题](https://blog.csdn.net/qq_31404603/article/details/85942546)
