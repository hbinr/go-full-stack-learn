# Docker 安装GitLab
## 创建挂载目录
在宿主机上创建以下目录，这个自定义：
```sh
mkdir -p ~/MyData/gitlab/config ~/MyData/gitlab/logs ~/MyData/gitlab/data
```
## 拉取镜像
默认拉取最新的
```sh
docker pull gitlab/gitlab-ce
```
## 创建容器并启动
```sh
docker run -d -p 443:443 -p 80:80 -p 222:22 --name gitlab --restart no \
-v /home/hblock/MyData/gitlab/config:/etc/gitlab \
-v /home/hblock/MyData/gitlab/gitlab/logs:/var/log/gitlab \
-v /home/hblock/MyData/gitlab/gitlab/data:/var/opt/gitlab gitlab/gitlab-ce
```
## 进入容器，修改相关配置
**进入`gitlab`容器**
```sh
docker exec -it gitlab bash
```
## 修改配置
**注意：以下操作都在 gitlab 容器中执行**

**修改/etc/gitlab/gitlab.rb：**
```sh
vim /etc/gitlab.rb
#unicorn['port'] = 8080 修改 8070  默认是注释的去掉前面的#
unicorn['port'] = 8070
#nginx['listen_port'] = nil 修改 8090  默认是注释的去掉前面的#
nginx['listen_port'] = 8090
```
**修改/var/opt/gitlab/gitlab-rails/etc/unicorn.rb：**
```sh
vim /var/opt/gitlab/gitlab-rails/etc/unicorn.rb
 
#listen "127.0.0.1:8080", :tcp_nopush => true
listen "127.0.0.1:8070", :tcp_nopush => true
```
**修改默认的gitlab nginx的web服务80端 /var/opt/gitlab/nginx/conf/gitlab-http.conf**
```sh
vim /var/opt/gitlab/nginx/conf/gitlab-http.conf
 
#listen *:80;
listen *:8090;
```

**重新配置gitlab：**
```sh
gitlab-ctl reconfigure
```

**重新启动gitlab：**
```sh
gitlab-ctl restart
```
## GitLab服务控制
因为是使用docker安装的，所以直接使用docker命令来管理：
```sh
# 启动
docker run gitlab
# 停止
docker stop gitlab
# 删除容器
docker rm gitlab
```