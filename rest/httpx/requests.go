package httpx

import (
	"io"
	"net/http"
	"strings"

	"github.com/gzorm/commons/core/errorx"
	"github.com/gzorm/commons/core/mapping"
	"github.com/gzorm/commons/rest/internal/encoding"
	"github.com/gzorm/commons/rest/internal/header"
	"github.com/gzorm/commons/rest/pathvar"
)

// @enhance
const (
	formKey           = "form"
	pathKey           = "path"
	maxMemory         = 32 << 20 // 32MB
	maxBodyLen        = 8 << 20  // 8MB
	separator         = ";"
	tokensInAttribute = 2
)

var (
	formUnmarshaler = mapping.NewUnmarshaler(formKey, mapping.WithStringValues(), mapping.WithOpaqueKeys())
	pathUnmarshaler = mapping.NewUnmarshaler(pathKey, mapping.WithStringValues(), mapping.WithOpaqueKeys())
	xValidator      = NewValidator()
)

// Parse parses the request.
func Parse(r *http.Request, v any, isValidate bool) error {
	if err := ParseJsonBody(r, v); err != nil {
		return errorx.NewCodeInvalidArgumentError(err.Error())
	}

	if err := ParsePath(r, v); err != nil {
		return errorx.NewCodeInvalidArgumentError(err.Error())
	}

	if err := ParseForm(r, v); err != nil {
		return errorx.NewCodeInvalidArgumentError(err.Error())
	}

	if err := ParseHeaders(r, v); err != nil {
		return errorx.NewCodeInvalidArgumentError(err.Error())
	}

	if isValidate {
		if errMsg := xValidator.Validate(v, r.Header.Get("Accept-Language")); errMsg != "" {
			return errorx.NewCodeError(xValidator.ErrorCode, errMsg)
		}
	}
	return nil
}

// ParseHeaders parses the headers request.
func ParseHeaders(r *http.Request, v any) error {
	return encoding.ParseHeaders(r.Header, v)
}

// ParseForm parses the form request.
func ParseForm(r *http.Request, v any) error {
	params, err := GetFormValues(r)
	if err != nil {
		return err
	}

	return formUnmarshaler.Unmarshal(params, v)
}

// ParseHeader parses the request header and returns a map.
func ParseHeader(headerValue string) map[string]string {
	ret := make(map[string]string)
	fields := strings.Split(headerValue, separator)

	for _, field := range fields {
		field = strings.TrimSpace(field)
		if len(field) == 0 {
			continue
		}

		kv := strings.SplitN(field, "=", tokensInAttribute)
		if len(kv) != tokensInAttribute {
			continue
		}

		ret[kv[0]] = kv[1]
	}

	return ret
}

// ParseJsonBody parses the post request which contains json in body.
func ParseJsonBody(r *http.Request, v any) error {
	if withJsonBody(r) {
		reader := io.LimitReader(r.Body, maxBodyLen)
		return mapping.UnmarshalJsonReader(reader, v)
	}

	return mapping.UnmarshalJsonMap(nil, v)
}

// ParsePath parses the symbols reside in url path.
// Like http://localhost/bag/:name
func ParsePath(r *http.Request, v any) error {
	vars := pathvar.Vars(r)
	m := make(map[string]any, len(vars))
	for k, v := range vars {
		m[k] = v
	}

	return pathUnmarshaler.Unmarshal(m, v)
}

func withJsonBody(r *http.Request) bool {
	return r.ContentLength > 0 && strings.Contains(r.Header.Get(header.ContentType), header.ApplicationJson)
}
