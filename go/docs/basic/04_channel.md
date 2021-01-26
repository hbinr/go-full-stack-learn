# channel基本知识

管道(channel)是Go在语言层面提供的协程间的通信方式，比UNIX的管道更易用也更轻便。

## 声明一个 channel
在 Go 语言中，声明一个 channel 非常简单，使用内置的 make 函数即可，如下所示：

复制代码
```go
ch:=make(chan string)
```
其中 chan 是一个关键字，表示是 channel 类型。后面的 string 表示 channel 里的数据是 string 类型。通过 channel 的声明也可以看到，chan 是一个集合类型。

定义好 chan 后就可以使用了，一个 chan 的操作只有两种：发送和接收。

- 接收：获取 chan 中的值，操作符为 <- chan。
- 发送：向 chan 发送值，把值放在 chan 中，操作符为 chan <-。

> 小技巧：这里注意发送和接收的操作符，都是 <- ，只不过位置不同。接收的 <- 操作符在 chan 的左侧，发送的 <- 操作符在 chan 的右侧。

### 实战


```go
func TestChanDefine(t *testing.T) {
	strCh := make(chan string)

	go func() {
		fmt.Println("sub goroutine exec")
		strCh <- "I'm sub goroutine"
	}()

    fmt.Println("main goroutine......")
    
	fmt.Println("value from sub goroutine:", <-strCh)
}
```
在上面的示例中，我们在新启动的 `goroutine` 中向 `chan` 类型的变量 `strCh` 发送值；在 `main goroutine` 中，从变量 `strCh` 接收值；如果 `strCh` 中没有值，则阻塞等待到 `strCh` 中有值可以接收为止。

`channel` 有点像在两个 `goroutine` 之间架设的管道，一个 `goroutine` 可以往这个管道里发送数据，另外一个可以从这个管道里取数据，有点类似于我们说的队列。

## 无缓冲 channel
上面的示例中，使用 `make` 创建的 `chan` 就是一个无缓冲 `channel` ，它的容量是 0，不能存储任何数据。所以无缓冲 `channel` 只起到传输数据的作用，数据并不会在 `channel` 中做任何停留。这也意味着，无缓冲 `channel` 的发送和接收操作是同时进行的，它也可以称为同步 `channel` 。

## 有缓冲 channel
有缓冲 `channel` 类似一个可阻塞的队列，内部的元素先进先出。通过 `make` 函数的第二个参数可以指定 `channel` 容量的大小，进而创建一个有缓冲 `channel` ，如下面的代码所示：
```go
cacheCh:=make(chan int,5)
```

我创建了一个容量为 5 的 channel，内部的元素类型是 int，也就是说这个 channel 内部最多可以存放 5 个类型为 int 的元素，如下图所示：
![](https://s0.lgstatic.com/i/image/M00/70/AE/CgqCHl-7fzmAVLu0AACSjW-neAE188.png)

一个有缓冲 channel 具备以下特点：

1. 有缓冲 `channel` 的内部有一个缓冲队列；
2. **发送**操作是向队列的**尾部插入元素**，如果队列已**满**，则**阻塞等待**，直到另一个 `goroutine` 执行，接收操作释放队列的空间；
3. **接收**操作是从**队列的头部获取元素**并把它**从队列中删除**，如果队列为**空**，则**阻塞等待**，直到另一个 `goroutine` 执行，发送操作插入新的元素。

因为有缓冲 `channel` 类似一个队列，可以获取它的**容量**和**里面元素的个数**。如下面的代码所示：
```go
// TestChanCache 有缓冲chan声明
func TestChanCache(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("cap(ch): ", cap(ch)) // cap(ch):  5
	fmt.Println("len(ch): ", len(ch)) // len(ch):  3
}
```

其中，通过内置函数 `cap` 可以获取 `channel` 的容量，也就是最大能存放多少个元素，通过内置函数 len 可以获取 `channel` 中元素的个数。

> 小提示：无缓冲 `channel` 其实就是一个容量大小为 0 的 `channel`。比如 `make(chan int,0)`。

## 关闭 channel
`channel` 还可以使用内置函数 `close()` 关闭，如下面的代码所示：
```go
close(cacheCh)
```
如果一个 `channel` 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 `painc` 异常。但是还可以接收 `channel` 里的数据，**如果 `channel` 里没有数据的话，接收的数据是元素类型的零值。**


## 单向 channel
有时候，我们有一些特殊的业务需求，比如限制一个 `channel` 只可以接收但是不能发送，或者限制一个 `channel` 只能发送但不能接收，这种 `channel` 称为单向 `channel`。

单向 `channel` 的声明也很简单，只需要在声明的时候带上 `<-` 操作符即可，如下面的代码所示：
```go
onlySend := make(chan<- int)
onlyReceive:=make(<-chan int)
```
> **注意:** 声明单向 `channel` `<-` 操作符的位置和上面讲到的发送和接收操作是一样的。

在**函数或者方法的参数**中，使用单向 `channel` 的较多，这样可以防止一些操作影响了 `channel` 。

下面示例中的 `counter` 函数，它的参数 `out` 是一个只能发送的 `channel` ，所以在 `counter` 函数体内使用参数 `out` 时，只能对其进行发送操作，如果执行接收操作，则程序不能编译通过。
```go
func counter(out chan<- int) {
  //函数内容使用变量out，只能进行发送操作
}
```

## 总结
如何通过 `channel` 实现 `goroutine` 间的数据传递，这些都是 `Go` 语言并发的基础，理解它们可以更好地掌握并发。

在 Go 语言中，提倡通过通信来共享内存，而不是通过共享内存来通信，其实就是提倡通过 `channel` 发送接收消息的方式进行数据传递，而不是通过修改同一个变量。所以在数据流动、传递的场景中要优先使用 `channel` ，它是并发安全的，性能也不错。

