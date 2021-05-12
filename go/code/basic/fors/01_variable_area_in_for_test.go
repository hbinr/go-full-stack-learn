package fors

import (
	"fmt"
	"testing"
)

func TestVariableAreaInFor(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println("i: ", i)
	}
	fmt.Println("--------------")
	// fmt.Println("i: ", i)  // 编译报错 undefined: i，这才是正常逻辑
	// {
	// s := "hello" // 编译报错 s declared but not used，这才是正常逻辑
	// }
	// fmt.Println("s: ", s)  // 编译报错 undefined: i，这才是正常逻辑
}
