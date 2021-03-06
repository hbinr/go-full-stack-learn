package goroutine

import (
	"fmt"
	"testing"
)

//  TestWildGoroutine 错误野生goroutine示例
func TestWildGoroutineIncorrect(t *testing.T) {
	go func() {
		panic("wild goroutine happened") // 如果程序panic，那么整个进程都会退出
	}()
}

//  TestWildGoroutineCorrect 正确goroutine示例
func TestWildGoroutineCorrect(t *testing.T) {
	Go(func() {
		fmt.Println("Hello world")
	})

}

func Go(x func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				panic("wild goroutine happened") // 如果程序panic，会被recover捕获，进程不会直接退出
			}
		}()
		x()
	}()
}
