# shard 特性用法

将单击表变为分片的表，增加读写性能

## 为 DB 激活特性

```sh
sh.enableSharding('mydb")
```

## 配置 hash 分片

```
sh.shardCollection("my_db.my_collection",{_id:"hashed"}, false,{numlnitialChunks: 10240})
```

对数据库 my_db 中的表 my_collection 配置 hash 分片，其中以 `_id` 建 hash 索引，预分配了 10240 个 chunk

分配原理： hash(\_id)/10240 的模等于多少就被分配到对应的 chunk 上

**注意**：

- 如果是按非 shard key 查询，请求被扇出给所有 shard，即会去查询所有数据库分片，性能就会差很多。所以查询时，按照 shard key 查询能大大提高性能
