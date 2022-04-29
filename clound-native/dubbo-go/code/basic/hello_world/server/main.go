package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hb.study/clound-native/dubbo-go/code/basic/hello_world/server/filter"
	"hb.study/clound-native/dubbo-go/code/basic/hello_world/server/provider"

	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/emicklei/go-restful/v3"

	_ "dubbo.apache.org/dubbo-go/v3/common/proxy/proxy_factory"
	"dubbo.apache.org/dubbo-go/v3/config"

	_ "dubbo.apache.org/dubbo-go/v3/protocol/rest"
	"dubbo.apache.org/dubbo-go/v3/protocol/rest/server/server_impl"

	_ "dubbo.apache.org/dubbo-go/v3/registry/protocol"

	_ "dubbo.apache.org/dubbo-go/v3/filter/filter_impl"

	_ "dubbo.apache.org/dubbo-go/v3/cluster/cluster_impl"

	_ "dubbo.apache.org/dubbo-go/v3/cluster/loadbalance"
	_ "dubbo.apache.org/dubbo-go/v3/config_center/nacos"
	_ "dubbo.apache.org/dubbo-go/v3/registry/nacos"
)

/**
先设置环境变量：
	export CONF_PROVIDER_FILE_PATH="./conf/server.yml"
	export APP_LOG_CONF_FILE="./conf/log.yml"
*/
func main() {

	config.SetProviderService(new(provider.UserProvider))
	extension.SetFilter("ErrResponseFilter", filter.GetErrResponseFilter)
	// 注意：import go-restful 版本为v3
	server_impl.AddGoRestfulServerFilter(func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		//gxlog.CInfo(request.SelectedRoutePath())
		//gxlog.CInfo("request %v", request)
		//gxlog.CInfo("response %v", response)
		chain.ProcessFilter(request, response)
	})

	config.Load()

	initSignal()
}

var (
	survivalTimeout = int(3e9)
)

func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
		// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("test exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider test exit now...")
			return
		}
	}
}
