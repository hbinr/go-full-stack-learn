# gorm-gen使用
GitHub: https://github.com/go-gorm/gen

基于 GORM, 更安全更友好的ORM工具。
# 概览
- 自动生成 CRUD 和 DIY 方法
- 自动根据表结构生成模型（model）代码
- 完全兼容 GORM
- 更安全、更友好
- 多种生成代码模式
## 安装
安装
```sh
go get -u gorm.io/gen
```

在工程中导入引用 Gen:
```go
import "gorm.io/gen"
```
## 代码生成
### 命令行工具
通过二进制文件安装 gen 命令行工具:
```sh
go install gorm.io/gen/tools/gentool@latest
```

用法：
```sh
gentool -h
Usage of gentool:
  -db string
      input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html] (default "mysql")
  -dsn string
      consult[https://gorm.io/docs/connecting_to_the_database.html]
  -fieldNullable
      generate with pointer when field is nullable  当字段为空时，用指针生成
  -fieldWithIndexTag
      generate field with gorm index tag 用gorm索引标签生成字段
  -fieldWithTypeTag
      generate field with gorm column type tag  用gorm列类型标签生成字段
  -modelPkgName string  
      generated model code's package name  生成的模型代码的包名
  -outFile string
      query code file name, default: gen.go    查询代码文件名，默认：gen.go
  -outPath string
      specify a directory for output (default "./dao/query") 为输出指定一个目录（默认：./dao/query）
  -tables string
      enter the required data table or leave it blank 输入所需的数据表,多个表用 ',' 分隔，留空则生成所有表
  -withUnitTest
      generate unit test for query code  为查询代码生成单元测试
```

示例：
```sh
gentool -dsn "user:pwd@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local" -tables "orders,doctor"
```

### 通过代码生成器
通过`gen.NewGenerator` 
```go
// 指定数据类型,如: 若是int 返回 int64
var dataMap = map[string]func(detailType string) (dataType string){
	"int":  func(detailType string) (dataType string) { return "int64" },
	"json": func(string) string { return "json.RawMessage" },
}

func main() {
    // 创建gen代码生成器
	g := gen.NewGenerator(gen.Config{
        // 代码生成目录
		OutPath: "../../dal/query",
        // 代码生成模式. 生成默认的query代码
		Mode:    gen.WithDefaultQuery,

        // 是否生成单元测试
		WithUnitTest: true,
        // 当字段为空时，用指针生成
		FieldNullable:     true,
        // 用gorm索引标签生成字段
		FieldWithIndexTag: true,
	})
    // 使用初始化好的 DB 连接
	g.UseDB(*gorm.DB)

	g.WithDataTypeMap(dataMap)
    // 使用 JSON tag 
	g.WithJSONTagNameStrategy(func(c string) string { return "-" })

    // 指定需要应用的模型, eg: User 
	g.ApplyBasic(model.User{})
    // 所有模型都应用
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
```
## CURD

### 创建
### 查询
#### 单条数据查询
GROM 提供了 `First、Take、Last` 方法从数据库中查询单条数据，在查询数据库时会自动添加 `LIMIT 1` 条件，如果没有找到记录则返回错误 `ErrRecordNotFound。`

```go
u := query.Use(db).User

// Get the first record ordered by primary key
user, err := u.WithContext(ctx).First()
// SELECT * FROM users ORDER BY id LIMIT 1;

// Get one record, no specified order
user, err := u.WithContext(ctx).Take()
// SELECT * FROM users LIMIT 1;

// Get last record, ordered by primary key desc
user, err := u.WithContext(ctx).Last()
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

// check error ErrRecordNotFound
errors.Is(err, gorm.ErrRecordNotFound)
```

### 根据主键查询数据
```go
u := query.Use(db).User

user, err := u.WithContext(ctx).Where(u.ID.Eq(10)).First()
// SELECT * FROM users WHERE id = 10;

users, err := u.WithContext(ctx).Where(u.ID.In(1,2,3)).Find()
// SELECT * FROM users WHERE id IN (1,2,3);
```

注意，不能有如下写法：
```go
// 为了简化代码，加了 WithContext
u := query.Use(db).User.WithContext(ctx)

// 然后继续去 Where 查询. 实际操作中是不可行的
user, err := u.WithContext(ctx).Where(u.ID.Eq(10)).First()

```

`query.Use(db).User` 返回值是 `Query.User`，如下：
```go
type Query struct {
	db *gorm.DB

	User user
}
```

`query.Use(db).User.WithContext(ctx)` 返回值是 `*userDo`，如下：
```go
type userDo struct{ gen.DO }

// DO (data object): implement basic query methods ：DO实现基本的查询方法
// the structure embedded with a *gorm.DB, and has a element item "alias" will be used when used as a sub query
// 嵌入了*gorm.DB的结构，并且有一个元素项 "alias"，在作为子查询时将被使用
type DO struct {
	db        *gorm.DB
	alias     string // for subquery
	modelType reflect.Type
	schema    *schema.Schema
}
```