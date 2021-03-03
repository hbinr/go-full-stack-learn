package err

import (
	"errors"
	"fmt"
	"testing"
)

//  ErrString 自定义error，使用 string 作为底层数据结构，这样有问题嘛？
type ErrString string

// Error 实现 error 接口
func (err ErrString) Error() string {
	return string(err)
}

func New(text string) error {
	return ErrString(text) // 返回值并不是指针类型
}

// 定义哨兵 error
var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

func TestErrPointStruct(t *testing.T) {

	// 底层使用 string 作为error数据结构，判断会为true，因为string进行相等判断时，是值比较
	// 如果error底层就是采用 string 来设计的话，那么不同的开发者在定义哨兵error(即类似ErrNamedType)的时候，正好定义完全相同了
	// 这样就导致逻辑混乱，我自定义的err竟然和别人相同了，在err比较时候会出现意外的bug，本来不希望相等，结果if 判断却返回了true
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}

	// 底层使用 struct 作为error数据结构
	// 如果error底层就是采用 Struct 来设计，并且返回 error 对象的是指针类型，那么不同的开发者在定义哨兵error(即类似ErrNamedType)的时候
	// 便不会有上述问题，因为底层指针指向了不同的内存地址，那么就是不同的对象，这样就不会导致逻辑混乱
	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}
}
