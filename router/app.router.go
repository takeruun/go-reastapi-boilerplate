package router

import (
	"app/controller"
	"net/http"
)

type AppRouter interface {
	HandleAppRequest(w http.ResponseWriter, r *http.Request)
}

type appRouter struct {
	appCon controller.AppController
}

func NewAppRouter(appCon controller.AppController) AppRouter {
	return &appRouter{
		appCon: appCon,
	}
}

func (ar *appRouter) HandleAppRequest(w http.ResponseWriter, r *http.Request) {
	ar.appCon.Index(w, r)
}
