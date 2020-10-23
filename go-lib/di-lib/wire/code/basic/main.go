package main

import (
	"errors"
	"fmt"
)

type Message struct {
	msg string
}
type Greeter struct {
	Message Message
}
type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent Event构造函数
// func NewEvent(g Greeter) Event {
// 	return Event{Greeter: g}
// }

// NewEvent Event构造函数 新增了错误返回
func NewEvent(g Greeter) (Event, error) {
	return Event{Greeter: g}, errors.New("error test")
}

// Greet Greeter结构体的方法
func (g Greeter) Greet() Message {
	return g.Message
}

// Start Event结构体的方法
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

/*
// 使用wire前
func main() {
	message := NewMessage("hello world")
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}
*/

// 使用wire后
func main() {
	event, _ := InitializeEvent("hello_world")
	event.Start()
}
