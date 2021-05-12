package structs

import (
	"fmt"
	"testing"
)

type User struct {
	name string
	age  int
}

func TestStructChange(t *testing.T) {
	user := User{
		name: "tom",
		age:  18,
	}
	fmt.Println("user.name: 1", user.name)
	fmt.Printf("user ptr old :%p\n", &user) // user ptr old :0xc000004090
	user.name = "bob"
	fmt.Println("user.name: 2", user.name)

	user = User{ // 重新赋值
		name: "bob",
		age:  18,
	}

	fmt.Println("user.name: 3", user.name)
	fmt.Printf("user ptr new :%p", &user) // user ptr new :0xc000004090 地址未变，还是指向原来的地址
}
