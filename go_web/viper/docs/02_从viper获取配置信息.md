# 从viper获取配置信息

在Viper中，有几种方法可以根据值的类型获取值。存在以下功能和方法:

- `Get(key string) : interface{}`
- `GetBool(key string) : bool`
- `GetFloat64(key string) : float64`
- `GetInt(key string) : int`
- `GetIntSlice(key string) : []int`
- `GetString(key string) : string`
- `GetStringMap(key string) : map[string]interface{}`
- `GetStringMapString(key string) : map[string]string`
- `GetStringSlice(key string) : []string`
- `GetTime(key string) : time.Time`
- `GetDuration(key string) : time.Duration`
- `IsSet(key string) : bool`
- `AllSettings() : map[string]interface{}`

需要认识到的一件重要事情是，每一个Get方法在**找不到值**的时候都会**返回零值**。为了检查给定的键是否存在，提供了`IsSet()`方法。

## 访问简单的key

有如下配置文件：
```go
{
  "version": 1.5
}
```

示例：
```go
viper.GetString("version") // 不区分大小写的设置和获取
```

## 访问嵌套的key
访问器方法也接受深度嵌套键的格式化路径。例如，如果加载下面的JSON文件：
```go
{
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}
```

Viper可以通过传入`.`分隔的路径来访问嵌套字段：
```go
viper.GetString("datastore.metric.host") // (返回 "127.0.0.1")
```

搜索路径将遍历其余配置注册表，直到找到为止。因为Viper支持从多种配置来源，例如 **磁盘上的配置文件>命令行标志位>环境变量>远程Key/Value存储>默认值**，我们在查找一个配置的时候如果在当前配置源中没找到，就会继续从后续的配置源查找，直到找到为止。

例如，在给定此配置文件的情况下，`datastore.metric.host`和`datastore.metric.port`均已定义（并且可以被覆盖）。如果另外在默认值中定义了`datastore.metric.protocol`，Viper也会找到它。

然而，如果`datastore.metric`被直接赋值覆盖（被flag命令行，环境变量，`set()`方法等等…），那么`datastore.metric`的所有子键都将变为未定义状态，它们被高优先级配置级别“遮蔽”（shadowed）了。

最后，如果存在**与分隔的键路径匹配**的键，则返回其值。例如：

```go
{
    "datastore.metric.host": "0.0.0.0",
    "host": {
        "address": "localhost",
        "port": 5799
    },
    "datastore": {
        "metric": {
            "host": "127.0.0.1",
            "port": 3099
        },
        "warehouse": {
            "host": "198.0.0.1",
            "port": 2112
        }
    }
}

viper.GetString("datastore.metric.host") // 返回 "0.0.0.0"
```

## 提取子树
从Viper中提取子树。

例如，viper实例现在代表了以下配置：
```yaml
app:
  cache1:
    max-items: 100
    item-size: 64
  cache2:
    max-items: 200
    item-size: 80
```

执行后：
```go
subv := viper.Sub("app.cache1")
```

subv现在就代表：

```yaml
max-items: 100
item-size: 64
```

假设我们现在有这么一个函数：
```go
func NewCache(cfg *Viper) *Cache {...}
```

它基于subv格式的配置信息创建缓存。现在，可以轻松地分别创建这两个缓存，如下所示：

```go
cfg1 := viper.Sub("app.cache1")
cache1 := NewCache(cfg1)

cfg2 := viper.Sub("app.cache2")
cache2 := NewCache(cfg2)
```
## 反序列化

你还可以选择将所有或特定的值解析到结构体、map等，这样viper加载完配置信息后就可以使用结构体变量保存配置信息。

有两种方法可以做到这一点：

```go
Unmarshal(rawVal interface{}) : error
UnmarshalKey(key string, rawVal interface{}) : error
```

举个例子：
```go

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host   string `mapstructure:"host"`
	DbName string `mapstructure:"dbname"`
	Port   int    `mapstructure:"port"`
}

func main() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("can't find config file, err:", err)
			return
		}
		fmt.Println("viper.ReadInConfig failed, err:", err)
	}
    var c Config
    // Unmarshal 支持解析到嵌入的结构体
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Println("viper.Unmarshal failed, err:", err)
	}
	fmt.Printf("viper.Unmarshal success, config:%#v\n", c)
}
```

Viper在后台使用github.com/mitchellh/mapstructure来解析值，其默认情况下使用`mapstructure` tag。

**注意：**
当我们需要将viper读取的配置反序列到我们定义的结构体变量中时，一定要使用`mapstructure` tag！
