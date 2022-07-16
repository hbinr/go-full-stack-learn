# 基于 go-zero 练习 PostgreSQL 开发

## 创建表

## 查看相关 goctl 命令

### goctl model -h 
```sh
> goctl model -h
Generate model code

Usage:
  goctl model [command]

Available Commands:
  mongo       Generate mongo model
  mysql       Generate mysql model
  pg          Generate postgresql model

Flags:
  -h, --help   help for model

Use "goctl model [command] --help" for more information about a command.
```

### goctl model pg -h
`go-zero`支持3种数据库：mysql、postgresql、mongo。

我们练习 `postgresql`，因此使用 `goctl model pg -h` 命令查看帮助信息

```sh
> goctl model pg -h
Generate postgresql model

Usage:
  goctl model pg [flags]
  goctl model pg [command]

Available Commands:
  datasource  Generate model from datasource

Flags:
  -h, --help   help for pg

Use "goctl model pg [command] --help" for more information about a command.
```


## 使用 goctl 命令生成相关CURD代码


```sh
# cd 到工作目录 (自定义)
cd  /Usr/Workspace/go/src/my-project/go-full-stack-learn/database/postgresql/code/go-zero-demo

# 生成代码
goctl model pg datasource -url="user:password@tcp(127.0.0.1:3306)/database" -table="*"  -dir="./model"
```
