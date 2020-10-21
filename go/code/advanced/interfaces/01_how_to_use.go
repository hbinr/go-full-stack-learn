package interfaces

import (
	"fmt"
	"testing"
)

// TestAppend int 转 int64/int32/intXX...的问题
func TestDataConvert(t *testing.T) {
	var tmp int = 10 // 当前系统为64为系统，底层会使用int64类型
	fmt.Printf("tmp 的数据类型：%T\n", tmp)
	tmp32 := int32(tmp)
	fmt.Printf("tmp32 的数据类型：%T\n", tmp32)
	tmp64 := int64(tmp)
	fmt.Printf("tmp64 的数据类型：%T\n", tmp64)

	var tmp2 int64 = 20
	fmt.Println(tmp + int(tmp2))

}
