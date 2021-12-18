package interfaces

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test(t *testing.T) {
	var data *byte
	var in interface{}
	fmt.Println(data, data == nil) // <nil> true
	fmt.Println(in, in == nil)     // <nil> true
	in = data

	fmt.Println(in, in == nil) // <nil> false
	// 获取 in 的真实值
	testIn := (*struct {
		data  uintptr
		_type uintptr
	})(unsafe.Pointer(&in))
	// in 是空接口类型，通过 in = data 这行代码，data指针有了值，但是 _type指针还是空，所以 in 此时已经 != nil了
	fmt.Printf("%+v:", testIn) // &{data:17751936 _type:0}
}
