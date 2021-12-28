`Dockerfile` 文件示例

```s
FROM golang:1.15 AS builder

COPY . /src
#  WORKDIR
# 指定工作目录。用 WORKDIR 指定的工作目录，会在构建镜像的每一层中都存在（WORKDIR 指定的工作目录，必须是提前创建好的）
# docker build 构建镜像过程中的，每一个 RUN 命令都是新建的一层。只有通过 WORKDIR 创建的目录才会一直存在
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

# 定制的镜像都是基于 FROM 的镜像
FROM debian:stable-slim

# 用于执行后面跟着的命令行命令
# 注意：Dockerfile 的指令每执行一次都会在 docker 上新建一层。所以过多无意义的层，会造成镜像膨胀过大。所以用 && 连接命令，这样执行后，只会创建 1 层镜像。
RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates  netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y \
    && apt-get autoclean -y

# COPY 复制指令，从上下文目录中复制文件或者目录到容器里指定路径。
COPY --from=builder /src/bin /app

WORKDIR /app

# EXPOSE 仅仅只是声明端口。
# 作用：
# - 帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射。
# - 在运行时使用随机端口映射时，也就是 docker run -P 时，会自动随机映射 EXPOSE 的端口。
EXPOSE 8000
EXPOSE 9000

# VOLUME 定义匿名数据卷。在启动容器时忘记挂载数据卷，会自动挂载到匿名卷。
# 作用：
# - 避免重要的数据，因容器重启而丢失，这是非常致命的。
# - 避免容器不断变大。
VOLUME /data/conf

# CMD 类似于 RUN 指令，用于运行程序，但二者运行的时间点不同:
# - CMD 在docker run 时运行。
# - RUN 是在 docker build。
# 作用：为启动的容器指定默认要运行的程序，程序运行结束，容器也就结束。CMD 指令指定的程序可被 docker run 命令行参数中指定要运行的程序所覆盖。

# 注意：如果 Dockerfile 中如果存在多个 CMD 指令，仅最后一个生效。
CMD ["./server", "-conf", "/data/conf"]

```

参考：
- https://www.runoob.com/docker/docker-dockerfile.html