#  Dapr 基础命令
## dapr 环境初始化
```bash
dapr init
```
初始化 Dapr 包括获取 Dapr sidecar 二进制文件并将其安装到本地.

此外，默认初始化过程还创建了一个开发环境，帮助简化 Dapr 的应用开发。 这包括下列步骤：

- 运行一个用于状态存储和消息代理的 Redis 容器实例
- 运行一个用于提供可观察性的 Zipkin 容器实例
- 创建具有上述组件定义的 默认组件文件夹，在 Linux/MacOS 中 Dapr 使用默认组件和文件的路径是 `$HOME.dap`r。
  - 执行 `ls $HOME/.dapr`, 可以看到：
  - > bin  components  config.yaml
- 运行用于本地演员支持的 Dapr placement 服务容器实例
## 启动应用
```bash
dapr run --app-id myapp --dapr-http-port 3500
```
其中 `myapp` 是自己 app 的名称， `--dapr-http-port` 可以通过配置文件来定义，也可以在命令行中指定参数

## 停止应用
```bash
dapr stop --app-id myapp
```
其中 `myapp` 是自己 app 的名称
