package defers

import (
	"errors"
	"fmt"
	"testing"
)

/*
	defer error应用
*/
func TestMutilErr(t *testing.T) {
	e1()
	e2()
	e3()
}
func e1() {
	var err error
	defer fmt.Println(err) // 值传递，参数在defer的时候就已经确定了
	err = errors.New("e1 defer err")
}

func e2() {
	var err error
	defer func() {
		fmt.Println(err) // e2 defer err
	}()
	err = errors.New("e2 defer err")

	/*
		执行流程：
		err = nil
		err = errors.New("e2 defer err")
		defer func()执行
	*/
}

func e3() {
	var err error
	defer func(err error) { // 值传递，参数在defer的时候就已经确定了
		fmt.Println(err) // nil
	}(err)
	err = errors.New("e3 defer err")
}
