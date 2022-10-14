package router

import "net/http"

type AuthRouter interface {
	SignInRoute(w http.ResponseWriter, r *http.Request)
	SignUpRoute(w http.ResponseWriter, r *http.Request)
}

type authRouter struct{}

func NewAuthRouter() AuthRouter {
	return &authRouter{}
}

func (ar *authRouter) SignInRoute(w http.ResponseWriter, r *http.Request) {}

func (ar *authRouter) SignUpRoute(w http.ResponseWriter, r *http.Request) {}
