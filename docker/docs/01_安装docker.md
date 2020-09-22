# Windows10 安装 Docker For Windows

https://blog.csdn.net/qq_43573375/article/details/107543592

**如何在 Windows 上通过 Kitematic 使用 Docker**
https://linux.cn/article-6131-1.html?pr

# Linux 安装 docker

[centos 安装 Docker 官方教程](https://docs.docker.com/engine/install/centos/)

## 一.移除旧版本

如果之前安装过 docker，需要先基础旧版本

```sh
sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine

```

## 二.安装

### 1.把 yum 包更新到最新

```sh
sudo yum update
```

（期间要选择确认，输入 y 即可）

### 2.安装需要的软件包， yum-util 提供 yum-config-manager 功能，另外两个是 devicemapper 驱动依赖的

```sh
sudo yum install -y yum-utils device-mapper-persistent-data lvm2
```

参考：https://blog.csdn.net/u014069688/article/details/100532774

### 3.设置 yum 源（选择其中一个）

> yum-config-manager --add-repo http://download.docker.com/linux/centos/docker-ce.repo（官方中央仓库）

> yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo（阿里仓库）

如设置阿里仓库，后面步骤下载 docker 引擎等要快很多

```sh
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
```

### 4. 下载 docker 引擎、容器

默认下载最新版本

```sh
sudo yum install docker-ce docker-ce-cli containerd.io
```

## 三.使用 docker

### 1.启动 docker

```sh
sudo systemctl start docker
```

设置开机自启

```sh
sudo systemctl enable docker

```

重启 docker：

```sh
sudo systemctl restart docker

```

### 2.查看 docker 版本

```sh
docker -v

输出结果：
[vagrant@localhost ~]$ docker -v
Docker version 19.03.12, build 48a66213fe
```

显示版本号则表示安装成功

### 2.验证 docker 是否安装成功

```sh
sudo docker run hello-world

```

此命令下载测试镜像并在容器中运行。 容器运行时，它会打印参考消息并退出。

### 3.配置 docker 阿里云加速器

**1）登陆阿里云并打开容器镜像服务**

控制台-》容器镜像服务(未开通的需要开通)

**2）选择镜像加速器(在最底部)**

选择操作系统，如选择 `centos`，通过修改 daemon 配置文件/etc/docker/daemon.json 来使用加速器

全部命令如下：

```sh
sudo mkdir -p /etc/docker

sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://xxxxxx.mirror.aliyuncs.com"]
}
EOF

sudo systemctl daemon-reload

sudo systemctl restart docker
```

加速地址是自己阿里云账号生成的，填入自己的加速地址即可
