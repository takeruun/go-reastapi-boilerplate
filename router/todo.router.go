package router

import (
	"app/controller"
	"net/http"
)

type TodoRouter interface {
	HandleTodoRequest(w http.ResponseWriter, r *http.Request)
}

type todoRouter struct {
	todoC controller.TodoController
}

func NewTodoRouter(todoC controller.TodoController) TodoRouter {
	return &todoRouter{
		todoC: todoC,
	}
}

func (ur *todoRouter) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ur.todoC.Index(w, r)
	case "POST":
		ur.todoC.Create(w, r)
	case "PUT":
		ur.todoC.Edit(w, r)
	case "DELETE":
		ur.todoC.Delete(w, r)
	default:
		w.WriteHeader(405)
	}
}
