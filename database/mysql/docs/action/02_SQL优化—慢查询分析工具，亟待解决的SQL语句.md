# 慢查询分析工具，亟待解决的SQL语句



## 慢查询分析工具--MySQL自带命令 `mysqldumpslow.pl`
使用MySQL自带的工具 `mysqldumpslow.pl`

我们想要知道以下关键信息：
- 慢SQL排序，时间最长的显示在最上面
- 指定如果慢SQL非常多，显示top n的SQL


`mysqldumpslow.pl` 在MySQL的bin目录下，要使用该工具，需要先安装好 perl 脚本语言

使用示例：
```sql
perl mysqldumpslow.pl -s t  -t 5 

```

参数解析：
- `perl`: 使用perl脚本语言运行该工具
- `mysqldumpslow.pl`: MySQL自带慢查询分析工具
- `-s t`: `-s` 表示排序，后面指定排序标准， `t` 表示SQL执行总时间。列举一些排序标准参数
  - `t`: 表示SQL执行总时间
  - `c`: 表示SQL执行总次数
  - `l`: 表示锁的时间
- `-t 5`： `-t` 表示top， `5` 表示5条。 该命令表示前5条



## 慢查询分析工具--pt-query-digest

[MySQL慢查询（二） - pt-query-digest详解慢查询日志](https://www.cnblogs.com/luyucheng/p/6265873.html)

