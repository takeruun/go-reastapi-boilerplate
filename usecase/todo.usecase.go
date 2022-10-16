package usecase

import (
	"app/database"
	"app/entity"
	"app/service"
	"context"
)

type TodoUsecase interface {
	FindAll(ctx context.Context) (todos []*entity.Todo, err error)
}

type todoUsecase struct {
	userRepo database.TodoRepository
	sessionS service.SessionService
}

func NewTodoUsecase(userRepo database.TodoRepository, sessionS service.SessionService) TodoUsecase {
	return &todoUsecase{
		userRepo: userRepo,
		sessionS: sessionS,
	}
}

func (uu *todoUsecase) FindAll(ctx context.Context) (todos []*entity.Todo, err error) {
	session, _ := uu.sessionS.GetSession(ctx, "session")
	userId := session.Values["userId"].(uint64)

	todos, err = uu.userRepo.FindAll(userId)

	return
}
