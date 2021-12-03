package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroupFor(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestWaitGroupForUseDefer(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			// use defer
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
