package defers

import (
	"errors"
	"fmt"
	"testing"
)

// 多个函数调用，多返回值
func TestMutilFuncs(t *testing.T) {
	e1()
	e2()
	e3()
}
func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("e1 defer err")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("e2 defer err")
}

func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("e3 defer err")
}
