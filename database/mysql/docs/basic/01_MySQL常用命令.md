# MySQL常用命令

## 查看当前mysql版本
```sql
select version();
```
## 查询当前使用的数据库
```sql
select database();
```
## 查看当前库中的表
```sql
show tables;
```
## 查看表结构  
```sql
desc 表名;
```

## 查看表创建语句
```sql
show create table <表名>

mysql> show create table dept;

CREATE TABLE `dept` (
  `DEPTNO` int(2) NOT NULL,
  `DNAME` varchar(14) DEFAULT NULL,
  `LOC` varchar(13) DEFAULT NULL,
  PRIMARY KEY (`DEPTNO`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4
```
