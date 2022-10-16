package router

import (
	"app/controller"
	"net/http"
	"strconv"
	"strings"
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
		id := strings.TrimPrefix(r.URL.Path, "/todos/") // URLを切り取る
		if id != "" {
			todoId, _ := strconv.Atoi(id)
			ur.todoC.Show(w, r, todoId)
		} else {
			ur.todoC.Index(w, r)
		}
	case "POST":
		ur.todoC.Create(w, r)
	case "PUT":
		id := strings.TrimPrefix(r.URL.Path, "/todos/") // URLを切り取る
		todoId, _ := strconv.Atoi(id)
		ur.todoC.Edit(w, r, todoId)
	case "DELETE":
		id := strings.TrimPrefix(r.URL.Path, "/todos/") // URLを切り取る
		todoId, _ := strconv.Atoi(id)
		ur.todoC.Delete(w, r, todoId)
	default:
		w.WriteHeader(405)
	}
}
