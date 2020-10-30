# 为go项目编写Makefile

在`main.go`所在目录下新建一个 `Makefile`文件，填写以下内容:
```sh
.PHONY: all build run gotool clean help

BINARY="test"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
```

解析：
- `.PHONY`后面的命令都可以通过`make`+ 命令来单独执行，如`make build`
- `all` 执行`gotool build`命令
- `build` 指定了操作系统类型和CPU类型，然后执行`go build -o ${BINARY}`命令，生成指定的名称，即test
- `run` 执行`@go run ./`命令，因为有`@`，所以不会在终端输出`go run ./`命令
- `gotool`  分别是格式化代码和代码风格、严谨型校验
- `clean` 删除生成的二进制文件名，如果存在的话
- `help`  输出帮助文档