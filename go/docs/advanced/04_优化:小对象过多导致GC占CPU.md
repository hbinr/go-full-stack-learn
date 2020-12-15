# 优化:小对象过多导致GC占CPU

## 问题背景
- 随着流量增大，请求超时增多
- 耗时毛刺严重，99分位耗时较长
- 总体内存变化不大

## 排查

查看多有对象+被引用的对象
```go
go tool pprof -alloc_objects

go tool pprof -inuse_objects

```
查看CPU占用情况，发现GC扫描函数占大量CPU，runtime.scanobject等
```go
go tool pprof bin/dupsdc
```
## 问题根源

小对象过多引起的服务吞吐问题


## 解决
思路：减少对象分配

对象如：`string`、 `map[key]value`、 `slice`、 
`*Type`

```go

原数据结构体                            优化数据结构                        优化点

map[string]SampleStruct        map[[32]byte]SampleStruct           key使用值类型避免对map遍历

map[int]*SampleStruct          map[int]SampleStruct                value使用值类型避免对map遍历

simpleSlice []float64           simpleSlice [32]float64             值类型代替对象类型
```




