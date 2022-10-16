package usecase

import (
	"app/database"
)

type AuthUsecase interface {
	SignIn() error
}

type authUsecase struct {
	userRepo database.UserRepository
}

func NewAuthUsecase(userRepo database.UserRepository) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
	}
}

func (uu *authUsecase) SignIn() error {
	return nil
}
