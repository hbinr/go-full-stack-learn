# SQL语句--条件查询

## 语法格式
```sql
select 
    字段1,字段2,字段...
from
    表名
where
    条件;
```
**执行顺序：一定要理清楚**

1. 先from，找到表
2. 然后where(条件)，筛选过滤数据
3. 然后select，查出想要的字段
4. 最后order by(排序)，根据指定字段将结果排序
   
**示例：**
```sql
SELECT
  dname
FROM
  DEPT
where
  DEPTNO = 10;
```
最基础的条件有：
- `>`或 `>=`  大于 或 大于等于
- `>`或 `<=`  小于 或 小于等于
- `<>` 或 `!=`  不等于
## between and
范围查询，左闭右闭，**等同于 >= and <=**
```sql
# 查询工资在[1000,3000]的员工姓名
SELECT
  ename
FROM
  emp
where
  sal between 1000
  and 3000;
```

## is null 和 is not null
```sql
# 查询津贴为null的员工姓名
SELECT
  ename,comm
FROM
  emp
where
  comm is null;
```
**注意: **  条件 <没有津贴> 不能使用 `comm = null` 来查询。因为在MySQL中null不是一个值，不能用等号进行衡量，它是一个特殊的内置的关键词，表示空。

零值，如空字符串、0、0.00 这些都是零值。实际查询中零值和值null的数据都要设置对应的条件。

```sql
# 查询没有津贴的员工姓名
SELECT
  ename,comm
FROM
  emp
where
  comm is null
  or comm = 0;
```

## and 和 or的优先级问题 
当 `and` 和 `or` 联合使用的时候，会涉及到二者的优先级问题。

```sql
# 找出薪资大于1000并且部门编号是20或30的员工
SELECT
  ename,sal,deptno
FROM
  emp
where
  sal > 1000
  and deptno = 20
  or deptno = 30;
```

上述SQL是存在问题，当 `and` 和 `or` 一起使用的时候，`and`的优先级要**高于** `or`，所以上述SQL的条件部分，其实是：

```sql
where
  (
    sal > 1000
    and deptno = 20
  )
  or deptno = 30
```
上述括号中检索范围明显不是我们想要的条件。我们想要的是工资大于1000(前提)，之后才是部门编号是20或30，修改SQL为：
```sql
where
  sal > 1000 (
    and deptno = 20
    or deptno = 30
  )
```

tips: **当运算符的优先级不确定的时候，加()来保证**


## in
等同于 `or`，`in()`中的值不是区间，是具体的值，表示检索条件在 `in` 的几个值当中

```sql
sal in (1000,3000)

# 等同于
sal = 1000 or 
sal = 3000
```
## not in 
和 `in`是相反的操作，表示检索条件不在 `not in` 的几个值当中

```sql
sal not in (1000,3000)

# 等同于
sal != 1000 or 
sal != 3000
```

## like
`like` 这个查询条件是比较重要的，又称模糊查询。根据`like`的的位置，分为三种情况：
- `%XX`：后缀查询，以XX为后缀
- `XX%`：前缀查询，以XX为前缀
- `%XX%`：任意查询，任意位置包含XX


`like`通常和 `%`、`_`搭配使用：
- `%`：表示任意个字符
- `_`：表示任意1个字符

**示例一：**
```sql
# 找出名字中包含 A 的员工
select
  ename
from
  EMP
where
  ename like '%A%';
```
**示例二：**
```sql
# 找出名字中第二个字母为 A 的员工
select
  ename
from
  EMP
where
  ename like '_A%';
```
**示例三：** 下划线作为普通字符时，可以使用 `\`转义
```sql
# 找出名字中包含 _ 的员工
select
  ename
from
  EMP
where
  ename like '%\_%';
```
**示例四：**
```sql
# 找出名字中最后一个字母为 A 的员工
select
  ename
from
  EMP
where
  ename like '%A';
```