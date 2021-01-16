package structs

import (
	"fmt"
	"reflect"
	"testing"
)

// People 可进行简单比较的结构体
type People struct {
	Name string
	Age  int
}

// Student 不能进行简单比较的结构体，只能使用反射来进行深度比较
type Student struct {
	Action  func()
	Books   map[string]string
	Message chan string
}

// TestStructSimpleCompare 简单比较
func TestStructSimpleCompare(t *testing.T) {
	p1 := People{
		Name: "Tom",
		Age:  20,
	}

	p2 := People{
		Name: "Tom",
		Age:  20,
	}

	fmt.Println("p1 == p2 :", p1 == p2) // p1 == p2 : true
}

// TestStructDeepCompare 深度比较
func TestStructDeepCompare(t *testing.T) {
	s1 := Student{
		Action: func() {},
		Books: map[string]string{
			"Bob": "How to study well",
		},
		Message: make(chan string),
	}

	s2 := Student{
		Action: func() {},
		Books: map[string]string{
			"Bob": "How to study well",
		},
		Message: make(chan string),
	}
	// fmt.Println("s1 == s2 :", s1 == s2) // 报错： invalid operation: s1 == s2 (struct containing func() cannot be compared)

	fmt.Println("s1 == s2 :", reflect.DeepEqual(s1, s2)) // s1 == s2 : false 可以比较,
	// false 不相等的原因是 func() {}和 make(chan string) 都在分配了各自的内存地址
}
