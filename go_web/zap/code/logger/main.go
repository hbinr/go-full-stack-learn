package main

import (
	"net/http"

	"go.uber.org/zap"
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

// InitLogger 初始化 zap.Looer和sugarLogger
func InitLogger() {
	logger, _ = zap.NewProduction()
	sugarLogger = logger.Sugar()
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
