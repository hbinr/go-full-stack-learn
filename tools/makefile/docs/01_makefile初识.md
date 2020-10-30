# Makefile初识

借助`Makefile`我们在编译过程中不再需要每次手动输入编译的命令和编译的参数，可以极大简化项目编译过程。
[Makefile教程](http://c.biancheng.net/view/7097.html)
## make介绍
`make`是一个构建自动化工具，会在当前目录下寻找`Makefile`或`makefile`文件。如果存在相应的文件，它就会依据其中定义好的规则完成构建任务。
## Makefile介绍
我们可以把`Makefile`简单理解为它定义了一个项目文件的编译规则。借助`Makefile`我们在编译过程中不再需要每次手动输入编译的命令和编译的参数，可以极大简化项目编译过程。同时使用`Makefile`也可以在项目中确定具体的编译规则和流程，很多开源项目中都会定义`Makefile`文件。

## 规则概述
`Makefile`由多条规则组成，每条规则主要由两个部分组成，分别是依赖的关系和执行的命令。

其结构如下所示：
```sh
[target] ... : [prerequisites] ...
<tab>[command]
    ...
    ...
```
其中：

- targets：规则的目标
- prerequisites：可选的要生成 targets 需要的文件或者是目标。
- command：make 需要执行的命令（任意的 shell 命令）。可以有多条命令，每一条命令占一行。

举个例子：
```sh
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o xx
```

## 语法

### @
表示`@`之后的命令内容不会显示在终端，只执行命令。不加则会先在终端输出命令内容，然后再执行去命令
### ${}
`${}`读取其包裹的内容。
```sh
NAME="test"

run:
    @echo ${NAME}
```
在命令行执行`make run`，则会输出 test
### .PHONY
`.PHONY`用来定义伪目标。不创建目标文件，而是去执行这个目标下面的命令

```sh
.PHONY: all build run gotool

NAME="test"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./
```
在终端执行`make all`就会执行：`gotool build`命令

如果不加`.PHONY`呢？命令如下：
```sh
BINARY="bluebell"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./
```

make开始执行`Makefile`定义的命令，同时会自动扫描当前目录及文件名，如果遇到与命令同名的不去执行定义的命令了。

因此，最好使用`.PHONY`用来定义伪目标，保证命令正确的执行。

待补充......  学多少，补多少