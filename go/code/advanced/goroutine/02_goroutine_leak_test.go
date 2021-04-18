package goroutine

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
	https://segmentfault.com/a/1190000019644257
*/
// goroutine 泄漏
func TestGoroutineLeak(t *testing.T) {
	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			requestWork(context.Background(), "any")
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
}

//
func hardWork(job interface{}) error {
	time.Sleep(time.Minute)
	return nil
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	done := make(chan error)
	go func() {
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
