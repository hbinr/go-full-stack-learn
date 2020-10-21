package datatype

import (
	"fmt"
	"testing"
)

// TestConvertType  数据类型转换
func TestConvertType(t *testing.T) {
	
	price := 1.55
	weight := 5
	sumPrice := price * float64(weight) // 变量类型转换
	fmt.Println("整数转float 不会丢失精度:", sumPrice)

	sumPrice2 := int(price) * weight
	fmt.Println("float转整数，直接去掉整数部分，不是四舍五入:", sumPrice2)

	sumPrice3 := int(price * float64(weight)) // 表达式类型转换
	fmt.Println("float转整数，直接去掉整数部分，不是四舍五入:", sumPrice3)
}

type int1 int
type int2 int

type int3 = int
type int4 = int

// TestAlias  起别名的变量参与计算
func TestAlias(t *testing.T) {
	//var a int1 = 1
	//var b int2 = 2
	//fmt.Println("别名没有等号，不能进行计算，因为是两个不同的类型", a+b)

	var c int3 = 1
	var d int4 = 2
	fmt.Println("别名有等号，能进行计算，因为是两个相同同的类型", c+d)
}
