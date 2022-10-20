package usecase

import (
	"app/controller/dto"
	"app/database"
	"app/entity"
	"app/service"
	"context"
	"errors"
)

type TodoUsecase interface {
	FindAll(ctx context.Context) (todos []*entity.Todo, err error)
	Create(ctx context.Context, createParams *dto.TodoCreateRequestDto) (todos *entity.Todo, err error)
	Show(ctx context.Context, todoId int) (todo *entity.Todo, err error)
	Edit(todoId int, updateParams *dto.TodoUpdateRequestDto) (todo *entity.Todo, err error)
	Delete(ctx context.Context, todoId int) error
}

type todoUsecase struct {
	todoRepo database.TodoRepository
	sessionS service.SessionService
}

func NewTodoUsecase(todoRepo database.TodoRepository, sessionS service.SessionService) TodoUsecase {
	return &todoUsecase{
		todoRepo: todoRepo,
		sessionS: sessionS,
	}
}

func (tu *todoUsecase) FindAll(ctx context.Context) (todos []*entity.Todo, err error) {
	userId, _ := tu.sessionS.GetSessionValue(ctx, "userId")

	todos, err = tu.todoRepo.FindAll(userId.(uint64))

	return
}

func (tu *todoUsecase) Create(ctx context.Context, createParams *dto.TodoCreateRequestDto) (todo *entity.Todo, err error) {
	userId, _ := tu.sessionS.GetSessionValue(ctx, "userId")

	t := entity.Todo{Title: createParams.Title, Description: createParams.Description, UserId: userId.(uint64)}
	todo, err = tu.todoRepo.Create(&t)
	if err != nil {
		return nil, err
	}

	return
}

func (tu *todoUsecase) Show(ctx context.Context, todoId int) (todo *entity.Todo, err error) {
	userId, _ := tu.sessionS.GetSessionValue(ctx, "userId")

	todo, err = tu.todoRepo.Get(todoId)
	if err != nil {
		return nil, err
	}

	if userId.(uint64) != todo.UserId {
		return nil, errors.New("no your todo")
	}

	return
}

func (tu *todoUsecase) Edit(todoId int, updateParams *dto.TodoUpdateRequestDto) (todo *entity.Todo, err error) {
	t := entity.Todo{ID: uint64(todoId), Title: updateParams.Title, Description: updateParams.Description}
	todo, err = tu.todoRepo.Update(todoId, &t)
	if err != nil {
		return nil, err
	}

	return
}

func (tu *todoUsecase) Delete(ctx context.Context, todoId int) error {
	if err := tu.todoRepo.Delete(&entity.Todo{ID: uint64(todoId)}); err != nil {
		return err
	}

	return nil
}
