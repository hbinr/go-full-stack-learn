package interfaces

import "testing"

/*
	inteface{}与*interface{}


*/
type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}
func TestInterface(t *testing.T) {
	// ABCD中哪一行存在错误？

	s := S{}
	p := &s
	f(s) // A
	// g(s) // B
	f(p) // C
	// g(p) // D

	var ptr *interface{} // E 可以通过编译
	g(ptr)
	// 看到这道题需要第一时间想到的是Golang是强类型语言，interface是所有go类型的父类函数中func f(x interface{})
	// 的interface{}可以支持传入golang的任何类型，包括指针

	// 但是函数func g(x *interface{})只能接受*interface{}

}
