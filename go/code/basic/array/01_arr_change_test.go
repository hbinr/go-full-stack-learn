package array

import (
	"fmt"
	"testing"
)

func TestArrChange(t *testing.T) {
	arr := [3]int{10, 20, 30}
	fmt.Println("arr[1]: 1", arr[1]) // arr[1]:  20

	arr = [3]int{10, 21, 30} // 重新赋值

	fmt.Println("arr[1]:  2", arr[1]) // arr[1]:  21
}
