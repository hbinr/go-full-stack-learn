# MySQL 预处理

## 什么是预处理？

**普通 SQL 语句执行过程：**

1. 客户端对 SQL 语句进行占位符替换得到完整的 SQL 语句。
2. 客户端发送完整 SQL 语句到 MySQL 服务端
3. MySQL 服务端执行完整的 SQL 语句并将结果返回给客户端。

**预处理执行过程：**

1. 把 SQL 语句分成两部分，命令部分与数据部分。
2. 先把命令部分发送给 MySQL 服务端，MySQL 服务端进行 SQL 预处理。
3. 然后把数据部分发送给 MySQL 服务端，MySQL 服务端对 SQL 语句进行占位符替换。
4. MySQL 服务端执行完整的 SQL 语句并将结果返回给客户端。

## 为什么要预处理？

1. 优化 MySQL 服务器重复执行 SQL 的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
2. 避免 SQL 注入问题。

## 什么情况下进行 MySQL 预处理？

- 批量执行同一个 SQL，只是替换占位符的数据不同。

## Go 实现 MySQL 预处理

`database/sql` 库中使用下面的 Prepare 方法来实现预处理操作。

```go
func (db *DB) Prepare(query string) (*Stmt, error
```

`Prepare`方法会先将 sql 语句发送给 MySQL 服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。

查询操作的预处理示例代码如下：

```go
func prepareQueryDemo() {

	sqlStr := "select * from demo where id > ? "
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("DB.Prepare failed, err:", err)
		return
	}
	// 非常重要：释放命令部分资源，当MySQL连接数被占满，资源不合理释放
	// 会导致sql一直等待可利用资源却资源沾满而无法执行
	defer stmt.Close()

	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Println("stmt.Query failed, err:", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var d Demo
		err := rows.Scan(&d.ID, &d.Age, &d.Name)
		if err != nil {
			fmt.Println("rows.Scan failed,err:", err)
			return
		}
		fmt.Println("Demo's data:", d)
	}
}
```

参考：
https://www.liwenzhou.com/posts/Go/go_mysql/#autoid-1-1-2
