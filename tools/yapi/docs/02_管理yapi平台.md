## pm2方式管理进程

要保证后台保持进程，需要安装pm2。全局安装:
```
npm install -g pm2
```
启动yapi：
```
pm2 start /root/my-yapi/vendors/server/app.js --watch -i 1
```
- `--watch`：监听应用目录的变化，一旦发生变化，自动重启。如果要精确监听、不见听的目录，最好通过配置文件。
- `-i --instances`：启用多少个实例，可用于负载均衡。如果-i 0或者-i max，则根据当前机器核数确定实例数目。 

通过pm2 save保存当前进程状态：

查看：
```
pm2 list 
```

重启：
```
pm2 restart /root/my-yapi/vendors/server/app.js --watch -i 1

```
停止：

可以先通过`pm2 list`获取应用的名字（--name指定的）或者进程id。
```
pm2 stop /root/my-yapi/vendors/server/app.js


pm2 stop app_name|app_id
```
如果要停止所有应用，可以:
```
pm2 stop all
```
升级:

升级项目版本是非常容易的，并且不会影响已有的项目数据，只会同步 vendors 目录下的源码文件。

```
cd {项目目录}
yapi ls //查看版本号列表
yapi update //更新到最新版本
yapi update -v {Version} //更新到指定版本
```


更多请看[PM2实用入门指南](https://imweb.io/topic/57c8cbb27f226f687b365636)