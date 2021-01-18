package goroutine

import (
	"fmt"
	"testing"
	"time"
)

// TestDeadLock 死锁
func TestDeadLock(t *testing.T) {
	ch1 := make(chan int)
	go fmt.Println(<-ch1)
	// time.Sleep(1 * time.Second) // 休眠即使位于此处也会死锁
	ch1 <- 5
	time.Sleep(1 * time.Second)

}
