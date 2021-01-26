package channel

import (
	"fmt"
	"sync"
	"testing"
)

// TestChannelDataRace channel 资源争夺案例 100个人抢10张票
func TestChannelDataRace(t *testing.T) {
	tickets := make(chan int, 10)

	for i := 0; i < 10; i++ {
		tickets <- i // 10张票
	}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(num int) {
			select {
			case ticket := <-tickets:
				fmt.Printf("People %d get ticket %d\n", num, ticket)
			default:
			}
			wg.Done()
		}(i) // 细节：实参为i
	}
	wg.Wait()
}
