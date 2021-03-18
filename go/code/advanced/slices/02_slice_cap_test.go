/*
 * @Author: duanhaobin
 * @Date: 2021-03-18 15:22:58
 * @LastEditTime: 2021-03-18 16:03:11
 * @FilePath: \go-full-stack-learn\go\code\advanced\slices\02_slice_cap_test.go
 */
package slices

import (
	"fmt"
	"testing"
)

func TestSliceCap(t *testing.T) {
	var arr [10]int
	var sli = arr[5:6]
	fmt.Println("len(slice)：", len(sli))
	fmt.Println("cap(slice)：", cap(sli))
}
func SliceRise(s []int) {
	s = append(s, 0)
	fmt.Println("s append 后：", s)
	fmt.Printf("s append地址为； %p\n", s) //  0xc000014360

	for i := range s {
		s[i]++
	}

}
func TestSliceCapAppend(t *testing.T) {
	s1 := []int{1, 2}
	fmt.Printf("1. s1地址为； %p\n", s1) //  0xc000014360

	s2 := s1
	s2 = append(s2, 3)               // 因为容量不够，所以扩容了，指向的新的array地址。变为[1,2,3]
	fmt.Printf("1. s2地址为； %p\n", s2) //  0xc000014360

	SliceRise(s1) // 值拷贝，所以即使调用改方法，s1的内容不变，因为底层指向array的地址未变
	SliceRise(s2) // 注意：s2 = append(s2,3)，扩容后为[1,2,3]，cap为4，s2指向了新的地址空间
	// 因此调用SliceRise中s = append(s, 0)便不会扩容，不会地址空间未更改，那么值就会变化
	fmt.Println(s1, s2)             // [1 2] [2 3 4]
	fmt.Printf("2. s1地址为；%p\n", s1) // 0xc000014360
	fmt.Println("len(s1)", len(s1)) // 2
	fmt.Println("cap(s1)", cap(s1)) // 2
}

func TestMakeSlice(t *testing.T) {
	var s1 []int
	fmt.Println("len(s1)：", len(s1))
	fmt.Println("cap(s1)：", cap(s1))

	s2 := make([]int, 1)
	fmt.Println("len(s2)：", len(s2))
	fmt.Println("cap(s2)：", cap(s2))
}
