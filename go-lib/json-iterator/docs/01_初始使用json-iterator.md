# 初始使用json-iterator

json-iterator 用法100%兼容go标准库，并且性能非常高，且不需要生成额外的代码。
## 安装
```go
go get github.com/json-iterator/go
```

## 使用
用法和标准库用法一致：
```go
package code

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
)

const jsonByte = `{"id":1,"Name":"Reds","Colors":["Crimson", "Red", "Ruby", "Maroon"]}`

func TestDemo(t *testing.T) {
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
	t.Logf("bytes %v", b)

	if err != nil {
		t.Error(err)
	}
	t.Log("------------split-----------")
	if err = jsoniter.Unmarshal([]byte(jsonByte), &group); err != nil {
		t.Error(err)
	}
	t.Logf("group %v", group)
}
```