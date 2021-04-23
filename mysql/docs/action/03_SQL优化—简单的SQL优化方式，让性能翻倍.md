# 简单的SQL优化方式，让性能翻倍

## 首先定位SQL慢的根本原因

一般SQL执行很慢的原因有：
- 多表关联查询
- 子查询
- 查询条件未设置索引
- 查询条件设置了索引，但是并未走索引 

给定一个SQL:
```sql
select 
    a.date_str,
    a.shopCode,
    a.add_car_pv,
    (
        select
        b.shop_type
        from
        dp_shop b
        where
        a.shopCode = b.shop_code
    )
    b.shop_type
    from 
    dp_car_copy a
    order by
        a.shopCode,
        a.add_car_pv,
        a.date_str
```
优化步骤:
1. 先看两个表的数据量有多少
   1. `select count(1) from dp_car_copy` 和 `select count(1) from dp_shop`
2. 分析这条SQL，关联多个表+子查询
3. 尝试将关联子查询改为关联查询

改为关联查询修改后的SQL:
```sql
select 
    a.date_str,
    a.shopCode,
    a.add_car_pv,
    b.shop_type
    from 
        dp_car_copy a,
        dp_shop b
    where
        a.shopCode = b.shop_code
    order by
        a.shopCode,
        a.add_car_pv,
        a.date_str
```

**知识补充：**

子查询就是查询中又嵌套的查询,表连接都可以用子查询，但不是所有子查询都能用表连接替换，子查询比较灵活，方便，形式多样，适合用于作为查询的筛选条件。

表连接更适合与查看多表的数据。子查询不一定需要两个表有关联字段，而连接查询必须有字段关联（所谓的主外键关系）
1. 表关联的效率要高于子查询，因为子查询走的是笛卡尔积
2. 表关联可能有多条记录，子查询只有一条记录，如果需要唯一的列，最好走子查询对于数据量多的肯定是用连接查询快些
   
**原因：**

因为子查询会多次遍历所有的数据（视你的子查询的层次而定），而连接查询只会遍历一次。

**场景：**
- 数据量：也就无所谓是连接查询还是子查询，视自己的习惯而定。
- 数据量大：优选使用关联查询from t1，t2

## 常见的优化方式
- SQL本身优化
- 反范式设计优化
- 索引优化
- 服务器硬件
### SQL本身优化

### 反范式设计优化

### 索引优化

### 服务器硬件
- 硬盘： 机械硬盘很慢，改成固态硬盘，性能会提升10倍以上
- CPU
- 内存

### MySQL服务器优化
服务器最好放在Linux