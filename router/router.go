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
	http.Handle("/users/", middleware.CorsMiddleware(middleware.SetHttpContextMiddleware(middleware.WriteHeaderMiddleware(http.HandlerFunc(r.userRoute.HandleUserRequest)))))
	http.Handle("/todos/", middleware.CorsMiddleware(middleware.SetHttpContextMiddleware(middleware.WriteHeaderMiddleware(http.HandlerFunc(r.todoRoute.HandleTodoRequest)))))
	http.Handle("/auth/sign_in/", middleware.CorsMiddleware(middleware.SetHttpContextMiddleware(middleware.WriteHeaderMiddleware(http.HandlerFunc(r.authRoute.SignInRequest)))))
	http.Handle("/auth/sign_up/", middleware.CorsMiddleware(middleware.SetHttpContextMiddleware(middleware.WriteHeaderMiddleware(http.HandlerFunc(r.authRoute.SignUpRequest)))))
}
