package contexts

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

// TestContextUse context cancel 使用示例
func TestContextCancel(t *testing.T) {
	var wg sync.WaitGroup
	// 初识化一个空context
	parent := context.Background()
	// 创建一个可取消的context
	ctx, cancel := context.WithCancel(parent)

	// goroutine 执行数量
	runTimes := 0

	wg.Add(1)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine done")
				return
			default:
				fmt.Println("goroutine running times : ", runTimes)
				runTimes++
			}
			if runTimes > 3 {
				cancel()
				wg.Done()
			}
		}
	}(ctx)
	wg.Wait()
}
