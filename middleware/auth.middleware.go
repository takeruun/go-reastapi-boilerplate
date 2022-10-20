package middleware

import (
	"app/config"
	"net/http"

	"github.com/wader/gormstore/v2"
)

func AuthMiddleware(next http.Handler, store *gormstore.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, config.SESSION_KEY_NAME)
		_, err := r.Cookie(config.SESSION_KEY_NAME)

		if err != nil || session.Values["userId"] == nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message":"ログインしてください。"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}
