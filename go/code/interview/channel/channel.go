package main

import (
	"fmt"
)

func main() {
	tmpCh := make(chan int, 1)     // 只有1个缓冲区的管道
	fmt.Println("tmpCh:", <-tmpCh) // fatal error: all goroutines are asleep - deadlock!
	tmpCh <- 1                     // 即使没有这行代码都会死锁
}
