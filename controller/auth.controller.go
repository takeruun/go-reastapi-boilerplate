package controller

import (
	"app/usecase"
	"net/http"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
}

type authController struct {
	authU usecase.AuthUsecase
}

func NewAuthController(authU usecase.AuthUsecase) AuthController {
	return &authController{
		authU: authU,
	}
}

func (authCon *authController) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authCon.authU.SignIn(r.Context())
}

func (authCon *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
