# 初始使用json-iterator

json-iterator 用法100%兼容go标准库，并且性能非常高，且不需要生成额外的代码。
## 安装
```go
go get github.com/json-iterator/go
```

## 使用

```go
type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}
group := ColorGroup{
	ID:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}
b, err := jsoniter.Marshal(group)

```