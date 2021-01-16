package slices

import (
	"bytes"
	"fmt"
	"testing"
)

/*
	参考：GO编程模式系列 https://coolshell.cn/articles/21128.html

	核心在slice的底层结构+扩容细节：
	type slice struct {
		array unsafe.Pointer //指向存放数据的数组指针
		len   int            //长度有多大
		cap   int            //容量有多大
	}

	知识点：
	1.切片分割会共享内存
	2.append()这个函数在 cap 不够用的时候就会重新分配内存以扩大容量，而如果够用的时候不不会重新分享内存！
*/

// TestSliceSplit 切片分割之坑
func TestSliceSplit(t *testing.T) {
	foo := make([]int, 5) // len:5,cap:5
	foo[3] = 42
	foo[4] = 100
	fmt.Println("foo is ", foo)         // [0 0 0 42 100]
	fmt.Println("len(foo): ", len(foo)) // 5
	fmt.Println("cap(foo): ", cap(foo)) // 5

	bar := foo[1:4]                     // 由foo分割而来，底层array指向同一个数组地址，foo 和 bar 的内存空间是共享的
	fmt.Println("bar is ", bar)         // [0 0 42]
	fmt.Println("len(bar): ", len(bar)) // 3
	fmt.Println("cap(bar): ", cap(bar)) // 4 因为bar是截取foo的部分得到的，从foo下标1开始截取直到末尾，因此bar容量便少了foo下标为 0 的部分

	bar[1] = 99                 // bar修改，foo也会修改
	fmt.Println("bar is ", bar) // [0 99 42]
	fmt.Println("foo is ", foo) // [0 0 99 42 100]

}

// TestSliceAppend 切片扩容之坑
func TestSliceAppend(t *testing.T) {
	a := make([]int, 32)
	b := a[1:16]                    // a 和 b 的内存空间是共享的
	fmt.Println("len(b): ", len(b)) // 15
	fmt.Println("cap(b): ", cap(b)) // 31

	// 扩容坑1
	a = append(a, 1) // 原容量为32，再新增元素会发生扩容，底层array指向了新的数组地址，a 和 b 的内存空间不再共享
	a[2] = 42

	fmt.Println("a is ", a)         // [0 0 42 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
	fmt.Println("len(a): ", len(a)) // 33
	fmt.Println("cap(a): ", cap(a)) // 64

	fmt.Println("b is ", b) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]  b中下标为1的地方并没有变为42，还是0，是因为a和b底层array指向的数组地址不同了
}

// TestSliceAppend 切片扩容之坑2
func TestSliceAppend2(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	fmt.Println("len(path): ", len(path)) // 14
	fmt.Println("cap(path): ", cap(path)) // 14

	sepIndex := bytes.IndexByte(path, '/') // sepIndex=4

	dir1 := path[:sepIndex]
	fmt.Println("dir1 is ", string(dir1)) // AAAA
	fmt.Println("len(dir1): ", len(dir1)) // 4
	fmt.Println("cap(dir1): ", cap(dir1)) // 14

	dir2 := path[sepIndex+1:]
	fmt.Println("dir2 is ", string(dir2)) // BBBBBBBBB
	fmt.Println("len(dir2): ", len(dir2)) // 9
	fmt.Println("cap(dir2): ", cap(dir2)) // 9
	// 扩容坑2，
	dir1 = append(dir1, "suffix"...) // suffix的长度为6， 因为 len(dir1)+6 = 10 ，小于cap(dir1) 14，
	// 因此，append操作并未发生扩容，dir1和dir2还共享内存

	fmt.Println("dir1 is ", string(dir1)) // AAAAsuffix
	fmt.Println("dir2 is ", string(dir2)) // uffixBBBB

}

// TestSliceAppend2Fix 解决 切片扩容之坑2 中的问题
func TestSliceAppend2Fix(t *testing.T) {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/') // sepIndex=4

	// 如果要解决这个问题，我们只需要修改一行代码
	dir3 := path[:sepIndex:sepIndex]      // dir1 := path[:sepIndex] 改为左侧的形式
	fmt.Println("dir3 is ", string(dir3)) // AAAA
	// 新的代码使用了 Full Slice Expression，其最后一个参数叫“Limited Capacity”，于是，后续的 append() 操作将会导致重新分配内存。

	dir4 := path[sepIndex+1:]
	fmt.Println("dir4 is ", string(dir4)) // BBBBBBBBB
	fmt.Println("len(dir4): ", len(dir4)) // 9
	fmt.Println("cap(dir4): ", cap(dir4)) // 9

	dir3 = append(dir3, "suffix"...)
	fmt.Println("dir3 is ", string(dir3)) // AAAAsuffix
	fmt.Println("dir4 is ", string(dir4)) // BBBBBBBBB  数据未变

}
