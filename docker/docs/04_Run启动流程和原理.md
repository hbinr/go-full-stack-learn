## `docker run` 命令执行后的流程图

![](../img/docker%20run原理图.jpg)

## 底层原理

### docker 是怎么工作的?

- Docker 是一个 Client-Server 结构的系统，Docker 的守护进程运行在主机上。通过 Socket 从客户端访问！
- Docker-Server 接收到 Docker-Client 的指令，就会执行这个命令！

![](../img/docker底层原理图.jpg)
