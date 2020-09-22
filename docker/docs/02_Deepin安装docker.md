# Deepin v20 正式版安装 Docker

## 简介

Docker 是一个开源的应用容器引擎，让开发者可以打包他们的应用以及依赖包到一个可移植的容器中，然后发布到任何流行的 Linux 机器上，也可以实现虚拟化。容器是完全使用沙箱机制，相互之间不会有任何接口。

## 关于 Deepin 中的 Docker

Deepin 官方的应用仓库已经集成了 docker，但不是类似于 docker-ce 这样的最新版本。由于 Deepin 是基于 debian 的 unstable 版本开发的，通过 \$(lsb_release -cs) 获取到的版本信息为 unstable，而 docker 官方源并没支持 debian 的 unstable 版本，因此使用 docker 官方教程是安装不成功的。如果你需要安装 docker-ce，请遵循下面的步骤进行安装：

## 在 Deepin 中安装 Docker 最新版的方法

#### 1.如果以前安装过老版本，要确保先卸载以前版本

```sh
sudo apt-get remove docker.io docker-engine
```

#### 2.安装密钥管理与下载相关的工具

```sh
## 密钥管理（add-apt-repository ca-certificates 等）与下载（curl 等）相关的工具
sudo apt-get install apt-transport-https ca-certificates curl python-software-properties software-properties-common
```

此处可能会提示 `python-software-properties` 未安装成功或其他的问题，最后导致 `add-apt-repository` 命令无法使用

**解决办法：**

```sh
sudo apt-get install python-software-properties
sudo apt-get update
sudo apt install software-properties-common
sudo apt-get update
```

#### 3.下载并安装密钥

鉴于国内网络问题，强烈建议使用国内源，官方源请在注释中查看。

国内源可选用清华大学开源软件镜像站或中科大开源镜像站，示例选用了中科大的。

为了确认所下载软件包的合法性，需要添加软件源的 GPG 密钥。

```sh
## 中科大源
curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/debian/gpg | sudo apt-key add -

## 官方源，能否成功可能需要看运气。
## curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
```

#### 4.查看密钥是否安装成功

sudo apt-key fingerprint 0EBFCD88

如果安装成功，会出现如下内容：

```sh
pub 4096R/0EBFCD88 2017-02-22    Key fingerprint = 9DC8 5822 9FC7 DD38 854A E2D8 8D81 803C 0EBF CD88

uid Docker Release (CE deb) <docker@docker.com>

sub 4096R/F273FCD8 2017-02-22
```

#### 5.在 source.list 中添加 docker-ce 软件源（请先查看后面的 Note）：

```sh
sudo add-apt-repository "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/debian buster stable"

## 官方源
## sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian buster stable"

## 15.10 会提示 aptsources.distro.NoDistroTemplateException: Error: could not find a distribution template for Deepin/stable

## 这里我们通过编辑 sudo vim /etc/apt/sources.list 添加一行即可，原因未知
sudo add-apt-repository "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/debian stretch stable"
```

**Note：** 官方在 buster 位置使用的是 \$(lsb_release -cs)，但之前已经解释过，在 deepin 里运行它得到的是 unstable，docker 官方不支持 unstable 版本！因此直接使用官方教程的命令会安装失败。

**更改方法：**将上述命令中的版本名称 buster，替换成 deepin 基于的 debian 版本对应的代号。查看版本号的命令为：

```sh
cat /etc/debian_version.
```

**举例：**

a). 对于 deepin v20，我操作上面的命令得到 debain 版本是 10.5，debian 10.5 的代号是 buster

b). 对于 deepin 15.9.2 基于 debian 9.0 , debian 9.0 的代号为 stretch, 所以 deepin 15.9.2 上完整的添加信息为:

```sh
sudo add-apt-repository "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/debian stretch stable"
```

具体代码可以去查一下：[debain 版本](https://www.debian.org/releases/index.zh-cn.html)

#### 6.更新仓库

```sh
sudo apt-get update
```

#### 7.安装 docker-ce

由于网络不稳定，可能会下载失败。如果下载失败了，可以多试几次或者找个合适的时间继续。

```sh
sudo apt-get install docker-ce
```

#### 8.命令行查看 docker 版本

```sh
docker version
```

#### 9.让普通用户也可运行 docker

上面第 8 步查看 docker 版本的时候如图末尾显示权限不足，是因为 docker 只允许 root 用户执行，为让普通用户也可运行 docker，执行

```sh
sudo usermod -aG docker username
```

将当前用户加入 docker 用户组，然后注销用户重新登录即可。

#### 10.启动 docker：

```sh
systemctl start docker
```

#### 11.验证 docker 是否被正确安装并且能够正常使用

```sh
sudo docker run hello-world
```

如果能够正常下载，并能够正常执行，则说明 docker 正常安装。

## 更换国内的 docker 加速器

如果使用 docker 官方仓库，速度会很慢，所以更换国内加速器就不可避免了。

#### 1.使用阿里云的 docker 加速器。

在阿里云申请一个账号
打开连接 https://cr.console.aliyun.com/#/accelerator 拷贝您的专属加速器地址。

修改 daemon 配置文件 /etc/docker/daemon.json 来使用加速器（下面是 4 个命令，分别单独执行）
Note： 这里的 https://jxus37ad.mirror.aliyuncs.com 是申请者的加速器地址，在此仅仅用于演示，而使用者要个根据自己的使用的情况填写自己申请的加速器地址。

```sh
sudo mkdir -p /etc/docker

## 创建配置文件
sudo tee /etc/docker/daemon.json <<-'EOF'
{
"registry-mirrors": ["https://jxus37ad.mirror.aliyuncs.com"]
}
EOF

## 重新加载配置文件
sudo systemctl daemon-reload
## 重启docker
sudo systemctl restart docker
```

其实 `daemon.json` 主要就是配置镜像源，不同账号，该源地址是不同的：

```sh
{
"registry-mirrors": ["https://jxus37ad.mirror.aliyuncs.com"]
}
```

#### 2.重启 docker 服务

```sh
sudo service docker restart
```

## 禁止开机自启

默认情况下 docker 是开机自启的，如果我们想禁用开机自启，可以通过安装 chkconfig 命令来管理 Deepin 自启项：

#### 安装 chkconfig

```sh
sudo apt-get install chkconfig
```

#### 移除自启

```sh
sudo chkconfig --del docker
```
