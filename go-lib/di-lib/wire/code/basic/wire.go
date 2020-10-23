// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

// InitializeEvent 声明injector的函数签名，这个需要自定义
func InitializeEvent(msg string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil //返回值没有实际意义，只需符合函数签名即可
}
