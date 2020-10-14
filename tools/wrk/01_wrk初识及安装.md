## 什么是wrk?
wrk 是一个很简单的 http 性能测试工具. 也可以叫做 http benchmark 工具. 只有一个命令行, 就能做很多基本的 http 性能测试.，他和 apache benchmark（ab）同属于性能测试工具，但是比 ab 功能更加强大，并且可以支持lua脚本来创建复杂的测试场景

wrk 的开源的, 代码在 github 上：https://github.com/wg/wrk

首先要说的一点是: wrk 只能运行在 Unix 类的系统上. 比如 linux, mac, solaris 等. 也只能在这些系统上编译. 

wrk 的一个很好的特性就是能用很少的线程压出很大的并发量. 原因是它使用了一些操作系统特定的高性能 io 机制, 比如 select, epoll, kqueue 等. 其实它是复用了 redis 的 ae 异步事件驱动框架.  确切的说 ae 事件驱动框架并不是 redis 发明的, 它来至于 Tcl的解释器 jim, 这个小巧高效的框架, 因为被 redis 采用而更多的被大家所熟知.

要用 wrk, 首先要编译 wrk.

你的机器上需要已经安装了 git 和基本的c编译环境. wrk 本身是用 c 写的. 代码很少. 并且没有使用很多第三方库.  所以编译基本不会遇到什么问题.

## 安装
1. 需要安装一下编译工具，通过运行下面命令来安装工具：

```sh
# 安装 make 工具
sudo apt-get install make
 
# 安装 gcc编译环境
sudo apt-get install build-essential
```
2. 安装完成之后安装git：

2.1.使用git -version命令查看环境是否安装git
2.2.如果没安装就使用命令：`apt-get install git` (需要root权限就使用：`sudo apt-get install git` )

3. 安装完成之后使用 git 下载 wrk 的源码到本地:

```go
#下载命令
git clone https://github.com/wg/wrk.git 

#切换路径到wrk目录下
cd wrk  

#使用make命令编译环境
make  
```

就 ok了. 
make 成功以后在目录下有一个 wrk 文件. 就是它了. 你可以把这个文件复制到其他目录, 比如 bin 目录. 或者就这个目录下执行. 

```sh
# 移到 bin目录下，就能全局使用了
sudo cp wrk /bin
```

### 验证
在终端输入 `wrk`，有以下提示表示安装成功：
```sh
Usage: wrk <options> <url>                            
  Options:                                            
    -c, --connections <N>  Connections to keep open   
    -d, --duration    <T>  Duration of test           
    -t, --threads     <N>  Number of threads to use   
                                                      
    -s, --script      <S>  Load Lua script file       
    -H, --header      <H>  Add header to request      
        --latency          Print latency statistics   
        --timeout     <T>  Socket/request timeout     
    -v, --version          Print version details      
                                                      
  Numeric arguments may include a SI unit (1k, 1M, 1G)
  Time arguments may include a time unit (2s, 2m, 2h)
```
### 异常情况：
如果编译过程中出现:
```c 
src/wrk.h:11:25: fatal error: openssl/ssl.h: No such file or directory  
 #include <openssl/ssl.h>  
```

是因为系统中没有安装openssl的库. 

```sh
sudo apt-get install libssl-dev 

或
sudo yum install  openssl-devel 
```


参考：

https://blog.csdn.net/qq_41030861/article/details/90553510