package middleware

import (
	"app/service"
	"context"
	"net/http"
)

func SetHttpContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpContext := service.HTTP{
			W: &w,
			R: r,
		}
		httpKeyContext := service.HTTPKey("http")

		ctx := context.WithValue(r.Context(), httpKeyContext, httpContext)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
