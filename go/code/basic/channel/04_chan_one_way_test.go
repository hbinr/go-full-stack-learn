package channel

import "testing"

// TestChanOneWay 单向通道
func TestChanOneWay(t *testing.T) {
	out := make(chan<- int)
	out <- 1
	out <- 1
	onlySend(out)
	// fmt.Printf("res :%v", <- onlySend(out))
	// 报错: invalid operation: cannot receive from send-only channel onlySend(out) (variable of type chan<- int)compiler

}

// onlySend 只能发送(只写)channel
func onlySend(out chan<- int) chan<- int {

	return out
 
}
