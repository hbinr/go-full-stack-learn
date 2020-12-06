package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"hb.study/clound-native/go-kit/code/app/endpoint"
)

// DecodeUserRequest 解码请求参数
func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	// http://localhost:8080/get/1
	params := httprouter.ParamsFromContext(r.Context())
	if idStr := params.ByName("id"); idStr != "" {
		id, _ := strconv.Atoi(idStr)
		return endpoint.UserRequest{
			UserID: id,
		}, nil
	}

	return nil, errors.New("请求参数错误")
}

// EncodeUserResponse 将响应参数进行编码
func EncodeUserResponse(c context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}
