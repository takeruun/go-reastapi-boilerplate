package router

import "net/http"

type TodoRouter interface {
	HandleTodoRequest(w http.ResponseWriter, r *http.Request)
}

type todoRouter struct{}

func NewTodoRouter() TodoRouter {
	return &todoRouter{}
}

func (ur *todoRouter) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {}
