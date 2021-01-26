package channel

import (
	"fmt"
	"testing"
	"time"
)

var ch = make(chan int, 10)

// producer 生产者
func producer(idx int) {
	ch <- idx
}

// consumer 消费者
func consumer() {
	fmt.Println("consumer : ", <-ch)
}

// TestProducerConsumer 生产者消费者模型
func TestProducerConsumer(t *testing.T) {

	for i := 0; i < 10; i++ { // 如果遍历超过10次,那么就会阻塞,因为不断的生产,消费不过来.也就是说写channel时,如果缓冲区满了,则会阻塞
		go producer(i)
	}

	for i := 0; i < 100; i++ { // 如果遍历超过10次,那么就会阻塞,因为不断的消费,生产不过来.也就是说读channel时,如果缓冲区空了,则会阻塞
		go consumer()
	}

	time.Sleep(5 * time.Second)
}
