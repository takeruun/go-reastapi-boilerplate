package database

import (
	"app/config"
	"app/entity"
)

type TodoRepository interface {
	FindAll(userId uint64) (todos []*entity.Todo, err error)
	Create(t *entity.Todo) (todo *entity.Todo, err error)
	Find(todoId int) (todo *entity.Todo, err error)
	Update(todoId int, t *entity.Todo) (todo *entity.Todo, err error)
	Delete(t *entity.Todo) error
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

func (todoRep *todoRepository) Create(t *entity.Todo) (todo *entity.Todo, err error) {
	err = todoRep.DB.Create(&t).Error
	if err != nil {
		return nil, err
	}

	err = todoRep.DB.Find(&todo, &t.ID).Error
	if err != nil {
		return nil, err
	}

	return
}

func (todoRep *todoRepository) Find(todoId int) (todo *entity.Todo, err error) {
	err = todoRep.DB.First(&todo, todoId).Error
	if err != nil {
		return nil, err
	}

	return
}

func (todoRep *todoRepository) Update(todoId int, t *entity.Todo) (todo *entity.Todo, err error) {
	err = todoRep.DB.Updates(&t).Error
	if err != nil {
		return nil, err
	}

	err = todoRep.DB.Find(&todo, &t.ID).Error
	if err != nil {
		return nil, err
	}

	return
}

func (todoRep *todoRepository) Delete(t *entity.Todo) error {
	if err := todoRep.DB.
		Delete(&t).
		Error; err != nil {
		return err
	}

	return nil
}
