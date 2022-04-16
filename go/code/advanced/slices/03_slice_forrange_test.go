package slices

import (
	"fmt"
	"testing"
)

type user struct {
	name string
	age  uint64
}

func TestForRange(t *testing.T) {
	u := []user{
		{"hbinr", 23},
		{"tom", 19},
		{"bob", 18},
	}
	n := make([]*user, 0, len(u))

	// for-range 表达式中，迭代变量(比如实例中的v)只会声明一次，在多次迭代中共享
	for _, v := range u {
		// v内存地址始终未变
		fmt.Println("v", &v)
		// fmt.Printf("v: %+v \n", &v)

		// 由于内存共享，所以最后一次迭代的值赋给了迭代变量初始化(声明)的地址
		n = append(n, &v)

		// 正确：
		// tmp := v
		// n = append(n, &tmp)
	}
	fmt.Println("->", n)

	for _, v := range n {
		fmt.Println("range: ", v)
	}
}
