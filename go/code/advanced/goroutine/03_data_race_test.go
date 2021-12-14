package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type A struct {
	a []int
	b int
}

// TestDataRace 并发问题; goroutine 使用不规范导致数据竞争
func TestDataRace(t *testing.T) {
	data := &A{
		a: make([]int, 0),
	}
	for i := 0; i < 200; i++ {
		go func(index int) {
			data.a = append(data.a, index)
			data.b++
		}(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(data.a, len(data.a))
	fmt.Println(data.b)
}

// TestDataRace 并发问题示例2
func TestDataRace2(t *testing.T) {
	data := &A{
		a: make([]int, 0),
	}
	// 使用了waitgroup, 可以保证所有的goroutine都执行完毕
	wg := sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func(index int) {
			data.a = append(data.a, index)
			data.b++
			wg.Done()
		}(i)
	}

	wg.Wait()
	time.Sleep(time.Second * 3)
	fmt.Println(data.a, len(data.a))
	fmt.Println(data.b)
}

// TestDataRace 修复并发问题
func TestDataRaceRepair(t *testing.T) {
	data := &A{
		a: make([]int, 0),
	}
	lock := sync.Mutex{}
	for i := 0; i < 200; i++ {
		go func(index int) {
			lock.Lock()
			data.a = append(data.a, index)
			data.b++
			lock.Unlock()
		}(i)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(data.a, len(data.a))
	fmt.Println(data.b)
}
