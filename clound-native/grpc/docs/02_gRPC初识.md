# gRPC 是什么？
在 gRPC 里客户端应用可以像调用本地对象一样直接调用另一台不同的机器上服务端应用的方法，使得您能够更容易地创建分布式应用和服务。

与许多 RPC 系统类似，gRPC 也是基于以下理念：
>定义一个服务，指定其能够被远程调用的方法（包含参数和返回类型）。在服务端实现这个接口，并运行一个 gRPC 服务器来处理客户端调用。在客户端拥有一个存根能够像服务端一样的方法。

gRPC 客户端和服务端可以在多种环境中运行和交互——如从 google 内部的服务器到你自己的电脑，并且可以用任何 gRPC 支持的语言来编写。

所以，你可以很容易地用 Java 创建一个 gRPC 服务端，用 Go、Python、Ruby 来创建客户端。此外，Google 最新 API 将有 gRPC 版本的接口，使你很容易地将 Google 的功能集成到你的应用里。


# 使用 protocol buffers
gRPC 默认使用 `protocol buffers`，这是 Google 开源的一套成熟的结构数据序列化机制（当然也可以使用其他数据格式如 JSON）。正如你将在下方例子里所看到的，你用 `proto files` 创建 gRPC 服务，用 `protocol buffers` 消息类型来定义方法参数和返回类型。

## Protocol buffers 版本

尽管 `protocol buffers` 对于开源用户来说已经存在了一段时间。

例子内使用的却一种名叫 proto3 的新风格的 `protocol buffers`，它拥有轻量简化的语法、一些有用的新功能，并且支持更多新语言。

在golang/protobuf Github 源码库里还有针对 Go 语言的生成器， 对更多语言的支持正在开发中。

你可以在 [proto3语言指南](https://blog.csdn.net/hulinku/article/details/80827018)里找到更多内容， 在与当前默认版本的发布说明比较，看到两者的主要不同点。
