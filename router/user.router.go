package router

import "net/http"

type UserRouter interface {
	HandleUserRequest(w http.ResponseWriter, r *http.Request)
}

type userRouter struct{}

func NewUserRouter() UserRouter {
	return &userRouter{}
}

func (ur *userRouter) HandleUserRequest(w http.ResponseWriter, r *http.Request) {}
