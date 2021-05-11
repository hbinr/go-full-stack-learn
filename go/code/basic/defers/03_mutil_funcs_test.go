package defers

import (
	"fmt"
	"testing"
)

/*
	defer 面试问题
*/
func TestMutilFuncs(t *testing.T) {
	fmt.Println("d1 :", d1(1)) // 2
	fmt.Println("d2 :", d2(1)) // 1
	fmt.Println("d3 :", d3(1)) // 2
}
func d1(i int) (v int) {
	v = i
	defer func() {
		v += i
	}()
	return v
	/*
		执行流程：
		1. v = i     -> v = 1
		2. defer func()  -> v = v + i = 1 + 1 = 2
		3. return v  -> v = 2
	*/
}

func d2(i int) int {
	v := i
	fmt.Println("v out: ", v) // 1
	defer func() {
		v += i
		fmt.Println("v in: ", v) // 2

	}()
	return v
	/*
		执行流程：
		1. v := i     -> v := i = 1
		2. anony = v = 1  -> 因为函数返回的定义为匿名返回值，go底层会默认生成一个返回变量(假如为anony)，然后将v的值复制给anony
		3. defer func()  -> v = v + i = 1 + 1 = 2  此时 v 的值为 2
		4. return anony  -> anony = 1
	*/
}

func d3(i int) (v int) {
	v = i
	defer func() {
		v += i
	}()
	return i

	/*
		执行流程：
		1. v = i             -> v = 1
		2. return i 第一步   -> 局部变量 i 入栈，i = 1
		3. return i 第二步   -> 函数返回的定义为具名返回值，所以将i复制给v， 即 v = i = 1
		4. defer func()     -> v = v + i = 1 + 1 = 2
		5. return i 第三步   -> 其实就是return v， 因为在第3步已经将返回值确定为具名返回值v了
	*/
}
