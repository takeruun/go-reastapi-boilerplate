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
	session, _ := tu.sessionS.GetSession(ctx, "session")
	userId := session.Values["userId"].(uint64)

	todos, err = tu.todoRepo.FindAll(userId)

	return
}

func (tu *todoUsecase) Create(ctx context.Context, createParams *dto.TodoCreateRequestDto) (todo *entity.Todo, err error) {
	session, _ := tu.sessionS.GetSession(ctx, "session")
	userId := session.Values["userId"].(uint64)

	entity := entity.Todo{Title: createParams.Title, Description: createParams.Description, UserId: userId}
	todo, err = tu.todoRepo.Create(&entity)
	if err != nil {
		return nil, err
	}

	return
}

func (tu *todoUsecase) Show(ctx context.Context, todoId int) (todo *entity.Todo, err error) {
	session, _ := tu.sessionS.GetSession(ctx, "session")
	userId := session.Values["userId"].(uint64)

	todo, err = tu.todoRepo.Get(todoId)
	if err != nil {
		return nil, err
	}

	if userId != todo.UserId {
		return nil, errors.New("no your todo")
	}

	return
}
