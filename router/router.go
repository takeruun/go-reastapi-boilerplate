package router

import (
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
	http.HandleFunc("/users/", r.userRoute.HandleUserRequest)
	http.HandleFunc("/todos/", r.todoRoute.HandleTodoRequest)
	http.HandleFunc("/sign_in/", r.authRoute.SignInRequest)
	http.HandleFunc("/sign_up/", r.authRoute.SignUpRequest)
}
