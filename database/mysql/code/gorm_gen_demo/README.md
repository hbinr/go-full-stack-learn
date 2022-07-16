# examples

## 使用命令行工具生成代码

```sh
gentool -dsn "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local" -tables "user" -outPath "./data/dao" -modelPkgName "entity" 
```