package gotos

import (
	"fmt"
	"testing"
)

/*
	1. Go 语言的 goto 语句可以无条件地转移到程序中指定的行。
	2. goto 语句通常与条件语句配合使用。可用来实现条件转移，跳出循环体等功能。
	3. 在 Go 程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序
*/
func PrintHello() {
	fmt.Println("Hello")
}

// TestGotoUse
func TestGotoUse(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i%6 == 0 {
			goto Label
		}
	}
Label:
	fmt.Println("world")
	go PrintHello()
}
