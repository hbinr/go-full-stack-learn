package structs

import (
	"fmt"
	"testing"
)

type User struct {
	name string
}

/*
	通过指针变量 p 访问其成员变量 name ，下面语法正确的是()
	A. p.name
	B. (*p).name
	C. (&p).name
	D. p->name  // C语言的写法

	参考答案：AB

	解析:
	指针变量也是变量，变量、指针和地址三者的关系是：每一个变量都有自己地址，每一个指针的值就是地址

	& 返回变量的内存地址，也叫取址符。

	* 返回变量的值, 其只能作用在指针上

	重点：一个指针的值是另一个变量的地址。一个指针对应变量在内存中的存储位置。并不是每一个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址。

*/
func TestPtrStruct(t *testing.T) {
	p := &User{
		name: "tom",
	}
	fmt.Println(p.name)
	fmt.Println((*p).name) // () 相当于包了一层，再加上 *，是一个指针对象。
	// fmt.Println((&p).name)  & 取址操作，取出地址，再 . 来获取字段是错误的

	fmt.Printf("%T\n", p)    // *structs.User
	fmt.Printf("%T\n", (*p)) // structs.User, 加了*返回 p 的值，那么p指向的对象不再是指针类型，而是一个普通的值类型
	fmt.Printf("%T\n", *p)   // structs.User, 同上
	fmt.Printf("%T\n", &p)   // **structs.User
	fmt.Printf("%T\n", (&p)) // **structs.User
}
