# wrk基本使用

终端运行 `wrk`命令，会有以下提示：
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
## 参数解析
- -c 需要模拟的连接数
- -d 测试的持续时间，时间越长样本越准确. 如果想测试系统的持续抗压能力
- -t 需要模拟的线程数

- -s 加载的Lua脚本文件
- -H, --header        增加请求头
       --latency      打印延迟统计   
       --timeout      超时的时间     
- -v 打印版本详细信息


## 实战
测试 www.baidu.com，12个线程，100个连接，压测30s。

```sh
hblock@hblock:~$ wrk -t12 -c100 -d30s http://www.baidu.com  

Running 30s test @ http://www.baidu.com
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.50s   434.66ms   2.00s    81.99%
    Req/Sec     7.36      8.09    70.00     89.34%
  1382 requests in 30.07s, 20.64MB read
  Socket errors: connect 0, read 0, write 0, timeout 249
Requests/sec:     45.97
Transfer/sec:    703.13KB
```

**结果参数解析：**

- Latency：可以理解为响应时间, 有平均值, 标准偏差, 最大值, 正负一个标准差占比
- Req/Sec：每个线程每秒钟的完成的请求数, 同样有平均值, 标准偏差, 最大值, 正负一个标准差占比
- 1382 requests in 30.07s, 20.64MB read：30.07s内共处理1382个请求，具体数字大小看个人电脑配置；读取数据量为20.64MB
- Socket errors: connect 0, read 0, write 0, timeout 249：错误统计, 0个连接错误， 0个读/写错误, 249个超时。wrk 默认超时时间是1秒，所以常会有`Socket errors`， 可以通过加参数设置，如`-T10s` 
- Requests/sec：一般称之为QPS（每秒请求数），这是一项压力测试的性能指标，通过这个参数我们可以看出应用程序的吞吐量
- Transfer/sec

- Avg：平均
- Max：最大
- Stdev：标准差
- +/- Stdev： 正负一个标准差占比


标准差`Stdev`如果太大说明样本本身离散程度比较高. 有可能系统性能波动很大.

### 加 `--latency`，看响应时间的分布情况

命令：
```sh
wrk -t12 -c100 -d30s -T10s --latency http://www.baidu.com  
```
结果：
```sh
Running 30s test @ http://www.baidu.com
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   659.45ms  319.44ms   5.06s    90.76%
    Req/Sec    13.82      7.74    70.00     77.74%
  Latency Distribution
     50%  638.71ms
     75%  778.38ms
     90%  881.88ms
     99%    2.09s 
  4454 requests in 30.06s, 66.15MB read
Requests/sec:    148.16
Transfer/sec:      2.20MB
```
可以看到，多了以下内容：
```sh
Latency Distribution
     50%  638.71ms
     75%  778.38ms
     90%  881.88ms
     99%    2.09s 
```

50%的请求在638.71ms以内响应, 75%则在778.38ms 以内， 90%则在881.88ms 以内， 99%则在2.09s 以内 

参考：

https://blog.csdn.net/qq_41030861/article/details/90553510