# SQL语句--排序

关键词 `order by`
## `order by` 默认是升序，从小到大

`order by 字段 asc`，其中 `asc`可以省略不写，默认就是 `asc`升序

```sql
# 查询员工姓名、薪资，并按薪资排序（升序）
select
  ename,
  sal
from
  EMP
order by
  sal
```
## `order by desc` 降序，从大到小
```sql
# 查询员工姓名、薪资，并按薪资排序（降序）
select
  ename,
  sal
from
  EMP
order by
  sal desc
```

## 按照多个字段排序
`order by 字段1,字段2......`

MySQL进行排序时，是按照指定字段的顺序进行排序的
```sql
# 查询员工姓名、薪资， 并按薪资排序（升序）
# 如果薪资一样，则再按照名字排序
select
  ename,
  sal
from
  EMP
order by
  sal,   # 薪资按升序排序 也可以写为 sal asc
  ename  # 姓名按降序排序
  desc
```

上述sql中，`sal`字段在前，起主导作用，只有 `sal`相等的情况下，`ename` 才会起作用

## 按照字段的位置进行排序
```sql
select
  ename,
  sal
from
  EMP
order by
  2 
```
`order by`后面的数字表示根据查询结果中的第二列(也就是  `sal`)进行排序。

实际开发中不推荐使用，作为了解即可。因为不易阅读，而且列的顺序很容易发生改变
