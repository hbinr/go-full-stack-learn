package waitgroup

import (
	"fmt"
	"sync"
	"testing"
)

// TestWaitGroupUse WaitGroup 使用示例
func TestWaitGroupUse(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1) // 增加1个goroutine 等待数
		go func(num int) {
			fmt.Println("goroutine test ", num)
			wg.Done() // 每完成1个goroutine便减1，直到0时，才释放阻塞时等待的goroutine；为负数会panic
		}(i)
	}
	wg.Wait() // 阻塞等待，直到WaitGroup计数器为零
}
