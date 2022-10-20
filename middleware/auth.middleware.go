package middleware

import (
	"net/http"

	"github.com/wader/gormstore/v2"
)

func AuthMiddleware(next http.Handler, store *gormstore.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "_goreset_session")
		_, err := r.Cookie("_goreset_session")

		if err != nil || session.Values["userId"] == nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message":"ログインしてください。"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}
