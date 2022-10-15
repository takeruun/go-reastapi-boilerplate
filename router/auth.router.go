package router

import (
	"app/controller"
	"net/http"
)

type AuthRouter interface {
	SignInRoute(w http.ResponseWriter, r *http.Request)
	SignUpRoute(w http.ResponseWriter, r *http.Request)
}

type authRouter struct {
	authC controller.AuthController
}

func NewAuthRouter(authC controller.AuthController) AuthRouter {
	return &authRouter{
		authC: authC,
	}
}

func (ar *authRouter) SignInRoute(w http.ResponseWriter, r *http.Request) {
	ar.authC.SignIn(w, r)
}

func (ar *authRouter) SignUpRoute(w http.ResponseWriter, r *http.Request) {
	ar.authC.SignUp(w, r)
}
