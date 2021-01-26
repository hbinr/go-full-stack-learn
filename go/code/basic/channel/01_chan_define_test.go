package channel

import (
	"fmt"
	"testing"
)

// TestChanDefine channel 声明及简单使用
func TestChanDefine(t *testing.T) {
	strCh := make(chan string)

	go func() {
		fmt.Println("sub goroutine exec")
		strCh <- "I'm sub goroutine"
	}()

	fmt.Println("main goroutine......")
	fmt.Println("value from sub goroutine:", <-strCh)
}
