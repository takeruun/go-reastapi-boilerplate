package database

import (
	"app/config"
	"app/entity"
)

type TodoRepository interface {
	FindAll() (todos *[]entity.Todo, err error)
}

type todoRepository struct {
	DB *config.DB
}

func NewTodoRepository(DB *config.DB) TodoRepository {
	return &todoRepository{DB: DB}
}

func (todoRep *todoRepository) FindAll() (todos *[]entity.Todo, err error) {
	err = todoRep.DB.Model(&entity.Todo{}).
		Find(&todos).
		Error

	if err != nil {
		return nil, err
	}

	return
}
