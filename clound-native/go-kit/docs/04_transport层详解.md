

## 编解码原则
针对go中的`struct`而言，现在网络中大都使用json格式来处理数据，因此返回给前端的数据主要以Json格式为主
### 解码request
将http发送过来的**请求**byte内容**解码**为go程序中能识别的**结构体**，`go-kit`框架中，服务端大部分都会解码请求，转成`struct`，示例如下：
#### 代码示例
```go
type SumRequest struct {
	A, B int
}


// decodeHTTPSumRequest 解码来自HTTP请求主体的JSON编码的请求。主要用于server端
func decodeHTTPSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req SumRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
```
### 编码request
将JSON结构编码到请求正文中。 主要用于client端，该方法可以封装为一个公共方法
#### 代码示例
```go
// encodeHTTPGenericRequest 将JSON结构编码到请求正文中。 主要用于client端
func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
```
** NopCloser 函数**
有时候我们需要传递一个`i`o.ReadCloser`的实例，而我们现在有一个`io.Reader`的实例，比如：strings.Reader，这个时候`NopCloser`就派上用场了。它包装一个`io.Reader`，返回一个`io.ReadCloser`，而相应的Close方法啥也不做，只是返回nil。

比如，在标准库`net/http`包中的`NewRequest`，接收一个`io.Reader`的body，而实际上，Request的Body的类型是`io.ReadCloser`，因此，代码内部进行了判断，如果传递的`io.Reader`也实现了`io.ReadCloser`接口，则转换，否则通过`ioutil.NopCloser`包装转换一下。
### 解码response
encodeResponse 解码HTTP响应主体（数据格式为json），所以可以通过json库将HTTP响应主体转为go中的结构体，供其它层调用

如果响应的状态代码不是200，我们会将其解释为错误，并尝试从响应正文中解码特定的错误消息。 主要用于客户端。
#### 代码示例
```go
type SumResponse struct {
	V   int   `json:"v"`
	Err error `json:"-"` // should be intercepted by Failed/errorEncoder
}

func decodeHTTPSumResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp SumResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
```

### 编码response
#### 代码示例
```go
type SumResponse struct {
	V   int   `json:"v"`
	Err error `json:"-"` // should be intercepted by Failed/errorEncoder
}
// encodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.

// encodeHTTPGenericResponse 它将响应编码为JSON编码到响应编写器。 主要用于server端
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(SumResponse); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
```