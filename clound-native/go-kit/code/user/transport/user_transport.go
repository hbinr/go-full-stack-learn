package transport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"hb.study/clound-native/go-kit/code/user/endpoint"
)

// DecodeUserRequest 解码请求参数
func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	// http:localhost:xxx?id="xx"
	if r.URL.Query().Get("id") != "" {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
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
