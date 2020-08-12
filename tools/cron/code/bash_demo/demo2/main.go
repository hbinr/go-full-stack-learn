package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	var (
		cmd        *exec.Cmd
		output     []byte
		err        error
		ctx        context.Context
		cancelFunc context.CancelFunc
		resultChan chan *result
	)
	resultChan = make(chan *result, 1000)
	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func() {
		// 创建命令cmd
		cmd = exec.CommandContext(ctx, "C:\\Windows\\System32\\bash.exe", "-c", "sleep 3;echo hello;")
		// 执行命令，捕获子进程的输出
		output, err = cmd.CombinedOutput()

		// 将结果写入通道中
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	// sleep 2  秒
	time.Sleep(2 * time.Second)

	// 睡2s后立即取消上下文，杀死子进程
	cancelFunc()

	// 在协程main里等待获取子协程的退出，并答应任务执行成果
	res := <-resultChan
	fmt.Printf("命令执行异常：%v,-----,命令执行结果：%v", res.err, string(res.output))
}
