package functionalprogram

import (
	"fmt"
	"testing"
)

//  TestMapForEach 函数式编程简单demo
func TestMapForEach(t *testing.T) {
	var list = []string{"Orange", "Apple", "Banana", "Grape"}
	var out = mapForEach(list, func(it string) int {
		return len(it)
	})
	fmt.Println("mapForEach:", out)
}

func mapForEach(arr []string, getStrLength func(str string) int) []int {
	var newArr = []int{}

	for _, v := range arr {
		newArr = append(newArr, getStrLength(v))
	}

	return newArr
}
