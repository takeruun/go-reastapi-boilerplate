package controller

import (
	"app/controller/dto"
	"app/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthController interface {
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var signInParams dto.AuthSignInRequestDto
	json.Unmarshal(body, &signInParams)

	authCon.authU.SignIn(r.Context(), &signInParams)
}

func (authCon *authController) SignUp(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var signUpParams dto.AuthSignUpRequestDto
	json.Unmarshal(body, &signUpParams)

	authCon.authU.SignUp(r.Context(), &signUpParams)
}

func (authCon *authController) Show(w http.ResponseWriter, r *http.Request) {
	user, err := authCon.authU.Show(r.Context())

	var data []byte
	data, err = json.MarshalIndent(&user, "", "\t\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.Write(data)
}

func (authCon *authController) Edit(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var userUpdateParams dto.AuthUserUpdateRequestDto
	json.Unmarshal(body, &userUpdateParams)

	user, err := authCon.authU.Edit(r.Context(), &userUpdateParams)

	var data []byte
	data, err = json.MarshalIndent(&user, "", "\t\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.Write(data)
}

func (authCon *authController) Delete(w http.ResponseWriter, r *http.Request) {
	err := authCon.authU.Delete(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
		return
	}

	w.Write([]byte(`{"message": "ok"}`))
}
