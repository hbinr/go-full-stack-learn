package interfaces

import (
	"fmt"
	"testing"
)

// 案例一：
type MyErr struct {
	Msg string
}

func GetErr() *MyErr {
	var myErr *MyErr
	return myErr

	// 上述代码 等同于: return nil
}

func (m *MyErr) Error() string {
	return ""
}

func TestGetErr(t *testing.T) {
	var e error

	e = GetErr()
	fmt.Println(e == nil) // false

	var e2 *MyErr

	e2 = GetErr()
	fmt.Println(e2 == nil) // true
}

// 案例二：

type Base interface {
	do()
}

type App struct {
}

func (a *App) do() {}

func GetApp() *App {
	return nil
}

func GetApp2() Base {
	/*
		app是一个指向nil的空指针，但是最后return app 会触发匿名变量 App = app值拷贝动作，
		所以最后GetApp2() 返回给上层的是一个 Base insterface{}类型，也就是一个iface struct{}类型。

		app为nil，只是 iface 中的data 为nil而已。 但是iface struct{}本身并不为nil.
		因为iface结构体中 tab  *itab 字段是有值的，具体包含了：
			1. _type 表示具体化的类型，有值，具体内容指向： type App struct{}
			2.fun 表示具体类型所实现的方法，有值，具体内容指向： func (a *App) do() {}

		非空接口定义：
		type iface struct {
			tab  *itab
			data unsafe.Pointer
		}

		itab结构体定义：
		type itab struct {
			inter  *interfacetype   // 接口自身的元信息
			_type  *_type           // 具体类型的元信息
			link   *itab
			bad    int32
			hash   int32            // _type里也有一个同样的hash，此处多放一个是为了方便运行接口断言
			fun    [1]uintptr       // 函数指针，指向具体类型所实现的方法
		}

		但是 data unsafe.Pointer 字段是nil: = *App = nil
	*/
	var app *App
	return app
}

func GetApp3() *App {
	return nil
}

//  具体解释：https://www.kancloud.cn/aceld/golang/1958316
func TestGetApp(t *testing.T) {
	var base Base
	base = GetApp()

	fmt.Println("GetApp()1", base)
	fmt.Println("GetApp()1", base == nil)

	var base2 Base
	base2 = GetApp2()

	fmt.Println("GetApp()2", base2)        // <nil>
	fmt.Println("GetApp()2", base2 == nil) // false

	var base3 *App
	base3 = GetApp3()

	fmt.Println("GetApp()3", base3)        // <nil>
	fmt.Println("GetApp()3", base3 == nil) // true 因为GetApp()3返回值只是一个普通的指针结构体，不是接口Base
}

// 案例三：空接口判断nil
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func TestFoo(t *testing.T) {
	var p *int = nil
	Foo(p) // non-empty interface
}

/*
	Foo()的形参x interface{}是一个空接口类型eface struct{}。

	空接口定义：

	type eface struct {
		_type *_type         // 类型信息
		data  unsafe.Pointer // 指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
	}

	在执行Foo(p)的时候，触发x interface{} = p语句，所以此时 x结构如下。

	1. _type *type  表示数据类型的描述，字段有值，指向*int
	2. data unsafe.Pointer  表示具体的数据类型或者说实现类，值为nil：= p = nil

	所以 x 结构体本身不为nil，而是data指针指向的p为nil。
*/
