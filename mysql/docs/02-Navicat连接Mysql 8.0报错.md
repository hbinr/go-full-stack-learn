# Navicat 连接 Mysql caching_sha2_password cannot be loaded :

参考：https://blog.csdn.net/dsl59741/article/details/107891415

## 原因

8.0 版本之前的 mysql 加密规则是 mysql_native_pssword，8.0 版本之后是 caching_sha2_password【强加密规则】，但是现在 navicat 还不支持最新的加密规则

## 解决

1. 进入 docker 容器内的 mysql 实例，修改 mysql 的加密规则为旧版本的【如果在非 docker 的情况下 navicat 连接 mysql，那么就直接在 centos 上进入自己的 Mysql，然后执行下边的修改密码即可，不用使用 docker 命令进入 mysql】

```sh
[root@localhost ~]# docker exec -it mysql bash        //进入Mysql的伪窗口
root@da94c9c49755:/# mysql -uroot -p   //回车输入开始docker运行镜像中的root密码进入mysql环境

```

2. 修改加密规则

```sh
mysql> use mysql;
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY 'root' PASSWORD EXPIRE NEVER;    //修改root用户的加密规则

mysql> ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root'; //修改root用户的密码为root
mysql> flush privileges;     //刷新权限
```

如果报错 ERROR 1396 (HY000): Operation ALTER USER failed for 'root'@'%' ：

则是远程访问权限不正确，先选择数据库，查看一下再更改：

```sh
use mysql;
Database changed

update user set host = 'localhost' where user ='root';

ALTER USER 'root'@'localhost' IDENTIFIED BY 'root' PASSWORD EXPIRE NEVER;

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'root';

update user set host = '%' where user ='root';

## 远程链接也直接就解决了

FLUSH PRIVILEGES;
```

## 测试

用 navicat 远程连接数据库的时候，注意把防火墙关闭或者把指定的端口开放，否则也连接不上
