package structs

/*
	如果函数外部没有引用，则优先放到栈中，即不会逃逸
	如果函数外部存在引用，则必定放到堆(heap)中，则必定发生逃逸
	cd structs

	使用命令 go build -gcflags=-m  可以看到是否发生内存逃逸。注意：针对非 XX_test.go 文件才可使用该命令
*/
type User struct {
	name string
}

// StructPtrMemoryEscape 内存逃逸之结构体指针逃逸
func StructPtrMemoryEscape(name string) *User {
	u := new(User)
	u.name = name
	return u
}

// CallStructPtrMemoryEscape 外部调用StructPtrMemoryEscape
func CallStructPtrMemoryEscape() {
	StructPtrMemoryEscape("Bob")
}

// StructPtrNoMemoryEscape 没有逃逸的示例
func StructPtrNoMemoryEscape(name string) {
	u := new(User)
	u.name = name // 没有被任何外部函数调用，只在本函数内部
}
