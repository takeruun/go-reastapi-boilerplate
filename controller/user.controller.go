package controller

import (
	"app/usecase"
	"net/http"
)

type UserController interface {
	Index(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userU usecase.UserUsecase
}

func NewUserController(userU usecase.UserUsecase) UserController {
	return &userController{
		userU: userU,
	}
}

func (userCon *userController) Index(w http.ResponseWriter, r *http.Request) {
}

func (userCon *userController) Show(w http.ResponseWriter, r *http.Request) {
}

func (userCon *userController) Edit(w http.ResponseWriter, r *http.Request) {
}

func (userCon *userController) Create(w http.ResponseWriter, r *http.Request) {
}

func (userCon *userController) Delete(w http.ResponseWriter, r *http.Request) {
}
