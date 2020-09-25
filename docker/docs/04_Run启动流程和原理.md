## `docker run` 命令执行后的流程图

![](../img/docker%20run原理图.jpg)

## 底层原理

### docker 是怎么工作的?

- Docker 是一个 Client-Server 结构的系统，Docker 的守护进程运行在主机上。通过 Socket 从客户端访问！
- Docker-Server 接收到 Docker-Client 的指令，就会执行这个命令！

![](https://img-blog.csdnimg.cn/20190618180935284.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzM3Mjk0ODM4,size_16,color_FFFFFF,t_70../img/docker底层原理图.jpg)

### docker 为什么比虚拟机快？

1. docker 有着比虚拟机更少的抽象层。由亍 docker 不需要 Hypervisor 实现硬件资源虚拟化,运行在 docker 容器上的程序直接使用的都是实际物理机的硬件资源。因此在 CPU、内存利用率上 docker 将会在效率上有明显优势。

2. docker 利用的是宿主机的内核,而不需要 Guest OS。因此,当新建一个容器时,docker 不需要和虚拟机一样重新加载一个操作系统内核。仍而避免引寻、加载操作系统内核返个比较费时费资源的过程,当新建一个虚拟机时,虚拟机软件需要加载 Guest OS,返个新建过程是分钟级别的。而 docker 由于直接利用宿主机的操作系统,则省略了返个过程,因此新建一个 docker 容器只需要几秒钟。

![](https://img-blog.csdnimg.cn/20190618180935324.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzM3Mjk0ODM4,size_16,color_FFFFFF,t_70)

**docker 与虚拟机对比：**

![](https://img-blog.csdnimg.cn/2019061818093616.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L20wXzM3Mjk0ODM4,size_16,color_FFFFFF,t_70)

参考：

[B 站视频](https://www.bilibili.com/video/BV1og4y1q7M4?p=8)
