# 使用 TiUP Playground 快速部署本地测试环境

- 适用场景：利用本地 Mac 或者单机 Linux 环境快速部署 TiDB 集群。可以体验 TiDB 集群的基本架构，以及 TiDB、TiKV、PD、监控等基础组件的运行。
- 耗时：1 分钟

作为一个分布式系统，最基础的 TiDB 测试集群通常由 2 个 TiDB 实例、3 个 TiKV 实例和 3 个 PD 实例来构成。通过 TiUP Playground，可以快速搭建出上述的一套基础测试集群。
## 1.下载并安装 TiUP。
```sh
curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh
```

下载成功示例：
```sh
hblock@hblock:~$ curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 8608k  100 8608k    0     0  1836k      0  0:00:04  0:00:04 --:--:-- 1836k
WARN: adding root certificate via internet: https://tiup-mirrors.pingcap.com/root.json
You can revoke this by remove /home/hblock/.tiup/bin/7b8e153f2e2d0928.root.json
Set mirror to https://tiup-mirrors.pingcap.com success
Detected shell: bash
Shell profile:  /home/hblock/.bashrc
/home/hblock/.bashrc has been modified to add tiup to PATH
open a new terminal or source /home/hblock/.bashrc to use it
Installed path: /home/hblock/.tiup/bin/tiup
===============================================
Have a try:     tiup playground
===============================================
```

## 2.声明全局环境变量。

**注意：**

> TiUP 安装完成会提示对应的 profile 文件的绝对路径，以下 source 操作需要根据实际位置进行操作。这个和个人本地环境有关

看第一步中的提示，`open a new terminal or source /home/hblock/.bashrc to use it`，我们直接执行：

```sh
source /home/hblock/.bashrc
```
## 3.在当前 session 执行以下命令启动集群。
有两种方式，看自己需求：
### 方式一：使用默认配置启动集群

直接运行 tiup playground 命令会运行最新版本的 TiDB 集群，其中 TiDB、TiKV 和 PD 实例各 1 个：
```sh
tiup playground
```
执行该命令后，会下载安装组件`palyground`

### 方式二：指定 TiDB 版本以及各组件实例个数
```sh
tiup playground v4.0.0 --db 2 --pd 3 --kv 3 --monitor
```

其中:
- ` v4.0.0 `：表示会在本地下载并启动一个 v4.0.0 版本的集群
- `--db 2`：表示TiDB实例启动1个
- `--pd 3`：表示PD实例启动3个
- `--kv 3`：表示TiKV实例启动3个
- `--monitor`：表示同时部署监控组件

 最新版本可以通过执行 `tiup list tidb` 来查看。 运行结果将显示集群的访问方式：

```sh
CLUSTER START SUCCESSFULLY, Enjoy it ^-^
To connect TiDB: mysql --host 127.0.0.1 --port 4000 -u root
To connect TiDB: mysql --host 127.0.0.1 --port 4001 -u root
To view the dashboard: http://127.0.0.1:2379/dashboard
To view the monitor: http://127.0.0.1:9090
```

## 4.新开启一个 session 以访问 TiDB 数据库。

1. 首先安装 MySQL 客户端。如果已安装 MySQL 客户端则可跳过这一步骤。
```sh
yum -y install mysql
```

2. 使用 MySQL 客户端连接 TiDB：
```sh
mysql --host 127.0.0.1 --port 4000 -u root
```
成功示例：
```sh
hblock@hblock:~$ mysql --host 127.0.0.1 --port 4000 -u root
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 3
Server version: 5.7.25-TiDB-v4.0.8 TiDB Server (Apache License 2.0) Community Edition, MySQL 5.7 compatible

Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```

看到`Server version: 5.7.25-TiDB-v4.0.8 TiDB Server` 则表示连接成功

## 5.通过 http://127.0.0.1:9090 访问 TiDB 的 Prometheus 管理界面。
## 6.通过 http://127.0.0.1:2379/dashboard 访问 TiDB Dashboard 页面，默认用户名为 root，密码为空。
## 7.测试完成后清理集群，绿色环保。通过 ctrl-c 停掉进程后，执行以下命令：
```sh
tiup clean --all
```
**注意：**
> TiUP Playground 默认监听 127.0.0.1，服务仅本地可访问；若需要使服务可被外部访问，你可以通过 `--host` 参数指定监听 0.0.0.0 或网卡绑定外部可访问的 IP。