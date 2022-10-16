package router

import (
	"app/middleware"
	"net/http"
)

type Router interface {
	SetRouting()
}

type router struct {
	appRoute  AppRouter
	userRoute UserRouter
	authRoute AuthRouter
	todoRoute TodoRouter
}

func NewRouter(appRoute AppRouter, userRoute UserRouter, authRoute AuthRouter, todoRoute TodoRouter) Router {
	return &router{
		appRoute:  appRoute,
		userRoute: userRoute,
		authRoute: authRoute,
		todoRoute: todoRoute,
	}
}

func (r *router) SetRouting() {

	http.HandleFunc("/", r.appRoute.HandleAppRequest)
	http.Handle("/users/", middleware.SetHttpContextMiddleware(http.HandlerFunc(r.userRoute.HandleUserRequest)))
	http.Handle("/todos/", middleware.SetHttpContextMiddleware(http.HandlerFunc(r.todoRoute.HandleTodoRequest)))
	http.Handle("/sign_in/", middleware.SetHttpContextMiddleware(http.HandlerFunc(r.authRoute.SignInRequest)))
	http.Handle("/sign_up/", middleware.SetHttpContextMiddleware(http.HandlerFunc(r.authRoute.SignUpRequest)))
}
