package datatype

import (
	"fmt"
	"testing"
)

//TestByte 字符显示定义和自动类型推导定义是不同的数据类型
func TestByte(t *testing.T) {
	// 显示定义
	var ch1 byte
	ch1 = 'a'
	fmt.Printf("ch1 的类型为：%T，值为：%c\n", ch1, ch1) // ch1 的类型为：uint8，值为：a
	// 自动类型推导则使用 int32 类型
	ch2 := 'b'
	fmt.Printf("ch2 的类型为：%T，值为：%c\n", ch2, ch2) // ch1 的类型为：int32，值为：a
}

// TestString  多行字符串，换行时，+ 必须在当前行的结尾
func TestString(t *testing.T) {
	var str = "hello " +
		"world"
	str += "!"
	fmt.Println(str)

}
