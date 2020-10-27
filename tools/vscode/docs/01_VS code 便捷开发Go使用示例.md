# VS code 便捷开发Go使用示例

主要记录：
- Go 插件使用
- 常用开发使用及快捷键操作

Go 插件使用操作方式：
- 命令式的操作：
  1. 选中需要操作的内容
  2. 按`ctrl+shift+p`，输入：`go:相关命令`
  3. 回车，然后输入对应命令需要的参数。没有参数的就回车生效
   
- 手动鼠标点
  1. 选中内容
  2. 右键弹出相关操作，选择go相关命令，如`Go:Generate Unit Tests For Function`
  3. 点击命令，然后输入对应命令需要的参数。没有参数的就生效

## 自动生成测试用例

### 命令：`Go:Generate Unit Tests For Function`

### 示例

测试代码：
```go
func Add(a, b int) int {
	return a + b 
}
```
操作命令后，会在函数所在文件的同级目录下生成测试代码：
```go

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
```
`TODO: Add test cases.`需要你填写自己的测试用例

不过要值的注意的是，在测试函数中，使用`t.Log()`函数是不会在终端输出内容的，需要加一个`-v`参数：
1. 在扩展商店里搜已经安装了的`go`插件，点击右下角的 **设置图标**
2. 选中扩展设置
3. 找到`Go:Build Flags`，添加 `-v`项，如下：

![](https://img-blog.csdnimg.cn/20201027181422863.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzQ3NDA0MTgx,size_16,color_FFFFFF,t_70#pic_center)

这样在执行`TestXXX`函数的时候，就能在终端输出`log`的内容了

## 自动生成结构体实例化
### 命令：`Go:fill Struct`

### 示例
在上面生成的单元测试中，在需要添加的测试case的地方来自动生成结构体实例化

1. 先敲一个`{}`，任意结构体的实例化也是这样，如:`u := &User{}`，光标在`{}`中，然后执行命令操作即可。
2. 然后命令操作即可

生成的代码如下：
```go
tests := []struct {
		name string
		args args
		want int
	}{
		{
		name: "",
		args: args{
			a: 0,
			b: 0,
		},
		want: 0,
	}
	}
```
不过要注意一点：

由于`tests`的类型是一个结构体切片`[]struct`，所以在生成的代码后面需要手动补一个逗号`,`，按保存就自动格式化代码了

这个经常用，我设置了快捷键 `alt + f`

## 自动实现接口
### 命令：`Go:Generate Interface Stubs`

该命令需要参数，输入三个内容：
- 方法接受者形参名eg：`s`
- 方法接受者名称，即实现类的名称，eg：`*Student`，你可以指定是值类型还是引用类型
- 要实现的接口的名称，需要加上包名eg：`code.Speaker`

完整内容：`s *Student code.Speaker`

### 示例
```go
package code

type Speaker interface {
	// Speak speak action
	Speak()
}

type Student struct {
}

// Speak speak action
func (s *Student) Speak() {
	panic("not implemented") // TODO: Implement
}
```
可以看到，连注释也顺便帮忙生成了。


## 自动增加/删除tag
### 增加命令：`Go:Add Tags To Struct Fileds`
### 删除命令：`Go:Remove Tags From Struct Fileds`

选中你要生成`tag`的字段，执行命令（增加和删除都是同理，需要选中字段，只会执行已选中的字段）

默认是只生成 `json tag`，可以自定义。在`setting.json`，加入`go.addTags`设置即可

**设置示例：**
```go
 // 结构体tag 设置
  "go.addTags": {
    // 可配置多个tag json,orm
    "tags": "json,form",
    // options 可以填入json=omitempty
    "options": "",
    "promptForTags": false,
    // snakecase:下划线分隔， camelcase:驼峰命名
    "transform": "camelcase"
  },
```

这个我也常用，设置快捷键`alt+a`，删除不常用，未设置。
## 快速导入包
#### 命令：`Go:Add import`
一般情况下，写代码智能提示或者保存的时候就会自动引入包。但是有时候vscode在引入同名的包时，会引入错误的包，这种情况主要是本地`pkg->mod`目录下有了重名的库，vscode无法知道是哪个。

这时候需要我们手动引入，执行命令，然后输入包名，从展示列表中选中你想要的回车即可，支持模糊搜索

这个常用，我也设置了快捷键：`alt+i`
## 查找接口的具体实现

1. 鼠标移到接口定义处
2. 快捷`ctrl+f12`
3. 或者右键，选中：`Go to implementations`

该操作会出现两种情况：
- 只有一个实现，那么直接跳转到实现的结构体上
- 如果有多个实现，那么会弹出一个界面，左侧框展示实现的代码，右侧框展示实现的列表，二者是联动的。然后选择你想看的实现，双击就会跳转



## 重构
### 重命名
1. 选中你想要重构的字段、方法、接口名等等
2. 按`F2`，然后输入你想要的命名

这个操作会将凡是引用该字段、方法、接口的地方全部重命名，并且支持跨文件

### 字段提取
#### 命令：`Go:Extract to variable`

字段提取主要用于**判断条件复杂**的场景，如果该条件判断在多个地方使用，最好是抽离出来，提取成一个变量

操作：
- 选中需要提取的内容
- 执行命令
- 输入要生成的变量名

等待1s，生成代码

### 函数提取
#### 命令：`Go:Extract to function`

函数提取主要用于**逻辑可复用**的场景。把同一段逻辑抽离成一个函数

操作：
- 选中需要提取的内容
- 执行命令
- 输入要生成的函数名

等待1s，生成代码


#### 示例
原代码：
```go
func ExtractFuncTest(a, b, c int) {
	if a > 0 {
	}
	if b > 0 {
	}
	if c > 0 {
	}
}
```
选中里面的逻辑，函数提取后的代码：
```go
func ExtractFuncTest(a, b, c int) {
	flag(a, b, c)
}

func flag(a int, b int, c int) {
	if a > 0 {
	}
	if b > 0 {
	}
	if c > 0 {
	}
}
```

## 第三方库加入到当前的workspace 
一般vscode左侧只会展示当前项目的目录，引入的第三方库并不会展示出来。并不像`Goland`一样，有个`External Libary`目录

这个功能不太好用，目前我只能导入一些`go`的库，如`fmt`,`errors`库，第三方库并不行。待解决

## 设置Go相关命令，右键就能显示常用的

在`settings.json`中，输入`go.editorContextMenuCommands`，事实上，输入前几个字母就会智能提示，回车后就会生成配置，如下：
```json
"go.editorContextMenuCommands": {
"toggleTestFile": true,
"addTags": false,
"removeTags": true,
"testAtCursor": true,
"testFile": true,
"testPackage": false,
"generateTestForFunction": true,
"generateTestForFile": false,
"generateTestForPackage": false,
"addImport": false,
"testCoverage": true,
"playground": true,
"debugTestAtCursor": true
},
```
`true`表示开启，设置完后，重启VS code即可

参考：

[B站——【教程】vscode-go插件的这些用法，你真的知道么?](https://www.bilibili.com/video/av94727284/)

