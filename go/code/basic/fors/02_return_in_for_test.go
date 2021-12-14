package fors

import (
	"fmt"
	"testing"
)

// TestReturnInFor 在 for 循环中执行 return
func TestReturnInFor(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println("i: ", i)
		if i == 3 {
			fmt.Println("3")
		} else {
			fmt.Println("return")
			// 直接return，不会执行后面的代码
			return
		}
	}

	fmt.Println("end")
}
