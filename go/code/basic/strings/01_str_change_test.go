package strings

import (
	"fmt"
	"testing"
)

func TestStrChange(t *testing.T) {
	str := "hello"
	fmt.Println("str[0]", str[0]) // str[0] 104  输出的是h的 ASCII 值

	fmt.Printf("str[0] %s\n", string(str[0])) // str[0] h  输出的是h的 ASCII 值
	// str[0] = "H"                              // 编译报错：cannot assign to str[0]
}
