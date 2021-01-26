package channel

import (
	"fmt"
	"testing"
)

// TestChanClose channel 关闭
func TestChanClose(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	fmt.Println("cap(ch): ", cap(ch)) // cap(ch):  5
	fmt.Println("len(ch): ", len(ch)) // len(ch):  3

	fmt.Println("可以从关闭channel中接收(读)数据，value:", <-ch) // 1 从队首开始读取数据
	ch <- 5                                          // 不能从关闭channel中发送(写)数据，会引发panic
}
