package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)
	// 生成 command，输出 go 的环境配置
	cmd = exec.Command("go", "env")

	// 执行命令，捕获子进程的输出
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println("执行命令异常，err:", err)
		return
	}

	// 打印输出
	fmt.Println(string(output))
}
