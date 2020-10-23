# 定制 zap logger

[详细代码见](../code/custom_logger/main.go)

## 日志级别

先了解下日志级别：

- Info: 记录一些关键的状态点，数据中转点
- Debug：通常在开发阶段使用，debug 时记录项目开发时遇到的问题，便于定位、解决问题、方便开发
- Waring：记录一些即将有“危险”的信息，这些逻辑在正常情况下不会执行，即使意外执行了也不会影响项目的正常运行，但是需要格外注意，因为这意味着项目的逻辑执行到“危险区域”了，需要注意、排查隐患。
- Error：记录错误日志，记录项目运行过程发生的错误信息

## 将日志写入文件而不是终端

我们要做的第一个更改是把日志写入文件，而不是打印到应用程序控制台。

- 我们将使用 `zap.New(…)`方法来手动传递所有配置，而不是使用像 `zap.NewProduction()`这样的预置方法来创建 logger。

```go
func New(core zapcore.Core, options ...Option) *Logger
```

`zapcore.Core` 需要三个配置——`Encoder`，`WriteSyncer`，`LogLevel`。

### 1.Encoder 编码器(如何写入日志)

我们将使用开箱即用的 `NewJSONEncoder()`，并使用预先设置的`ProductionEncoderConfig()`。

```go
zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
```

### 2.WriterSyncer ：指定日志将写到哪里去

我们使用 `zapcore.AddSync()`函数并且将打开的文件句柄传进去。

```go
file, _ := os.Create("./test.log")
writeSyncer := zapcore.AddSync(file)
```

### 3.Log Level：哪种级别的日志将被写入

我们并重写 `InitLogger()`方法，并新增了 `getEncoder()`——设置编码，`getLogWriter()`——指定日志将写到哪里去，两个方法

```go
func InitLogger() {
	writeSyncer := getLogWriter()
    encoder := getEncoder()
    // 使用debug级别记录日志
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}
```

当使用这些修改过的 logger 配置调用详细代码中的 main()函数时，以下输出将打印在文件——test.log 中。

```sh
{"level":"error","ts":1602124670.2491782,"msg":"Error fetching url..","url":"www.baidu.com","error":"Get \"www.baidu.com\": unsupported protocol scheme \"\""}
{"level":"info","ts":1602124670.2766364,"msg":"Success..","statusCode":"200 OK","url":"http://www.baidu.com"}
{"level":"debug","ts":1602124670.2767293,"msg":"Trying to hit GET request for www.baidu.com"}
{"level":"error","ts":1602124670.276776,"msg":"Error fetching URL www.baidu.com : Error = Get \"www.baidu.com\": unsupported protocol scheme \"\""}
{"level":"debug","ts":1602124670.2767882,"msg":"Trying to hit GET request for http://www.baidu.com"}
{"level":"info","ts":1602124670.3166301,"msg":"Success! statusCode = 200 OK for URL http://www.baidu.com"}
```

## 将 JSON Encoder 更改为普通的 Log Encoder

现在，我们希望将编码器从 JSON Encoder 更改为普通 Encoder。为此，我们需要将 `NewJSONEncoder()`更改为 `NewConsoleEncoder()`。

```go
return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
```

当使用这些修改过的 logger 配置调用详细代码中的 main()函数时，以下输出将打印在文件——test.log 中。

```sh
1.6021245580757637e+09	error	Error fetching url..	{"url": "www.baidu.com", "error": "Get \"www.baidu.com\": unsupported protocol scheme \"\""}
1.6021245581827517e+09	info	Success..	{"statusCode": "200 OK", "url": "http://www.baidu.com"}
1.60212455818284e+09	debug	Trying to hit GET request for www.baidu.com
1.6021245581828763e+09	error	Error fetching URL www.baidu.com : Error = Get "www.baidu.com": unsupported protocol scheme ""
1.6021245581828847e+09	debug	Trying to hit GET request for http://www.baidu.com
1.6021245582454193e+09	info	Success! statusCode = 200 OK for URL http://www.baidu.com
```

## 更改时间编码并添加调用者详细信息

鉴于我们对配置所做的更改，有下面两个问题：

- 时间是以非人类可读的方式展示，例如 1.572161051846623e+09
- 调用方函数的详细信息没有显示在日志中

我们要做的第一件事是覆盖默认的 `ProductionConfig()`，并进行以下更改:

- 修改时间编码器
- 在日志文件中使用大写字母记录日志级别

```go
func getEncoder() zapcore.Encoder {
    encoderConfig := zap.NewProductionEncoderConfig()
    // 将time.Time序列化为ISO8601格式的字符串毫秒精度。
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    // 将Level序列化为全大写字符串。例如，InfoLevel被序列化为“ INFO”。
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
```

接下来，我们将修改 zap logger 代码，添加将调用函数信息记录到日志中的功能。为此，我们将在 `zap.New(..)`函数中添加一个 `Option`

```go
logger := zap.New(core, zap.AddCaller())
```

当使用这些修改过的 logger 配置调用详细代码中的 main()函数时，以下输出将打印在文件——test.log 中。

```sh
2020-10-08T10:42:14.381+0800	ERROR	custom_logger/main.go:62	Error fetching url..	{"url": "www.baidu.com", "error": "Get \"www.baidu.com\": unsupported protocol scheme \"\""}
2020-10-08T10:42:14.473+0800	INFO	custom_logger/main.go:67	Success..	{"statusCode": "200 OK", "url": "http://www.baidu.com"}
2020-10-08T10:42:14.473+0800	DEBUG	custom_logger/main.go:76	Trying to hit GET request for www.baidu.com
2020-10-08T10:42:14.473+0800	ERROR	custom_logger/main.go:79	Error fetching URL www.baidu.com : Error = Get "www.baidu.com": unsupported protocol scheme ""
2020-10-08T10:42:14.473+0800	DEBUG	custom_logger/main.go:76	Trying to hit GET request for http://www.baidu.com
2020-10-08T10:42:14.517+0800	INFO	custom_logger/main.go:81	Success! statusCode = 200 OK for URL http://www.baidu.com
```
