# sqlx 库使用

在项目中我们通常可能会使用 database/sql 连接 MySQL 数据库。本文借助使用 sqlx 实现批量插入数据的例子，介绍了 sqlx 中可能被你忽视了的 `sqlx.In` 和 `DB.NamedExec` 方法。

在项目中我们通常可能会使用 database/sql 连接 MySQL 数据库。sqlx 可以认为是 Go 语言内置 database/sql 的超集，它在优秀的内置 database/sql 基础上提供了一组扩展。这些扩展中除了大家常用来查询的

- `Get(dest interface{}, ...) error`
- `Select(dest interface{}, ...) error`

外还有很多其他强大的功能。

## 安装 sqlx

```go
go get github.com/jmoiron/sqlx
```

[文章详细代码](https://github.com/hblcok/go-full-stack-learn/tree/master/mysql/code/sqlx_demo)

## 基本使用

### 连接数据库

记得匿名导入 mysql 驱动 `_ "github.com/go-sql-driver/mysql"`

```go
package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB 数据库句柄
var DB *sqlx.DB

// InitMysql 初始化MySQL连接
func InitMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gateway?charset=utf8mb4&parseTime=True"
	// sqlx.Connect() 底层做了Open和Ping
	if DB, err = sqlx.Connect("mysql", dsn); err != nil {
		fmt.Println("sqlx.Connect failed, err:", err)
		return
	}
	// 设置MySQL相关配置
	DB.SetConnMaxLifetime(20)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
	fmt.Println("mysql connect success......")
	return
}
```

### 查询

#### 查询单行数据示例代码如下：

```go
// Demo Demo结构体
type Demo struct {
	ID   int
	Age  int
	Name string
}

// queryRowDemo 查询单条数据示例
func queryRowDemo() {
	sqlStr := `select id,age,name from demo where id = ?`
	var d Demo
	// Get()返回单条数据
	if err := db.DB.Get(&d, sqlStr, 1); err != nil {
		fmt.Println("db.DB.Get failed, err:", err)
		return
	}
	fmt.Println("queryRowDemo success, data:", d)
}

```

#### 查询多行数据示例代码如下：

```go
// queryMultiRowDemo 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := `select id,age,name from demo where id > ?`
	var ds []Demo
	// Select()返回多条数据
	if err := db.DB.Select(&ds, sqlStr, 1); err != nil {
		fmt.Println("db.DB.Select failed, err:", err)
		return
    }

	fmt.Println("queryMultiRowDemo success, data:", ds)
}
```

### 增、删、改

sqlx 中的 exec 方法与原生 sql 中的 exec 使用基本一致

```go
// insertRowDemo 新增一条数据
func insertRowDemo() {
	sqlStr := `insert into demo(age,name) values(?,?)`
	res, err := db.DB.Exec(sqlStr, 20, "Bob")
	if err != nil {
		fmt.Println("insert demo's data failed, err:", err)
		return
	}
	lstID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("res.LastInsertId() failed, err:", err)
		return
	}
	fmt.Printf("insert demo's data success,lstID:%d", lstID)
}

// updateRowDemo 更行数据
func updateRowDemo() {
	sqlStr := `update demo set age = ? where id = ?`
	res, err := db.DB.Exec(sqlStr, 25, 3)
	if err != nil {
		fmt.Println("update demo's data failed, err:", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("update res.RowsAffected error, err:", err)
		return
	}
	fmt.Println("update demo's data success,rowsAffected:", n)
}

// deleteRowDemo 删除数据
func deleteRowDemo() {
	sqlStr := `delete from demo where id = ?`
	res, err := db.DB.Exec(sqlStr, 4)
	if err != nil {
		fmt.Println("delete demo's data failed,err:", err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println("delete res.RowsAffected error, err:", err)
		return
	}
	fmt.Println("delete demo's data success,rowsAffected:", n)

}
```

### NamedExec 操作 sql，针对增删改

`DB.NamedExec`方法用来绑定 SQL 语句与**结构体**或 **map** 中的同名字段

```go
// insertRowByNameExecDemo 使用NamedExec()操作sql
func insertRowByNameExecDemo() {
	sqlStr := `insert into demo(age,name) values(:age,:name)`
	// 使用结构体
	ronger := &Demo{Age: 19, Name: "ronger"}

	res, err := db.DB.NamedExec(sqlStr, ronger)
	if err != nil {
		fmt.Println("insert demo's data failed,err:", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("res.LastInsertId() failed, err:", err)
		return
	}
	fmt.Println("insert demo's data success, 5id:", id)
}
```

**特别注意：**

- sql 变化：values()中需要指明字段名，不再是占位符 `?`
- 指明的字段名的顺序和结构体中字段的顺序要保持一致，否则会匹配错误。如 ` values(:age,:name)` 对应的结构体应该是 `&Demo{Age: 19, Name: "ronger"}`，不是 `&Demo{Name: "ronger",Age: 19,}`

### NamedQuery 操作 sql，针对查

与 `DB.NamedExec` 同理，这里是支持查询。

```go
// queryRowByNamedQueryDemo 使用NamedExec()操作sql，针对查
func queryRowByNamedQueryDemo() {
	sqlStr := `SELECT * from demo WHERE name=:name`
	demo := Demo{
		Name: "Bob",
	}
	rows, err := db.DB.NamedQuery(sqlStr, &demo)
	if err != nil {
		fmt.Println("select demo's data by NamedQuery failed,err:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		// 方式一：使用StructScan，一次扫描一行
		var d Demo
		err := rows.StructScan(&d)
		if err != nil {
			fmt.Println("rows.StructScan failed, err:", err)
			continue
		}
		fmt.Println("row's data :", d)

		// 方式二：使用SliceScan，一次扫描多行
		res, err := rows.SliceScan()
		if err != nil {
			fmt.Println("rows.SliceScan failed,err:", err)
			return
		}
		fmt.Println("row's res :", res)
	}
}
```

#### rows.StructScan()

一次扫描一行数据，返回的结果是单个对象的数据，推荐场景：只需要单条数据

#### rows.SliceScan()

一次扫描多行数据，底层是封装了 for 循环 scan 每一行数据。推荐场景：需要多条数据，直接返回的就是多条数据，不需要`append`每一条数据到切片中

### 事务

我们可以使用 sqlx 中提供的 `db.Beginx()`和 `tx.Exec()`方法，和原生 `database/sql` 库操作基本一致。示例代码如下

```go
// transactionDemo 使用sqlx进行事务操作
func transactionDemo() (err error) {
	var (
		tx *sql.Tx
		rs sql.Result
		n  int64
	)
	if tx, err = db.DB.Begin(); err != nil {
		fmt.Println("db.DB.Begin() failed, err:", err)
		return err
	}
	// defer 来进行事务回滚
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			fmt.Println("rollback")
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
			fmt.Println("commit")
		}
	}()

	sqlStr1 := "Update demo set age=30 where id=?"
	rs, err = tx.Exec(sqlStr1, 1)
	if err != nil {
		fmt.Println("tx.Exec sqlStr1 failed,err:", err)
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		fmt.Println("rs.RowsAffected sqlStr1 failed,err:", err)
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}

	sqlStr2 := "Update demo set age=50 where id=?"
	rs, err = tx.Exec(sqlStr2, 5)
	if err != nil {
		fmt.Println("tx.Exec sqlStr2 failed,err:", err)
		return err
	}
	n, err = rs.RowsAffected()
	if err != nil {
		fmt.Println("rs.RowsAffected RowsAffected failed,err:", err)
		return err
	}
	if n != 1 {
		return errors.New("exec sqlStr1 failed")
	}
	return err

}
```

这里需要注意的是，因为 defer 函数里使用的 err 是函数返回时便定义了的变量，所以在整个 `transactionDemo()`内，err 不能再通过 `:=` 来重声明，否则便判断失败，事务失效

## sqlx.In()

### sqlx.In 的批量插入示例

#### 定义结构体

定义一个 Demo 结构体，字段通过 tag 与数据库中 demo 表的列一致。

```go
type Demo struct {
	ID   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}
```

#### bindvars（绑定变量）

查询占位符?在内部称为 `bindvars` （查询占位符）,它非常重要。你应该始终使用它们向数据库发送值，因为它们可以防止 SQL 注入攻击。database/sql 不尝试对查询文本进行任何验证；它与编码的参数一起按原样发送到服务器。除非驱动程序实现一个特殊的接口，否则在执行之前，查询是在服务器上准备的。因此 `bindvars` 是特定于数据库的:

- MySQL 中使用 ?
- PostgreSQL 使用枚举的 $1、$2 等 bindvar 语法
- SQLite 中 ? 和 \$1 的语法都支持
- Oracle 中使用 :name 的语法

`bindvars` 的一个常见误解是，它们用来在 sql 语句中插入值。它们其实仅用于参数化，不允许更改 SQL 语句的结构。例如，使用 `bindvars` 尝试参数化列或表名将不起作用：

```go
// ？不能用来插入表名（做SQL语句中表名的占位符）
db.Query("SELECT * FROM ?", "mytable")

// ？也不能用来插入列名（做SQL语句中列名的占位符）
db.Query("SELECT ?, ? FROM people", "name", "location")
```

#### 使用 sqlx.In 实现批量插入

**前提**是需要我们的结构体实现 driver.Valuer 接口：

```go
// Value 使用sqlx.In 必须实现的接口
func (d Demo) Value() (driver.Value, error) {
	return []interface{}{d.Age, d.Name}, nil
}
```

使用 `sqlx.In` 帮我们拼接语句和参数,

```go
// BatchInsertDemos 批量新增数据
func BatchInsertDemos(ds []interface{}) error {
	sqlStr := `INSERT  INTO  demo(age,name) values(?),(?),(?)` // 多个一个占位符
	query, args, err := sqlx.In(sqlStr, ds...)                 // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开ds
	if err != nil {
		fmt.Println("sqlx.In insert demo's data failed,err:", err)
		return err
	}
	fmt.Println("查看生成的querystring:", query)
	fmt.Println("查看生成的args", args)

	_, err = db.DB.Exec(query, args...)
	if err != nil {
		fmt.Println("db.DB.Exec insert demo's data failed,err:", err)
		return err
	}
	return nil
}
func main() {
	if err := db.InitMysql(); err != nil {
		fmt.Println("sqlx.InitMysql failed,err:", err)
		return
	}
	defer db.DB.Close()
	d1 := Demo{Age: 18, Name: "小明"}
	d2 := Demo{Age: 28, Name: "小明2"}
	d3 := Demo{Age: 38, Name: "小明3"}
	ds := []interface{}{&d1, &d2, &d3}
	BatchInsertDemos(ds)
}

```

**注意：**

- sql 的写法要注意：values 列是单独包括一个占位符，并始终比字段列多一列
- BatchInsertDemos**传入的参数**是`[]interface{}`，是空接口，一定要注意
- 实现的 Value 方法、sql、 Demo 对象构造，他们的字段顺序要完全保持一致，否则数据匹配错误

#### 使用 NamedExec 实现批量插入

**注意：**该功能目前有人已经推了[#285 PR](https://github.com/jmoiron/sqlx/pull/285)，但是作者还没有发 release，所以想要使用下面的方法实现批量插入需要暂时使用 master 分支的代码：

在项目目录下执行以下命令下载并使用 master 分支代码：

```
go get github.com/jmoiron/sqlx@master
```

使用 NamedExec 实现批量插入的代码如下：

```go
// BatchInsertDemos2 NamedExec批量新增数据
func BatchInsertDemos2(demos []*Demo) error {
	_, err := db.DB.NamedExec("INSERT INTO demo (age, name) VALUES (:age, :name)", demos)
	return err
}

demos2 := []*Demo{&d1, &d2, &d3}
BatchInsertDemos2(demos2)
```

这种方法不需要实现 driver.Valuer 接口

在 `main`方法中测试上述两种方法：

```go
func main() {
	if err := db.InitMysql(); err != nil {
		fmt.Println("sqlx.InitMysql failed,err:", err)
		return
	}
	defer db.DB.Close()
	d1 := Demo{Age: 18, Name: "小明"}
	d2 := Demo{Age: 28, Name: "小明2"}
	d3 := Demo{Age: 38, Name: "小明3"}
	demos1 := []interface{}{&d1, &d2, &d3}
	BatchInsertDemos(demos1)

	demos2 := []*Demo{&d1, &d2, &d3}
	BatchInsertDemos2(demos2)
}
```

### sqlx.In 的查询示例

关于 sqlx.In 这里再补充一个用法，在 sqlx 查询语句中实现 `in` 查询和 `FIND_IN_SET` 函数。

即实现 `SELECT _ FROM demo WHERE id in (3, 2, 1);`

和 `SELECT _ FROM demo WHERE id in (3, 2, 1) ORDER BY FIND_IN_SET(id, '3,2,1');`。

#### in 查询

查询 id 在给定 id 集合中的数据。

```go
// QueryByIDs 根据给定ID查询
func QueryByIDs(ids []int) (demos []Demo, err error) {
	query, args, err := sqlx.In(`SELECT * FROM demo WHERE id IN (?)`, ids)
	if err != nil {
		fmt.Println("QueryByIDs select demo's data failed,err:", err)
		return
	}
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它
	query = db.DB.Rebind(query)


	fmt.Println("查询的sql为：", query) //查询的sql为： SELECT * FROM demo WHERE id IN (?, ?, ?, ?)
	fmt.Println("拼接的参数为：", args)  // 拼接的参数为： [1 2 3 5]
	// Select() 查询多条数据
	err = db.DB.Select(&demos, query, args...)
	return
}

// main.go中调用：
QueryByIDs([]int{1,2,3,5})
```

#### in 查询和 FIND_IN_SET 函数

查询 id 在给定 id 集合的数据并维持给定 id 集合的顺序。

```go
// QueryAndOrderByIDs 按照指定id查询并维护顺序
func QueryAndOrderByIDs(ids []int) (demos []Demo, err error) {
	// 将int型切片转成string切片参数，用作sql参数
	strIDs := make([]string, 0, len(ids))
	for _, v := range ids {
		strIDs = append(strIDs, strconv.Itoa(v))
	}

	// 拼接sql，一定要加strIDs用 ','隔开
	query, args, err := sqlx.In(`SELECT * FROM demo WHERE id IN (?) ORDER BY FIND_IN_SET(id,?)`,
		ids, strings.Join(strIDs, ","))
	if err != nil {
		fmt.Println("QueryAndOrderByIDs select demo's data failed,err:", err)
		return
	}
	query = db.DB.Rebind(query)
	err = db.DB.Select(&demos, query, args...)
	return
}
```

当然，在这个例子里面你也可以先使用 IN 查询，然后通过代码按给定的 ids 对查询结果进行排序。

参考：

https://www.liwenzhou.com/posts/Go/sqlx/
