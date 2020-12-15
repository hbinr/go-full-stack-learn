

版本：
/common/logger/logger.go 131 行左右
使用zap，并且是`Sugar Logger`

```go 
zapLogger, _ := zapLoggerConfig.Build(zap.AddCallerSkip(1))
logger = &DubboLogger{Logger: zapLogger.Sugar(), dynamicLevel: zapLoggerConfig.Level}

// set getty log
getty.SetLogger(logger)
```