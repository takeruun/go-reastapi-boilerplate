package controller

import (
	"net/http"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
}

type authController struct{}

func NewAuthController() AuthController {
	return &authController{}
}

func (authCon *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (authCon *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
