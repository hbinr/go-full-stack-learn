# 生成代码


## 知识点
可以通过定义好的.proto文件来生成Go,ava,Python,C++, Ruby, JavaNano, Objective-C,或者C# 代码，需要基于.proto文件运行protocol buffer编译器protoc。

如果你没有安装编译器，下载安装包并遵照README安装。

对于Go,你还需要安装一个特殊的代码生成器插件。你可以通过GitHub上的protobuf库找到安装过程

通过如下方式调用protocol编译器（各语言全量命令，根据需要选择）：

```go

protoc --proto_path=IMPORT_PATH 
--cpp_out=DST_DIR 
--java_out=DST_DIR 
--python_out=DST_DIR 
--go_out=DST_DIR 
--ruby_out=DST_DIR 
--javanano_out=DST_DIR 
--objc_out=DST_DIR 
--csharp_out=DST_DIR 

path/to/file.proto
```

**各命令参数解释：**
- IMPORT_PATH声明了一个.proto文件所在的解析import具体目录。如果忽略该值，则使用当前目录。如果有多个目录则可以多次调用--proto_path，它们将会顺序的被访问并执行导入。-I=IMPORT_PATH是--proto_path的简化形式。
- 当然也可以提供一个或多个输出路径：
- --cpp_out 在目标目录DST_DIR中产生C++代码，可以在C++代码生成参考中查看更多。
- --java_out 在目标目录DST_DIR中产生Java代码，可以在 Java代码生成参考中查看更多。
- --python_out 在目标目录 DST_DIR 中产生Python代码，可以在Python代码生成参考中查看更多。
- --go_out 在目标目录 DST_DIR 中产生Go代码，可以在[GO代码生成](https://developers.google.com/protocol-buffers/docs/reference/go-generated?hl=zh-cn)参考中查看更多。
- --ruby_out在目标目录 DST_DIR 中产生Go代码，参考正在制作中。
- --javanano_out在目标目录DST_DIR中生成JavaNano，JavaNano代码生成器有一系列的选项用于定制自定义生成器的输出：你可以通过生成器的README查找更多信息，JavaNano参考正在制作中。
- --objc_out在目标目录DST_DIR中产生Object代码，可以在Objective-C代码生成参考中查看更多。
- --csharp_out在目标目录DST_DIR中产生Object代码，可以在C#代码生成参考中查看更多。
- --php_out在目标目录DST_DIR中产生Object代码，可以在PHP代码生成参考中查看更多。



## 案例：以生成go代码为例

**前提是要安装 protoc 工具**([未安装的可参考](../docs/01_grpc-go安装.md))

[具体实践请看](../docs/03_初步使用.md)