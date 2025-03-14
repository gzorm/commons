package handler

import (
	"context"
	"github.com/gzorm/commons/rest/enum"
	"net/http"
)

// I18nHandler returns a middleware that inject accept-language into context
func I18nHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add lang to context
		if lang := r.Header.Get("Accept-Language"); lang == "" {
			ctx = context.WithValue(ctx, enum.I18N_CTX_KEY, "zh-CN")
		} else {
			ctx = context.WithValue(ctx, enum.I18N_CTX_KEY, lang)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
