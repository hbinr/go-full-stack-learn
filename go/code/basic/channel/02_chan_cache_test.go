package channel

import (
	"fmt"
	"testing"
)

// TestChanCache 有缓冲chan声明
func TestChanCache(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("cap(ch): ", cap(ch)) // cap(ch):  5
	fmt.Println("len(ch): ", len(ch)) // len(ch):  3
}
