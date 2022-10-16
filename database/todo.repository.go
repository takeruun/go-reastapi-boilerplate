package database

import (
	"app/config"
	"app/entity"
)

type TodoRepository interface {
	FindAll(userId uint64) (todos []*entity.Todo, err error)
	Create(entity *entity.Todo) (todo *entity.Todo, err error)
	Get(todoId int) (todo *entity.Todo, err error)
}

type todoRepository struct {
	DB *config.DB
}

func NewTodoRepository(DB *config.DB) TodoRepository {
	return &todoRepository{DB: DB}
}

func (todoRep *todoRepository) FindAll(userId uint64) (todos []*entity.Todo, err error) {
	err = todoRep.DB.Model(&entity.Todo{}).
		Where("user_id = ?", userId).
		Find(&todos).
		Error

	if err != nil {
		return nil, err
	}

	return
}

func (todoRep *todoRepository) Create(entity *entity.Todo) (todo *entity.Todo, err error) {
	err = todoRep.DB.Create(&entity).Error
	if err != nil {
		return nil, err
	}

	err = todoRep.DB.Find(&todo, &entity.ID).Error
	if err != nil {
		return nil, err
	}

	return
}

func (todoRep *todoRepository) Get(todoId int) (todo *entity.Todo, err error) {
	err = todoRep.DB.Find(&todo, todoId).Error
	if err != nil {
		return nil, err
	}

	return
}
