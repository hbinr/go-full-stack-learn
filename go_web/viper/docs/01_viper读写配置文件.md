# viper 基本使用

[Viper](https://github.com/spf13/viper)是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。

## 安装
```go
go get github.com/spf13/viper
```

## 什么是Viper？

Viper是适用于Go应用程序（包括Twelve-Factor App）的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。它支持以下特性：

- 设置默认值
- 从JSON、TOML、YAML、HCL、envfile和Java properties格式的配置文件读取配置信息
- 实时监控和重新读取配置文件（可选）
- 从环境变量中读取
- 从远程配置系统（etcd或Consul）读取并监控配置变化
- 从命令行参数读取配置
- 从buffer读取配置
- 显式配置值

## 为什么选择Viper?

在构建现代应用程序时，你无需担心配置文件格式；你想要专注于构建出色的软件。Viper的出现就是为了在这方面帮助你的。

Viper能够为你执行下列操作：

- 查找、加载和反序列化JSON、TOML、YAML、HCL、INI、envfile和Java properties格式的配置文件。
- 提供一种机制为你的不同配置选项设置默认值。
- 提供一种机制来通过命令行参数覆盖指定选项的值。
- 提供别名系统，以便在不破坏现有代码的情况下轻松重命名参数。
- 当用户提供了与默认值相同的命令行或配置文件时，可以很容易地分辨出它们之间的区别。

Viper会按照下面的优先级。每个项目的优先级都高于它下面的项目:

- 显示调用Set设置值
- 命令行参数（flag）
- 环境变量
- 配置文件
- key/value存储
- 默认值


**重要： 目前Viper配置的键（Key）是大小写不敏感的。目前正在讨论是否将这一选项设为可选。**

## 把值存入Viper


### 建立默认值

一个好的配置系统应该支持默认值。键不需要默认值，但如果没有通过配置文件、环境变量、远程配置或命令行标志（flag）设置键，则默认值非常有用。

例如：
```go
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
```

Viper是开箱即用的。你不需要配置或初始化即可开始使用Viper。由于大多数应用程序都希望使用单个中央存储库管理它们的配置信息，所以viper包提供了这个功能。它类似于单例模式。

当然你可以使用多个viper实例。每个都有自己独特的一组配置和值。每个人都可以从不同的配置文件，key value存储区等读取数据。每个都可以从不同的配置文件、键值存储等中读取。viper包支持的所有功能都被镜像为viper实例的方法。

```go
x := viper.New()
y := viper.New()

x.SetDefault("ContentDir", "content")
y.SetDefault("ContentDir", "foobar")

//...
```

### 读取配置文件

Viper需要最少知道在哪里查找配置文件的配置。Viper支持`JSON`、`TOML`、`YAML`、`HCL`、`envfile`和`Java properties`格式的配置文件。Viper可以搜索多个路径，但目前单个Viper实例只支持单个配置文件。Viper不默认任何配置搜索路径，将默认决策留给应用程序。

下面是一个如何使用Viper搜索和读取配置文件的示例。不需要任何特定的路径，但是至少应该提供一个配置文件预期出现的路径。
```go
viper.SetConfigFile("./config.yaml") // 指定配置文件路径
viper.SetConfigName("config") // 配置文件名称(无扩展名)
viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
viper.AddConfigPath(".")               // 还可以在工作目录中查找配置
err := viper.ReadInConfig() // 查找并读取配置文件
if err != nil { // 处理读取配置文件的错误
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}
```
在加载配置文件出错时，你可以像下面这样处理找不到配置文件的特定情况：
```go
if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        // 配置文件未找到错误；如果需要可以忽略
    } 
    // 配置文件被找到，但产生了另外的错误
    
}

// 配置文件找到并成功解析
```

当你使用如下方式读取配置时，viper会从`./conf`目录下查找任何以`config`为文件名的配置文件，如果同时存在`./conf/config.json`和`./conf/config.yaml`两个配置文件的话，viper会从哪个配置文件加载配置呢？

```go
viper.SetConfigName("config")
viper.AddConfigPath("./conf")
```

它会先加载 `config.json`配置文件，因为其在`/conf`目录下的排序位置靠前，如果在`config.json`中找到需要的配置，如:`port:8080`，那么就会正常输出，不再读取之后的配置文件。

如果没有找到的对应的配置，那么就会报错，中断程序，错误信息如下：

> While parsing config: unexpected end of JSON input

如果增加`viper.SetConfigType("yaml")`来指定配置文件类型，这样是否就能正确获得配置信息？

**答案是否定的**，因为其使用常场景是获取远程配置信息的。比如，配置中心有个名为 `config.yaml`，远程获取配置的时候需要解析网络传输中的字节流，但是以什么格式解析呢？是`JSON`还是`yaml`？这时`viper.SetConfigType("yaml")`就派上用场了，它告诉调用方使用`yaml`格式解析配置信息。

那上述同名配置的场景如何正确的获取配置信息呢？

**可以指定配置文件名+文件后缀**，`viper.SetConfigFile("config.yaml")`，使用该方法后，便不需要配置 `SetConfigName`和`SetConfigType`，它会直接寻找文件名为 `config.yaml`的文件。

### 从io.Reader读取配置

Viper预先定义了许多配置源，如文件、环境变量、标志和远程K/V存储，但你不受其约束。你还可以实现自己所需的配置源并将其提供给viper。

```go
viper.SetConfigType("yaml") // 或者 viper.SetConfigType("YAML")

// 任何需要将此配置添加到程序中的方法。
var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

viper.ReadConfig(bytes.NewBuffer(yamlExample))

viper.Get("name") // 这里会得到 "steve"
```


### 写入配置文件
从配置文件中读取配置文件是有用的，但是有时你想要存储在运行时所做的所有修改。为此，可以使用下面一组命令，每个命令都有自己的用途:
- WriteConfig - 将当前的viper配置写入预定义的路径并覆盖（如果存在的话）。如果没有预定义的路径，则报错。
- SafeWriteConfig - 将当前的viper配置写入预定义的路径。如果没有预定义的路径，则报错。如果存在，将不会覆盖当前的配置文件。
- WriteConfigAs - 将当前的viper配置写入给定的文件路径。将覆盖给定的文件(如果它存在的话)。
- SafeWriteConfigAs - 将当前的viper配置写入给定的文件路径。不会覆盖给定的文件(如果它存在的话)。

根据经验，标记为safe的所有方法都不会覆盖任何文件，而是直接创建（如果不存在），而默认行为是创建或截断。

**一个小示例：**
```go
viper.WriteConfig() // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // 因为该配置文件写入过，所以会报错
viper.SafeWriteConfigAs("/path/to/my/.other_config")
```

**实战练习：**
```go
// WriteConfigDemo viper写入配置文件示例
func WriteConfigDemo() {
	viper.SetConfigFile("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

    viper.Set("version", 1.5) // 修改配置

    // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
	if err := viper.WriteConfig(); err != nil {
		fmt.Println("viper SafeWriteConfig failed,err:", err)
		return
    }
	
	// viper.GetString("version") 获取配置的值
	fmt.Println("viper write config success, version's value:", viper.GetString("version"))
}
```

## 监控并重新读取配置文件

Viper支持在运行时实时读取配置文件的功能。

需要重新启动服务器以使配置生效的日子已经一去不复返了，viper驱动的应用程序可以在运行时读取配置文件的更新，而不会错过任何消息。

只需告诉viper实例watchConfig。可选地，你可以为Viper提供一个回调函数，以便在每次发生更改时运行。

```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
  // 配置文件发生变更之后会调用的回调函数
	fmt.Println("Config file changed:", e.Name)
})
```

**注意：**

- 确保在调用`WatchConfig()`之前添加了所有的配置路径
- `fsnotify` 使用的库是 `github.com/fsnotify/fsnotify`

**实战练习：**

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// WatchConfigDemo viper监控配置文件示例
func WatchConfigDemo() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	// 查找并读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			fmt.Println("未找到配置文件，err:", err)
			return
		}
		// 配置文件被找到，但产生了另外的错误
		fmt.Println("配置文件被找到，但产生了另外的错误，err:", err)
		return
	}

	// 实时监控配置
	viper.WatchConfig()
	// 配置文件发生变更之后会调用回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// 验证监控配置示例
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})
	r.Run()
}

func main(){
    WatchConfigDemo()
}
```