







[Why I Recommend to Avoid Using the go-kit Library](https://gist.github.com/posener/330c2b08aaefdea6f900ff0543773b2e)




# 观点

### 可控制
实际上，我之所以从go-micro转到go-kit，是因为你可以完全控制从传输到服务。我很沮丧，不知道为什么我的Java GRPC服务不能调用我的go-micro服务，因为那个漂亮的简单的API隐藏了各种魔法，而这些魔法对我来说并不适用，我无法控制它。Go-Kit虽然比较啰嗦，但可以轻松调试任何问题，从而在做复杂的编排时节省开发时间。


## 分层规范，利于团队
这个想法是要有一个标准化的api和简洁的架构设计。我喜欢go-kit，因为它强制执行了团队中的秩序和协议。一旦你了解了这些概念，就不会再有黑客和破坏依赖关系的行为。当你运行一个大型项目时，有序的事情和创造最佳实践才是最重要的。

我喜欢传输和实现之间的干净分离，较小的代码块做一件事，而不是混合不同的关注点。

当你来构建你的服务时，中间件是一个好处，就像乐高一样，你可以在不改变核心服务代码的情况下添加仪器和跟踪层。

使代码可测试是一个巨大的好处.注入依赖和抽象接口是使它成为可能的原因。

是的，如果你运行的是一个非常小的项目，那就是矫枉过正了，否则，我推荐它。

# 额外获取到的知识

## 基于`.proto`文件生成代码
> The idea, which I kind of like is to define the services in a protobuf file (as you would do for gRPC)

> This project does this really nicely. https://github.com/moul/protoc-gen-gotemplate
And the examples show some go-kit integration:

> - https://github.com/moul/kafka-gateway/
> - https://github.com/moul/translator
> - https://github.com/moul/acl
