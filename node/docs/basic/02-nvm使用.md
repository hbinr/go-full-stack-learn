# Nvm——Node.js版本管理工具

nvm (node.js version manager)

## nvm安装
**安装nvm之前，先删除node**

### Windows
1. 先删除
   1. 控制面板-》找到node-》删除即可
2. 后安装nvm
   1. 下载nvm:  https://github.com/coreybutler/nvm-windows/releases
   2. 执行安装程序即可
   3. 记得自行选择nvm安装路径和node安装路径
3. 配置nvm环境环境变量，Path中新建一个环境变量，填入nvm的安装路径
4. 测试
   1. 在cmd中执行命令： `nvm -v`，能显示出版本号，说明安装成功
5. 配置镜像源，下载就更快了
   1. 打开nvm安装路径，找到 `settings.txt`
   2. 填入以下内容：
```
node_mirror: https://npm.taobao.org/mirrors/node/
npm_mirror: https://npm.taobao.org/mirrors/npm/
```
## 安装/管理nodejs
1. 查看本地安装的所有版本；有可选参数available，显示所有可下载的版本。

`nvm list [available]`
2. 安装，命令中的版本号可自定义，具体参考命令1查询出来的列表

`nvm install 11.13.0`

3. 使用特定版本

`nvm use 11.13.0`

4. 卸载node

`nvm uninstall 11.13.0`
### nvm 命令提示
- nvm arch ：显示node是运行在32位还是64位。
- nvm install <version> [arch] ：安装node， version是特定版本也可以是最新稳定版本latest。可选参数arch指定安装32位还是64位版本，默认是系统位数。可以添加--insecure绕过远程服务器的SSL。
- nvm list [available] ：显示已安装的列表。可选参数available，显示可安装的所有版本。list可简化为ls。
- nvm on ：开启node.js版本管理。
- nvm off ：关闭node.js版本管理。
- nvm proxy [url] ：设置下载代理。不加可选参数url，显示当前代理。将url设置为none则移除代理。
- nvm node_mirror [url] ：设置node镜像。默认是https://nodejs.org/dist/。如果不写url，则使用默认url。设置后可至安装目录settings.txt文件查看，也可直接在该文件操作。
- nvm npm_mirror [url] ：设置npm镜像。https://github.com/npm/cli/archive/。如果不写url，则使用默认url。设置后可至安装目录settings.txt文件查看，也可直接在该文件操作。
- nvm uninstall <version> ：卸载指定版本node。
- nvm use [version] [arch] ：使用制定版本node。可指定32/64位。
- nvm root [path] ：设置存储不同版本node的目录。如果未设置，默认使用当前目录。
- nvm version ：显示nvm版本。version可简化为v。


参考:
https://www.cnblogs.com/gaozejie/p/10689742.html