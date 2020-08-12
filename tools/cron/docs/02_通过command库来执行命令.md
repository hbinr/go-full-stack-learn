# 通过 command 库执行命令案例

> 声明：Windows10 下使用 bash.exe 命令，需要先安装和开启 Linux Bash

[安装和开启 Linux Bash](../docs/03_Win10安装和开启Linux%20Bash命令.md)

开发还是会用 Linux 环境吧，Windows 平台下总会有诸多的不便，遇不到还好，一旦遇到了解决问题就会耗费很多时间和精力，但是主要的事情却还做。相反，Linux 平台下已经早做完主要事情了

## 案列 1：通过调用 `go env`命令，输出 go 的环境配置

```go
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
```

`exec.Command()` 的参数，主要是能执行的命令就能使用，脚本(如 python 脚本)也可以

如：

```go
exec.Command("python","/usr/python/xxx.py")
```

## 案列 2：主动杀死正在执行的命令

```sh
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
	resultChan = make(chan *result, 10)
	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func() {
		// 创建命令cmd
		cmd = exec.CommandContext(ctx, "C:\\Windows\\System32\\bash.exe", "-c", "sleep 3;echo hello")
		// 执行命令，捕获子进程的输出
		output, err = cmd.CombinedOutput()

		// 将结果写入通道中
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()

	// sleep 2  秒
	time.Sleep(1 * time.Second)

	// 取消上下文，杀死子进程
	cancelFunc()

	// 在协程main里等待获取子协程的退出，并答应任务执行成果
	res := <-resultChan
	fmt.Printf("命令执行异常：%v,-----,命令执行结果：%v", res.err, string(res.output))
}
```

**实现的逻辑：**

执行 1 个 cmd，让它在一个协程里去执行，先睡 3 秒，然后输出 hello -> `sleep 3;echo hello`

然后主进程睡 2 秒，通过调用 `cancelFunc()` 杀死 cmd 子进程，这样命令 `echo hello`就不会执行了。

**`cancelFunc()`杀死进程原因解析：**

- `context` 接口有一个`Done()`方法，其是一个通道 `chan`，context 内部方法 `propagateCancel` 是通过 `select {case < ctx.Done()}`，来监听上下文 `context` 是否被关闭，调用 `cancelFunc()`即可让 `context`关闭
- 而`cancelFunc` 底层是调用 `cancel(removeFromParent bool, err error)`来取消上下文的，其内部做了 `close(ctx.Done)`，这样就保证了通道的关闭，不被阻塞。

- 上下文关闭后，操作系统就会 kill `C:\\Windows\\System32\\bash.exe` 程序，即 `kill pid` 杀死子进程

**第一次运行，控制台输出：**

> 命令执行异常：<nil>,-----,命令执行结果：sleep: cannot read realtime clock: Invalid argument

该问题主要原因是：

Windows 的子系统 Ubuntu20.04 是有 Bug 的（Ubuntu 20.04 LTS for Windows 10 Released on Microsoft Store），

解决：

sleep 命令是一个在 `/bin/sleep` 下的执行文件，用于暂停一定时间，先把 sleep 备份：

```sh
sudo mv /bin/sleep /bin/sleep~
```

那就写一个好了：

```sh
vi /bin/sleep
```

输入以下脚本：

```python
#!/bin/python3
import sys
from time import sleep
v=sys.argv[1]
u=v[-1]
if u=='m':
  sleep(int(v[:-1])*60)
elif u=='h':
  sleep(int(v[:-1])*3600)
elif u=='d':
  sleep(int(v[:-1])*86400)
elif u=='s':
  sleep(int(v[:-1]))
else:
  sleep(int(v))
```

sleep 就可以用了，只是性能稍微差点，不过也不会有太大影响。

**然后再运行程序，输出结果：**

> 命令执行异常：exit status 1,-----,命令执行结果：

可以看到，cmd 子进程被杀死后，output 是没有执行结果的，只有 err 有结果
