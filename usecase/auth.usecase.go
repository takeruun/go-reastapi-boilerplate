package usecase

import (
	"app/database"
	"app/service"
	"context"
)

type AuthUsecase interface {
	SignIn(context.Context) error
	SignUp() error
}

type authUsecase struct {
	userRepo database.UserRepository
	sessionS service.SessionService
}

func NewAuthUsecase(userRepo database.UserRepository, sessionS service.SessionService) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		sessionS: sessionS,
	}
}

func (uu *authUsecase) SignIn(ctx context.Context) error {
	return nil
}

func (uu *authUsecase) SignUp() error {
	return nil
}
