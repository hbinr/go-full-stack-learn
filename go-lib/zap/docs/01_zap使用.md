# zap 日志库使用

## 介绍

在许多 Go 语言项目中，我们需要一个好的日志记录器能够提供下面这些功能：

- 能够将事件记录到文件中，而不是应用程序控制台。
- 日志切割-能够根据文件大小、时间或间隔等来切割日志文件。
- 支持不同的日志级别。例如 INFO，DEBUG，ERROR 等。
- 能够打印基本信息，如调用文件/函数名和行号，日志时间等。

## 默认的 Go Logger

在介绍 Uber-go 的 zap 包之前，让我们先看看 Go 语言提供的基本日志功能。Go 语言提供的默认日志包是https://golang.org/pkg/log/。

### 实现 Go Logger

实现一个 Go 语言中的日志记录器非常简单——创建一个新的日志文件，然后设置它为日志的输出位置。

#### 设置 Logger

我们可以像下面的代码一样设置日志记录器

```go
func SetupLogger() {
	logFileLocation, _ := os.OpenFile("/Users/q1mi/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}
```

#### 使用 Logger

让我们来写一些虚拟的代码来使用这个日志记录器。

在当前的示例中，我们将建立一个到 URL 的 HTTP 连接，并将状态代码/错误记录到日志文件中。

```go
func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}
}
```

#### Logger 的运行

现在让我们执行上面的代码并查看日志记录器的运行情况。

```go
func main() {
	SetupLogger()
	simpleHttpGet("www.google.com")
	simpleHttpGet("http://www.google.com")
}
```

当我们执行上面的代码，我们能看到一个 test.log 文件被创建，下面的内容会被添加到这个日志文件中。

> 2019/05/24 01:14:13 Error fetching url www.google.com : Get www.google.com: unsupported protocol scheme ""
> 2019/05/24 01:14:14 Status Code for http://www.google.com : 200 OK

### Go Logger 的优势和劣势

#### 优势

它最大的优点是使用非常简单。我们可以设置任何 `io.Writer` 作为日志记录输出并向其发送要写入的日志。

#### 劣势

- **仅限**基本的**日志级别**

  - 只有一个 `Print` 选项。不支持 `INFO/DEBUG` 等多个级别。

- 对于错误日志，它有 `Fatal` 和 `Panic`

  - `Fatal` 日志通过调用 `os.Exit(1)`来结束程序
  - `Panic` 日志在写入日志消息之后抛出一个 panic
  - 但是它缺少一个 `ERROR` 日志级别，这个级别可以在**不抛出 panic **或**退出程序的情况**下记录错误

- 缺乏**日志格式化**的能力——例如记录调用者的函数名和行号，格式化日期和时间格式。等等。
- 不提供**日志切割**的能力。

## Uber-go Zap

[Zap](https://github.com/uber-go/zap) 是非常快的、结构化的，分日志级别的 Go 日志库。

### 特点

- 它同时提供了结构化日志记录和 printf 风格的日志记录
- 它非常的快

### 安装

运行下面的命令安装 zap

```go
go get -u go.uber.org/zap
```

### 配置 Zap Logger

Zap 提供了两种类型的日志记录器——Sugared Logger 和 Logger。

在**性能很好但不是很关键的上下文**中，使用 `SugaredLogger` 。它比其他结构化日志记录包快 4-10 倍，并且支持结构化和 printf 风格的日志记录。

在**每一微秒和每一次内存分配都很重要**的上下文中，使用 `Logger` 。它甚至比 `SugaredLogger` 更快，内存分配次数也更少，但它**只支持强类型**的**结构化**日志记录。

#### Logger

- 通过调用 `zap.NewProduction()`/`zap.NewDevelopment`()或者 `zap.Example()`创建一个 Logger。
- 上面的每一个函数都将创建一个 logger。唯一的区别在于它将记录的信息不同。例如 production logger 默认记录调用函数信息、日期和时间等。
- 通过 Logger 调用 Info/Error 等。
- 默认情况下日志都会打印到应用程序的 console 界面。

```go
package main

import (
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

// InitLogger 初始化 zap.Looer
func InitLogger() {
	logger, _ = zap.NewProduction()
}
func main() {
	InitLogger()
	// Sync调用基础Core的Sync方法，刷新所有缓冲的日志词条
	// 应用程序应注意退出前调用Sync
	defer logger.Sync()
	testHTTPGet("www.baidu.com")        // error
	testHTTPGet("http://www.baidu.com") // success
}

// testHttpGet 测试日志
func testHTTPGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
```

在上面的代码中，我们首先创建了一个 Logger，然后使用 Info/ Error 等 Logger 方法记录消息。

日志记录器方法的语法是这样的：

```go
func (log *Logger) MethodXXX(msg string, fields ...Field)
```

其中 `MethodXXX` 是一个可变参数函数，可以是 Info / Error/ Debug / Panic 等。每个方法都接受一个消息字符串和任意数量的 `zapcore.Field` 场参数。

每个 zapcore.Field 其实就是一组键值对参数。

我们执行上面的代码会得到如下输出结果：

> {"level":"error","ts":1601641640.1690316,"caller":"logger/main.go:28","msg":"Error fetching url..","url":"www.baidu.com","error":"Get \"www.baidu.com\": unsupported protocol scheme \"\"","stacktrace":"main.testHTTPGet\n\t/home/hblock/Develop/WorkSpace/go/src/hb.study/go_web/zap/code/logger/main.go:28\nmain.main\n\t/home/hblock/Develop/WorkSpace/go/src/hb.study/go_web/zap/code/logger/main.go:20\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:204"}

> {"level":"info","ts":1601641640.1990101,"caller":"logger/main.go:33","msg":"Success..","statusCode":"200 OK","url":"http://www.baidu.com"}

#### Sugared Logger

现在让我们使用 Sugared Logger 来实现相同的功能。

- 大部分的实现基本都相同。
- 惟一的区别是，我们通过调用主 logger 的`.Sugar()`方法来获取一个 `SugaredLogger`。
- 然后使用 `SugaredLogger` 以 `printf` 格式记录语句

下面是修改过后使用 `SugaredLogger` 代替 `Logger` 的代码：

```go
var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	// Sync调用基础Core的Sync方法，刷新所有缓冲的日志词条
	// 应用程序应注意退出前调用Sync
	defer logger.Sync()

	testHTTPGet2("www.baidu.com")        // error
	testHTTPGet2("http://www.baidu.com") // success
}

// InitLogger 初始化 zap.Looer和sugarLogger
func InitLogger() {
	logger, _ = zap.NewProduction()
	sugarLogger = logger.Sugar()
}


// testHttpGet 测试sugarLogger日志
func testHTTPGet2(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
```

当你执行上面的代码会得到如下输出：

> {"level":"error","ts":1601642124.335311,"caller":"logger/main.go:51","msg":"Error fetching URL www.baidu.com : Error = Get \"www.baidu.com\": unsupported protocol scheme \"\"","stacktrace":"main.testHTTPGet2\n\t/home/hblock/Develop/WorkSpace/go/src/hb.study/go_web/zap/code/logger/main.go:51\nmain.main\n\t/home/hblock/Develop/WorkSpace/go/src/hb.study/go_web/zap/code/logger/main.go:20\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:204"}

> {"level":"info","ts":1601642124.367956,"caller":"logger/main.go:53","msg":"Success! statusCode = 200 OK for URL http://www.baidu.com"}

你应该注意到的了，到目前为止这两个 logger 都打印输出 JSON 结构格式。

参考：

https://www.liwenzhou.com/posts/Go/zap/

```

```
