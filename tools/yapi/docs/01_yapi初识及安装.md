# YApi
在开始使用 YApi 之前，我们先来熟悉一下 YApi 的网站结构，这将让你快速了解YApi。


## 登录与注册
想要使用 YApi ，首先要注册账号。
![](https://yapi.baidu.com/doc/documents/images/usage/login.png)
### 首页
登录后进入首页，首页展示了分组与项目。

此时你作为新用户，没有任何分组与项目的权限，因此只能搜索、浏览 “公开项目” 的接口，如果在首页找不到任何项目，请联系管理员将你加入对应项目。

1. 首页头部展示了当前所在的位置、搜索框、新建项目、查看文档和用户信息。
2. 首页左侧展示分组信息，“分组”是“项目”的集合，只有超级管理员可以管理分组。
3. 首页右侧是分组下的项目和成员列表，点击左侧的某个分组，右侧会出现该分组下的项目和成员信息。
4. 点击项目右上角的星星即可关注项目，关注的项目可以在“我的关注”页面查看。

![](https://yapi.baidu.com/doc/documents/images/usage/index.png)

### 项目页
点击一个项目，进入项目页，项目页展示了属于该项目的全部接口，并提供项目、接口的全部操作。

此时你作为新用户，只能浏览接口信息，不可以编辑项目或接口，如果需要编辑，请联系管理员将你加入该项目。

1. 项目页左侧的 “接口列表” 展示了该项目下的所有接口，右侧默认显示该项目下所有接口的列表。
2. 点击左侧的某个接口，右侧会出现“预览”、“编辑”和“运行”。
3. 点击左侧的 “测试集合” 使用测试集功能。
4. 点击二级导航的“设置”，项目组长即可编辑项目信息和管理成员列表。
5. 点击二级导航的“动态”，即可查看项目的操作日志。

![](https://yapi.baidu.com/doc/documents/images/usage/project.png)
### 个人中心
鼠标移动到右上角的用户头像或用户名上，即可点击“个人中心”查看个人信息。

![](https://yapi.baidu.com/doc/documents/images/usage/hover.png)

在个人信息页面可以查看并修改自己的用户名、密码等信息。
![](https://yapi.baidu.com/doc/documents/images/usage/user.png)


## 内网部署

### 环境要求
- nodejs（7.6+）
- mongodb（2.6+）

记得先启动了 mongodb
### 安装
**可视化部署[推荐]**
执行 yapi server 启动可视化部署程序，输入相应的配置和点击开始部署，就能完成整个网站的部署。部署完成之后，可按照提示信息，执行 `node/{网站路径/server/app.js}` 启动服务器。在浏览器打开指定url, 点击登录输入您刚才设置的管理员邮箱，默认密码(ymfe.org) 登录系统（默认密码可在个人中心修改）。

```
npm install -g yapi-cli --registry https://registry.npm.taobao.org
yapi server
```
### 部署yapi
执行 `yapi server`命令后，根据提示，在浏览器中输入：0.0.0.0:9090，会弹出以下页面：

![](https://s1.ax1x.com/2020/09/23/wjZPfA.png)

其中：
- 部署路径自定义
- 数据库地址需要将无改为自己启动`MongoDB`的地址

根据项目实际需要可以修改部署信息，点击【开始部署】，完成部署任务

常见部署错误：
- MongoDB数据库连接失败，检查是不是配置文件的ip、port出错。
- node版本较高，如下：

![](https://s1.ax1x.com/2020/09/23/wjKeiD.png)

解决办法就是[更换node版本](https://blog.csdn.net/m0_47404181/article/details/109771872)，比如 node 12.X就可以。

### 安装成功提示：
> (省略安装过程.............)
> 
> 依赖库安装完成，正在初始化数据库mongodb...
 
> yapi-vendor@1.9.2 install-server /home/hblock/Develop/my-yapi/vendors
>  node server/install.js


> log: mongodb load success...

> 初始化管理员账号成功,账号名："admin@admin.com"，密码："ymfe.org"

> 部署成功，请切换到部署目录，输入： "node vendors/server/app.js" 指令启动服务器, 然后在浏览器打开 http://127.0.0.1:3000 访问

### 启动yapi

cd到部署目录：
```
cd /home/hblock/develop/my-yapi
```

node启动：
```
node vendors/server/app.js
```

然后在浏览器打开 http://127.0.0.1:3000 访问