# staticcheck Go代码检查利器

[staticcheck GitHub](https://github.com/dominikh/go-tools)

Staticcheck是Go编程语言的最先进的linter。通过静态分析，它可以发现错误和性能问题，提供简化服务，并强制执行样式规则。

## Windows安装

[官方地址](https://github.com/dominikh/go-tools/releases)

下载对应文件，解压后得到  `staticcheck.exe` 可执行文件，放到 `$GOPATH/bin` 目录下


## 使用教程

https://play-with-go.dev/using-staticcheck_go115_en/


### 查看版本 `staticcheck -version`

### 代码检查 `staticcheck .` 
`cd`到相关代码目录，执行  `staticcheck .` 就会检查当前目录下的代码

**示例：**
```sh
$ staticcheck .
pets.go:23:14: Printf format %v reads arg #1, but call has only 0 args (SA5009)
pets.go:25:10: should use fmt.Errorf(...) instead of errors.New(fmt.Sprintf(...)) (S1028)
pets.go:30:7: receiver name should be a reflection of its identity; don't use generic names such as "this" or "self" (ST1006)
pets.go:31:9: the argument is already a string, there's no need to use fmt.Sprintf (S1025)

```

正如你从输出中看到的，Staticcheck报告的错误很像Go编译器。每一行代表一个问题，从文件位置开始，然后是问题的描述，行末括号里是Staticcheck检查号。

Staticcheck检查分为不同的类别，每个类别用不同的代码前缀来标识。下面列出一些。

- 代码简化 `S1XXX`
- 正确性问题 `SA5XXX`
- 风格问题 `ST1XXX`
[Staticcheck官方网站](https://staticcheck.io/docs/checks)列出并记录了所有的类别和检查。许多检查甚至有例子。你也可以使用`-explain`标志在命令行获得细节。

```sh
$ staticcheck -explain SA5009
Invalid Printf call

Available since
    2019.2
```
### 配置 Staticcheck
Staticcheck在提供一些合理的、经过实践检验的默认值。然而，Staticcheck的各个方面都可以通过配置文件进行自定义。

你也在Staticcheck网站上注意到，检查代号`ST1000`(包注释)恰好涵盖了这种情况，但默认情况下它没有启用。

Staticcheck的配置文件名为`staticcheck.conf`，可以填写自定义配置，比如填写开启包注释检查：
```sh
checks = ["inherit", "ST1000"]
```

让我们创建一个Staticcheck配置文件来启用check ST1000，继承Staticcheck的默认值。

### 忽略错误 `lint:ignore`
在某些情况下，我们不需要 Staticcheck 检查某个"错误"——比如人为制造的，可以使用`lint:ignore`来指定某些代码无需检查。
#### 忽略单行 `line-based`

```go
//lint:ignore SA4018 trying out line-based linter directives
food = food
```
增加注释后，`food = food`这行代码便不会被检查

### 忽略整个文件 `file-based`
如果要忽略很多行代码呢？像使用工具生成的代码变不需要检查，比如`gRPC`生成的 `XX.pb.go`代码，可以在代码前面加入以下注释
```go
//lint:file-ignore SA4018 trying out file-based linter directives
```
**示例：**
```go
// Package pets contains useful functionality for pet owners
package pets

import (
	"fmt"
)

//lint:file-ignore SA4018 trying out file-based linter directives

type Animal int

const (
	Dog Animal = iota
	Snake
)

type Pet struct {
	Kind Animal
	Name string
}

func (p Pet) Walk() error {
	switch p.Kind {
	case Dog:
		fmt.Printf("Will take %v for a walk around the block\n", p.Name)
	default:
		return fmt.Errorf("cannot take %v for a walk", p.Name)
	}
	return nil
}

func (p Pet) Feed(food string) {
	food = food
	fmt.Printf("Feeding %v some %v\n", p.Name, food)
}

func (p Pet) String() string {
	return p.Name
}
```

整个文件的代码便不会被检查，在 `pet.go`目录下执行 `staticcheck .`是没有任何建议输出的。

