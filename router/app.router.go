package router

import (
	"net/http"
)

type AppRouter interface {
	HandleAppRequest(w http.ResponseWriter, r *http.Request)
}

type appRouter struct{}

func NewAppRouter() AppRouter {
	return &appRouter{}
}

func (ar *appRouter) HandleAppRequest(w http.ResponseWriter, r *http.Request) {

}
