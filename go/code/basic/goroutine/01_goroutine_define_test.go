package goroutine

import (
	"fmt"
	"testing"
)

func action() {
	fmt.Println("action ......")
}

// TestDefine goroutine 定义
func TestDefine(t *testing.T) {
	go action()
}
