## wire 安装

### install
在终端运行以下命令：
```go
go get github.com/google/wire/cmd/wire
```
确保你的 `$GOPATH/bin` 被加入到了 `$PATH`中

### 验证
在终端输入命令 `wire`：
```go
hblock@hblock:~$ wire
wire: go [list -e -json -compiled=true -test=false -export=false -deps=true -find=false -tags=wireinject -- .]: exit status 1: go: cannot find main module; see ’go help modules‘

wire: generate failed
```

有上述提示表示安装成功
