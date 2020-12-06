package sync

import (
	"fmt"
	"sync"
	"testing"
)

func out(in func()) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Println("recover。。。，err", e)
				wg.Done()
			}
		}()
		in()
		wg.Done()
	}()
	wg.Wait()
}

func out2(in func()) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if e := recover(); e != nil {
				fmt.Println("recover。。。，err", e)
				//wg.Done() // 此处Done()会报错，如果panic发送了，waitGroup计数为负数了
			}
		}()
		in()
		//wg.Done() // 此处Done()会报错，waitGroup计数为负数了
	}()
	wg.Wait()
}
func out3(in func()) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			if e := recover(); e != nil {
				fmt.Println("recover。。。，err", e)
			}
		}()
		in()
	}()
	wg.Wait()
}

func noPainc() {
	fmt.Println("no panic")
}
func triggerPainc() {
	panic("triggerPainc")
}

// TestWaitGroup .
func TestWaitGroup(t *testing.T) {
	out(noPainc)
	out(triggerPainc)

	out2(noPainc)
	out2(triggerPainc)

	out3(noPainc)
	out3(triggerPainc)
}
