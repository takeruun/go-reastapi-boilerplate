package controller

import (
	"app/controller/dto"
	"app/usecase"
	"encoding/json"
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

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var signInParams dto.AuthSignInRequestDto
	json.Unmarshal(body, &signInParams)

	authCon.authU.SignIn(r.Context(), &signInParams)
}

func (authCon *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
