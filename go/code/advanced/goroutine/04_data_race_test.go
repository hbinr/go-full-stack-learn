package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var counter int

//全局变量
var ticket = 10 // 10张票

var wg sync.WaitGroup
var mx sync.Mutex // 创建锁头

func TestDataRaceCorrectUse(t *testing.T) {
	// 4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。
	wg.Add(4)
	go saleTickets("售票口1") // g1
	go saleTickets("售票口2") // g2
	go saleTickets("售票口3") // g3
	go saleTickets("售票口4") // g4
	wg.Wait()              // 要等待

	// time.Sleep(5 * time.Second) 有了wait 可以去掉 sleep
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()

	for {
		mx.Lock()
		if ticket > 0 {
			//睡眠
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			mx.Unlock() //解锁
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
		mx.Unlock() //解锁
	}
}

func saleTickets2(name string) {
	rand.Seed(time.Now().UnixNano())

	for {
		if ticket > 0 { // 有票的时候开抢
			// 随机睡眠时间, 模仿业务耗时
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", ticket)
			ticket--
		} else {
			fmt.Println(name, "售罄，没有票了...")
			break
		}
	}
}

// TestDataRaceRepairByWaitgroup 通过 waitgroup 修复并发问题
func TestDataRaceRepairByWaitgroup(t *testing.T) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch2 <- 200
	}()
	go func() {
		ch1 <- 100
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1中取数据...", num1)
	case num2, ok := <-ch2: // 第二写法，更推荐(安全)
		if ok {
			fmt.Println("ch2中取数据...", num2)
		} else {
			fmt.Println("ch2通道已经关闭...")
		}
	}
}

// TestDataRaceRepairByMutex 通过锁修复并发问题
func TestDataRaceRepairByMutex(t *testing.T) {

}
