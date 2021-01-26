package channel

import (
	"fmt"
	"testing"
)

// TestChannelImplMutexLock 通过channel实现互斥锁
func TestChannelImplMutexLock(t *testing.T) {
	counter := 0
	tmpCh := make(chan int, 1) // 只有1个缓冲区的管道也是关键

	tmpCh <- 1 // 写操作类似于加锁
	counter++
	res := <-tmpCh // 读操作类似于释放锁

	fmt.Println("tmpCh:", res)
}

// TestReadChannelLock1 先直接读取有缓冲区的chan，会死锁，并不是加锁操作
func TestReadChannelLock1(t *testing.T) {

	tmpCh := make(chan int, 1)     // 只有1个缓冲区的管道
	fmt.Println("tmpCh:", <-tmpCh) // fatal error: all goroutines are asleep - deadlock!
	tmpCh <- 1                     // 即使没有这行代码都会死锁
}

// TestReadChannelLock2 先直接读取无缓冲区的chan，会死锁，并不是加锁操作
func TestReadChannelLock2(t *testing.T) {

	tmpCh := make(chan int)        // 无缓冲区的管道
	fmt.Println("tmpCh:", <-tmpCh) // fatal error: all goroutines are asleep - deadlock!
	tmpCh <- 1                     // 即使没有这行代码都会死锁
}
