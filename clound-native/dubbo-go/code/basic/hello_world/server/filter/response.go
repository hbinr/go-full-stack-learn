package filter

import (
	"context"
	"errors"
	"fmt"

	"github.com/apache/dubbo-go/filter"
	"github.com/apache/dubbo-go/protocol"
)

// func init() {
// 	extension.SetFilter("ErrResponseFilter", GetErrResponseFilter)
// }

type errResponseFilter struct {
	Code int
	Data interface{}
}

func (mf errResponseFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	// the logic put here...
	// you can get many params in url. And the invocation provides more information about
	fmt.Println("Invoke=============================")
	fmt.Printf("invoker:%v\n", invoker)
	fmt.Println("---------")
	fmt.Printf("invocation:%v\n", invocation)
	return invoker.Invoke(ctx, invocation)
}

func (mf errResponseFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	fmt.Println("OnResponse=============================")
	fmt.Printf("result:%v\n", result)
	fmt.Printf("invoker.GetUrl().Service():%v\n", invoker.GetUrl().Service()) // com.hbstudy.user.UserProvider
	// invalidParm := errors.New("invalid param")
	// if errors.Is(result.Error(), invalidParm) {
	// }\
	result.SetError(errors.New("ssssssssssssssssssssssssssss")) // 修改响应错误
	result.SetResult("sssssss-------wdddddddddddd")
	fmt.Printf("invoker:%v\n", invoker)
	fmt.Printf("invocation:%v\n", invocation)
	return result
}
func GetErrResponseFilter() filter.Filter {
	return &errResponseFilter{}
}
