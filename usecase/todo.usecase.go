package usecase

import (
	"app/database"
	"app/entity"
)

type TodoUsecase interface {
	FindAll() (todos *[]entity.Todo, err error)
}

type todoUsecase struct {
	userRepo database.TodoRepository
}

func NewTodoUsecase(userRepo database.TodoRepository) TodoUsecase {
	return &todoUsecase{
		userRepo: userRepo,
	}
}

func (uu *todoUsecase) FindAll() (todos *[]entity.Todo, err error) {
	return uu.userRepo.FindAll()
}
