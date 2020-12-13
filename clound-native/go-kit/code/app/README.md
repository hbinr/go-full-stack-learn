httprouter、chi、gorilla/mux三个路由性能测试：


# 测试命令1：
```sh
 wrk -t12 -c100 -d30s http://localhost:8080/get/1
```

## httprouter

```sh
Running 30s test @ http://localhost:8080/get/1
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.12ms    1.74ms  34.84ms   88.87%
    Req/Sec    13.75k     2.00k   22.09k    68.92%
  4928647 requests in 30.03s, 606.34MB read
Requests/sec: 164115.67
Transfer/sec:     20.19MB

```

## gorilla/mux

```sh
Running 30s test @ http://localhost:8080/get/1
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.23ms    1.92ms  35.28ms   89.08%
    Req/Sec    12.74k     2.22k   21.02k    69.53%
  4568829 requests in 30.05s, 766.86MB read
  Non-2xx or 3xx responses: 4568829
Requests/sec: 152056.76
Transfer/sec:     25.52MB
```


## chi

```sh
Running 30s test @ http://localhost:8080/get/1
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.17ms    1.70ms  28.51ms   87.98%
    Req/Sec    12.60k     1.93k   20.72k    68.56%
  4516049 requests in 30.03s, 758.00MB read
  Non-2xx or 3xx responses: 4516049
Requests/sec: 150360.08
Transfer/sec:     25.24MB
```

# 测试命令2：

```sh
wrk -t30 -c500 -d30s http://localhost:8080/get/1
```


## httprouter

```sh
Running 30s test @ http://localhost:8080/get/1
  30 threads and 500 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.87ms    9.73ms 133.22ms   86.66%
    Req/Sec     4.62k     1.64k   63.12k    81.31%
  4145617 requests in 30.10s, 510.01MB read
Requests/sec: 137744.74
Transfer/sec:     16.95MB
```

## gorilla/mux

```sh
Running 30s test @ http://localhost:8080/get/1
  30 threads and 500 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.47ms    9.33ms 137.87ms   87.02%
    Req/Sec     4.91k     1.57k   24.58k    78.07%
  4405965 requests in 30.07s, 739.53MB read
  Non-2xx or 3xx responses: 4405965
Requests/sec: 146518.76
Transfer/sec:     24.59MB
```


## chi

```sh
Running 30s test @ http://localhost:8080/get/1
  30 threads and 500 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.77ms   10.08ms 230.53ms   87.43%
    Req/Sec     4.95k     1.64k   18.40k    76.87%
  4439506 requests in 30.06s, 745.16MB read
  Non-2xx or 3xx responses: 4439506
Requests/sec: 147699.61
Transfer/sec:     24.79MB
```

# 测试命令3：

```sh
wrk -t100 -c1000 -d30s http://localhost:8080/get/1
```


## httprouter

```sh
Running 30s test @ http://localhost:8080/get/1
  100 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.74ms   12.12ms 174.20ms   88.28%
    Req/Sec     1.44k   654.71    13.56k    89.02%
  4005587 requests in 30.10s, 492.78MB read
  Socket errors: connect 82, read 0, write 0, timeout 0
Requests/sec: 133080.41
Transfer/sec:     16.37MB
```

## gorilla/mux

```sh
Running 30s test @ http://localhost:8080/get/1
  100 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.18ms   11.05ms 153.56ms   88.58%
    Req/Sec     1.49k   755.47    14.28k    91.03%
  4114447 requests in 30.10s, 690.60MB read
  Socket errors: connect 82, read 0, write 0, timeout 0
  Non-2xx or 3xx responses: 4114447
Requests/sec: 136699.70
Transfer/sec:     22.94MB
```


## chi

```sh
Running 30s test @ http://localhost:8080/get/1
  100 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.38ms   11.74ms 182.56ms   88.14%
    Req/Sec     1.53k   779.60    14.78k    90.57%
  4208703 requests in 30.10s, 706.42MB read
  Socket errors: connect 82, read 0, write 0, timeout 0
  Non-2xx or 3xx responses: 4208703
Requests/sec: 139844.14
Transfer/sec:     23.47MB
```