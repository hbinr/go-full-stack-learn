package test

import (
	"errors"
	"net/http"

	"dubbo.apache.org/dubbo-go/v3/common"

	"dubbo.apache.org/dubbo-go/v3/protocol/rest/config"
	"dubbo.apache.org/dubbo-go/v3/protocol/rest/server"
	"dubbo.apache.org/dubbo-go/v3/protocol/rest/server/server_impl"
)

// TestGoRestfulServer
func TestGoRestfulServer() {
	grs := server_impl.NewGoRestfulServer()
	url, _ := common.NewURL("http://127.0.0.1:43121")
	grs.Start(url)
	rmc := &config.RestMethodConfig{
		Produces:   "*/*",
		Consumes:   "*/*",
		MethodType: "GET",
		Path:       "/hello",
	}
	grs.Deploy(rmc, Hello)
	grs.UnDeploy(rmc)
	grs.Destroy()
}

func Hello(req server.RestServerRequest, res server.RestServerResponse) {
	msg := req.QueryParameter("msg")

	if msg == "" {
		res.WriteError(http.StatusInternalServerError, errors.New("invalid param"))
	}
	var entity = make(map[string]interface{})
	entity["code"] = 200
	entity["msg"] = msg
	_ = res.WriteEntity(entity)
}
