package slice

import (
	"testing"
)

// TestAppend 切片追加问题，不要忽略默认值
func TestAppend(t *testing.T) {
	tmp := make([]int, 2) // 长度为2的切片
	tmp = append(tmp, 1, 2, 3)
	t.Log(tmp) // [0 0 1 2 3]
}

