## RabbitMQ 简介

在介绍 RabbitMQ 之前实现要介绍一下 MQ，MQ 是什么？

MQ 全称是 Message Queue，可以理解为消息队列的意思，简单来说就是消息以管道的方式进行传递。

RabbitMQ 是一个实现了 AMQP（Advanced Message Queuing Protocol）高级消息队列协议的消息队列服务，用 Erlang 语言的。

## 使用场景

在我们秒杀抢购商品的时候，系统会提醒我们稍等排队中，而不是像几年前一样页面卡死或报错给用户。

像这种排队结算就用到了消息队列机制，放入通道里面一个一个结算处理，而不是某个时间断突然涌入大批量的查询新增把数据库给搞宕机，所以 RabbitMQ 本质上起到的作用就是削峰填谷，为业务保驾护航。

## 为什么选择 RabbitMQ

现在的市面上有很多 MQ 可以选择，比如 ActiveMQ、ZeroMQ、Appche Qpid，那问题来了为什么要选择 RabbitMQ？

- 除了 Qpid，RabbitMQ 是唯一一个实现了 AMQP 标准的消息服务器；
- 可靠性，RabbitMQ 的持久化支持，保证了消息的稳定性；
- 高并发，RabbitMQ 使用了 Erlang 开发语言，Erlang 是为电话交换机开发的语言，天生自带高并发光环，和高可用特性；
- 集群部署简单，正是应为 Erlang 使得 RabbitMQ 集群部署变的超级简单；
- 社区活跃度高，根据网上资料来看，RabbitMQ 也是首选；

## 工作机制

### 生产者、消费者和代理

在了解消息通讯之前首先要了解 3 个概念：生产者、消费者和代理。

**生产者：** 消息的创建者，负责创建和推送数据到消息服务器；

**消费者：** 消息的接收方，用于处理数据和确认消息；

**代理：** 就是 RabbitMQ 本身，用于扮演“快递”的角色，本身不生产消息，只是扮演“快递”的角色。

### 消息发送原理

首先你必须连接到 Rabbit 才能发布和消费消息，那怎么连接和发送消息的呢？

你的应用程序和 Rabbit Server 之间会创建一个 TCP 连接，一旦 TCP 打开，并通过了认证，认证就是你试图连接 Rabbit 之前发送的 Rabbit 服务器连接信息和用户名和密码，有点像程序连接数据库。

后面代码会详细介绍使用 Go 进行连接认证，一旦认证通过你的应用程序和 Rabbit 就创建了一条 AMQP 信道（Channel）。

信道是创建在“真实”TCP 上的虚拟连接，AMQP 命令都是通过信道发送出去的，每个信道都会有一个唯一的 ID，不论是发布消息，订阅队列或者介绍消息都是通过信道完成的。

### 为什么不通过 TCP 直接发送命令？

对于操作系统来说创建和销毁 TCP 会话是非常昂贵的开销，假设高峰期每秒有成千上万条连接，每个连接都要创建一条 TCP 会话，这就造成了 TCP 连接的巨大浪费，而且操作系统每秒能创建的 TCP 也是有限的，因此很快就会遇到系统瓶颈。

如果我们每个请求都使用一条 TCP 连接，既满足了性能的需要，又能确保每个连接的私密性，这就是引入信道概念的原因。

![](http://icdn.apigo.cn/blog/rabbit_channel.png)

## 你必须知道的 Rabbit

想要真正的了解 Rabbit 有些名词是你必须知道的。

包括：ConnectionFactory（连接管理器）、Channel（信道）、Exchange（交换器）、Queue（队列）、RoutingKey（路由键）、BindingKey（绑定键）。

**ConnectionFactory（连接管理器）**：应用程序与 Rabbit 之间建立连接的管理器，程序代码中使用；

**Channel（信道）**：消息推送使用的通道；

**Exchange（交换器）**：用于接受、分配消息；

**Queue（队列）**：用于存储生产者的消息；

**RoutingKey（路由键）**：用于把生成者的数据分配到交换器上；

**BindingKey（绑定键）**：用于把交换器的消息绑定到队列上，通过不同方式的绑定，可以实现 rabbitmq 不同的工作模式

看到上面的解释，最难理解的路由键和绑定键了，那么他们具体怎么发挥作用的，请看下图：

![](http://icdn.apigo.cn/blog/rabbit-producer.gif)

## 消息持久化

Rabbit 队列和交换器有一个不可告人的秘密，就是默认情况下重启服务器会导致消息丢失，那么怎么保证 Rabbit 在重启的时候不丢失呢？答案就是消息持久化。

当你把消息发送到 Rabbit 服务器的时候，你需要选择你是否要进行持久化，但这并不能保证 Rabbit 能从崩溃中恢复，想要 Rabbit 消息能恢复必须满足以下条件：
TODO1

- 投递消息的时候 `durable` 设置为 true，消息持久化，代码：`channel.queueDeclare(x, true, false, false, null)`，参数 2 设置为 true 持久化；
- 设置投递模式 `deliveryMode` 设置为 2（持久），代码：`channel.basicPublish(x, x, MessageProperties.PERSISTENT_TEXT_PLAIN,x)`，参数 3 设置为存储纯文本到磁盘；
- 消息已经到达持久化交换器上；
- 消息已经到达持久化的队列；

### 持久化工作原理

Rabbit 会将你的持久化消息写入磁盘上的持久化日志文件，等消息被消费之后，Rabbit 会把这条消息标识为等待垃圾回收。

### 持久化的缺点

消息持久化的优点显而易见，但缺点也很明显，那就是性能，因为要写入硬盘要比写入内存性能较低很多，从而降低了服务器的吞吐量，尽管使用 SSD 硬盘可以使事情得到缓解，但它仍然吸干了 Rabbit 的性能，当消息成千上万条要写入磁盘的时候，性能是很低的。

所以使用者要根据自己的情况，选择适合自己的方式。

## 虚拟主机 Virtual Host

每个 Rabbit 都能创建很多 vhost，我们称之为虚拟主机，每个虚拟主机其实都是 mini 版的 RabbitMQ，拥有自己的队列，交换器和绑定，拥有自己的权限机制，这样能使数据隔离，区分队列

### vhost 特性

RabbitMQ 默认的 vhost 是“/”开箱即用；

多个 vhost 是隔离的，多个 vhost 无法通讯，并且不用担心命名冲突（队列和交换器绑定），实现了多层分离；

创建用户的时候必须指定 vhost；

### vhost 操作

可以通过 rabbitmqctl 工具命令创建：

```sh
rabbitmqctl add_vhost[vhost_name]
```

删除 vhost：

```sh
rabbitmqctl delete_vhost[vhost_name]

```

查看所有的 vhost：

```
rabbitmqctl list_vhosts
```

## 代码实现

### Rabbit 的连接

### 使用最简单的方式发布和消费消息

**参考**:

- https://www.cnblogs.com/vipstone/p/9275256.html
