# dapr 绑定 MySQL 组件
## MySQL
创建 MySQL 组件，需要绑定类型为bindings.mysql 的组件 

Dapr 底层使用 [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) 来处理MySQL相关操作
## 定义 MySQL 组件配置
```yaml
apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: dapr-admin-mysql       # 自定义
  namespace: dapr-admin-mysql  # 自定义
spec:
  type: bindings.mysql
  version: v1
  metadata:
    - name: url # Required, define DB connection in DSN format
      value: "root:123456@tcp(127.0.0.1:3306)/study?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"
    - name: maxIdleConns
      value: "10"
    - name: maxOpenConns
      value: "10"
    - name: connMaxLifetime
      value: "12s"
    - name: connMaxIdleTime
      value: "12s"
```

> **注意**

> 上面的例子MySQL DSN直接显示了(demo)，建议使用secret store进行加密，[具体请看](https://docs.dapr.io/operations/components/component-secrets/)。




## 绑定的 MySQL 支持的API
此组件支持与以下操作的输出绑定：
- `exec`
- `query`
- `close`

### exec

`exec` 操作可用于DDL操作（如表创建），以及插入，更新，删除仅返回元数据（例如受影响行的数量）
#### 请求示例
```json
{
  "operation": "exec",
  "metadata": {
    "sql": "INSERT INTO foo (id, c1, ts) VALUES (1, 'demo', '2020-09-24T11:45:05Z07:00')"
  }
}
```
#### 响应示例
```json
{
  "metadata": {
    "operation": "exec",
    "duration": "294µs",
    "start-time": "2020-09-24T11:13:46.405097Z",
    "end-time": "2020-09-24T11:13:46.414519Z",
    "rows-affected": "1",
    "sql": "INSERT INTO foo (id, c1, ts) VALUES (1, 'demo', '2020-09-24T11:45:05Z07:00')"
  }
}
```

### query
`query` 操作用于SELECT语句，其以行值数组的形式返回元数据以及数据。
#### 请求示例
```json
{
  "operation": "query",
  "metadata": {
    "sql": "SELECT * FROM foo WHERE id < 3"
  }
}
```
#### 响应示例
```json
{
  "metadata": {
    "operation": "query",
    "duration": "432µs",
    "start-time": "2020-09-24T11:13:46.405097Z",
    "end-time": "2020-09-24T11:13:46.420566Z",
    "sql": "SELECT * FROM foo WHERE id < 3"
  },
  "data": "[
    [0,\"test-0\",\"2020-09-24T04:13:46Z\"],
    [1,\"test-1\",\"2020-09-24T04:13:46Z\"],
    [2,\"test-2\",\"2020-09-24T04:13:46Z\"]
  ]"
}
```

### close
可以使用 `close` 操作来显式关闭DB连接并将其返回到池中。 此操作没有任何响应
#### 请求示例
```json
{
  "operation": "close"
}
```
> 注意

> MySQL binding 本身不会阻止SQL注入，如与任何数据库应用程序一起验证在执行查询之前的输入。

> 推荐使用 '?' 占位符来解决SQL注入问题

**参考：**
- https://docs.dapr.io/reference/components-reference/supported-bindings/mysql/