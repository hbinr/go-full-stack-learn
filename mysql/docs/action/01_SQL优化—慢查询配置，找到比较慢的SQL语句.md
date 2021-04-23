# SQL优化—慢查询解决

## 慢查询

查询慢的SQL，每个查询慢的SQL都会生成慢查询日志，我们可以根据这个日志来找到哪些SQL需要优化。

**慢查询日志**： 是指MySQL记录所有执行时间超过 `long_query_time` 参数设定的时间阈值的SQL语句的日志。

默认情况下，慢查询日志是关闭的，需要开启。

## 慢查询配置
慢查询日志是关闭的，需要开启。那我们怎么配置，怎么就能开启呢？

1. 首先找到MySQL的基础配置文件 `my.conf`（Linux下），Windows下为 `my.ini`
2. 打开该文件后，找到以下内容：
   1. `slow_quey_log=1`: 慢查询开启开关，1表示开启，0表示关闭
   2. `slow_query_log_file="DESKTOP-2EKGEE5-slow.log"`: 慢查询日志保存的文件名称
   3. `long_query_time=10`: 慢查询时间阈值，10s，默认以秒为单位
3. 根据自己需要修改上面3个配置项即可

当这些都配置好，去查看 `DESKTOP-2EKGEE5-slow.log` 文件的内容即可。该文件保存在 `mysql server 5.6/data/DESKTOP-2EKGEE5-slow.log`  （以MySQL 5.6版本为例）

该日志中需要关注三点：
- `Query_time`: 查询花费时间，以秒为单位
- `SET timestamp`: 查询开始时间，时间戳显示
- `select ... `: 完整查询SQL

但是该文件一般都会很大，直观去查找慢SQL的时候，需要我们自己去看耗时、SQL语句，这些可以通过工具解决的。


## 开启慢查询-设置步骤

### 1.查看慢查询相关参数
```sql

# 查看以slow_query为前缀相关配置项字段的配置情况
mysql> show variables like 'slow_query%';
+---------------------------+----------------------------------+
| Variable_name             | Value                            |
+---------------------------+----------------------------------+
| slow_query_log            | OFF                              |
| slow_query_log_file       | /mysql/data/localhost-slow.log   |
+---------------------------+----------------------------------+


# 查看long_query_time字段的配置情况
mysql> show variables like 'long_query_time';
+-----------------+-----------+
| Variable_name   | Value     |
+-----------------+-----------+
| long_query_time | 10.000000 |
+-----------------+-----------+
```

### 2.修改配置
**方法一：全局变量设置**
将 `slow_query_log` 全局变量设置为“ON”状态

```sql
mysql> set global slow_query_log='ON'; 
```

设置慢查询日志存放的位置
```sql
mysql> set global slow_query_log_file='/usr/local/mysql/data/slow.log';

```

查询超过1秒就记录
```sql
mysql> set global long_query_time=1;
```

**方法二：配置文件设置**

修改配置文件`my.cnf`：

1. 首先找到MySQL的基础配置文件 `my.cnf`（Linux下），Windows下为 `my.ini`
2. 打开该文件后，找到以下内容：
   1. `slow_quey_log=0`: 慢查询开启开关，1表示开启，0表示关闭
   2. `slpw_query_log_file="DESKTOP-2EKGEE5-slow.log"`: 慢查询日志保存的文件名称
   3. `long_query_time=10`: 慢查询时间阈值，10s，默认以秒为单位
3. 根据自己需要修改上面3个配置项即可

### 3.重启MySQL服务
```sql
service mysqld restart
```

### 4.查看设置后的参数

```sql
mysql> show variables like 'slow_query%';
+---------------------+--------------------------------+
| Variable_name       | Value                          |
+---------------------+--------------------------------+
| slow_query_log      | ON                             |
| slow_query_log_file | /usr/local/mysql/data/slow.log |
+---------------------+--------------------------------+

mysql> show variables like 'long_query_time';
+-----------------+----------+
| Variable_name   | Value    |
+-----------------+----------+
| long_query_time | 1.000000 |
+-----------------+----------+
```