package router

import (
	"app/controller"
	"net/http"
)

type UserRouter interface {
	HandleUserRequest(w http.ResponseWriter, r *http.Request)
}

type userRouter struct {
	userC controller.UserController
}

func NewUserRouter(userC controller.UserController) UserRouter {
	return &userRouter{
		userC: userC,
	}
}

func (ur *userRouter) HandleUserRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ur.userC.Index(w, r)
	case "POST":
		ur.userC.Create(w, r)
	case "PUT":
		ur.userC.Edit(w, r)
	case "DELETE":
		ur.userC.Delete(w, r)
	default:
		w.WriteHeader(405)
	}
}
