package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie("_goreset_session"); err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message":"ログインしてください。"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}
