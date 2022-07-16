# SQL语句练习
**涉及到的表：**

**部门表：**
```sql
CREATE TABLE `DEPT` (
  `DEPTNO` int(2) NOT NULL comment '部门编号',
  `DNAME` varchar(14) DEFAULT NULL comment '部门名称',
  `LOC` varchar(13) DEFAULT NULL comment '部门位置',
  PRIMARY KEY (`DEPTNO`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

```
**员工表：** 与部门表有关联，通过 `DEPTNO` 部门编号字段
```sql
CREATE TABLE `EMP` (
  `EMPNO` int(4) NOT NULL comment '员工编号',
  `ENAME` varchar(10) DEFAULT NULL comment '员工姓名',
  `JOB` varchar(9) DEFAULT NULL comment '工作',
  `MGR` int(4) DEFAULT NULL comment '领导编号',
  `HIREDATE` date DEFAULT NULL comment '入职时间',
  `SAL` double(7,2) DEFAULT NULL comment '薪资',
  `COMM` double(7,2) DEFAULT NULL comment '补助',
  `DEPTNO` int(2) DEFAULT NULL comment '部门编号',
  PRIMARY KEY (`EMPNO`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
```

**职级表：**
```sql
CREATE TABLE `SALGRADE` (
  `GRADE` int(11) DEFAULT NULL comment '职级序号',
  `LOSAL` int(11) DEFAULT NULL comment '最低工资',
  `HISAL` int(11) DEFAULT NUL Lcomment '最高工资'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
```
## 找出工资在[1250,3000]的员工信息，要求按薪资降序排列

```sql
select
  ename,
  sal
from
  EMP
where
  sal between 1250
  and 3000
order by
  sal desc;
```