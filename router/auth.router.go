package router

import (
	"app/controller"
	"net/http"
)

type AuthRouter interface {
	SignInRequest(w http.ResponseWriter, r *http.Request)
	SignUpRequest(w http.ResponseWriter, r *http.Request)
	UserRequest(w http.ResponseWriter, r *http.Request)
}

type authRouter struct {
	authC controller.AuthController
}

func NewAuthRouter(authC controller.AuthController) AuthRouter {
	return &authRouter{
		authC: authC,
	}
}

func (ar *authRouter) SignInRequest(w http.ResponseWriter, r *http.Request) {
	ar.authC.SignIn(w, r)
}

func (ar *authRouter) SignUpRequest(w http.ResponseWriter, r *http.Request) {
	ar.authC.SignUp(w, r)
}

func (ar *authRouter) UserRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ar.authC.Show(w, r)
	case "PUT":
		ar.authC.Edit(w, r)
	case "DELETE":
		ar.authC.Delete(w, r)
	}
}
