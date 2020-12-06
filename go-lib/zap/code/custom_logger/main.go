package main

import (
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	// Sync调用基础Core的Sync方法，刷新所有缓冲的日志词条
	// 应用程序应注意退出前调用Sync
	defer logger.Sync()
	testHTTPGet("www.baidu.com")        // error
	testHTTPGet("http://www.baidu.com") // success

	testHTTPGet2("www.baidu.com")
	testHTTPGet2("http://www.baidu.com")
}

// InitLogger 定制 zap.Logger和sugarLogger
func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	// 使用debug级别记录日志
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// zap.AddCaller() 将调用函数信息记录到日志中的功能
	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

// getEncoder 设置日志编码
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 将time.Time序列化为ISO8601格式的字符串毫秒精度。
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 将Level序列化为全大写字符串。例如，InfoLevel被序列化为“ INFO”。
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
	// 使用JSON编码，并使用开发环境的默认配置
	// return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// 使用普通Encoder编码，并使用开发环境的默认配置
	// return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}

// getLogWriter 设置日志写入到的文件路径
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./httprouter_test.log")
	return zapcore.AddSync(file)
}

// testHttpGet 测试logger日志
func testHTTPGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		logger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

// testHttpGet 测试sugarLogger日志
func testHTTPGet2(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}
