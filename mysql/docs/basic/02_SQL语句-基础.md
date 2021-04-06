# 基础SQL语句

## 别名 as  
```sql
select empno as '员工编号' from emp;

# 等同于：
select empno '员工编号' from emp;
```
别名可以省略
## 查询结果使用中文展示
```sql
select empno as '员工编号',sal * 12 as '年薪' from emp;
```
标准SQL语句中，**字符串(包括中文)**尽量使用**单引号**括起来，虽然MySQL支持双引号，但不建议使用，因为Oracle、SQL server中都只识别单引号。
