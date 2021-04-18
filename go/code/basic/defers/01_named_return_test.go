package defers

import (
	"fmt"
	"testing"
)

// 命名返回值 测试
func TestNamedReturn(t *testing.T) {
	fmt.Println("NamedReturn1: ", NamedReturn1())
	fmt.Println("NamedReturn2: ", NamedReturn2())
	fmt.Println("NamedReturn3: ", NamedReturn3())
}

func NamedReturn1() (r int) {
	i := 1
	defer func() {
		i = i + 1
	}()
	return i
}
func NamedReturn2() (r int) {
	defer func(r int) {
		r = r + 2
	}(r)

	return 2
}

func NamedReturn3() (r int) {
	defer func(r *int) {
		*r = *r + 2
	}(&r)

	return 2
}
