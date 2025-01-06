package httpx

import (
	"bytes"
	"net/http"
)

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int           // 记录状态码
	Body       *bytes.Buffer // 记录响应体
}

func (rec *ResponseRecorder) WriteHeader(code int) {
	rec.StatusCode = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *ResponseRecorder) Write(data []byte) (int, error) {
	if rec.Body == nil {
		rec.Body = &bytes.Buffer{}
	}
	rec.Body.Write(data) // 记录响应体内容
	return rec.ResponseWriter.Write(data)
}
