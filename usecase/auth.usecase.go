package usecase

import (
	"app/controller/dto"
	"app/database"
	"app/service"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	SignIn(ctx context.Context, signInParams *dto.AuthSignInRequestDto) error
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

func (uu *authUsecase) SignIn(ctx context.Context, signInParams *dto.AuthSignInRequestDto) error {
	loginUser, err := uu.userRepo.FindByEmail(signInParams.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginUser.HashPassword), []byte(signInParams.Password))
	if err != nil {
		return err
	}

	err = uu.sessionS.SaveSession(ctx, "userId", loginUser.ID)
	if err != nil {
		return err
	}

	return nil
}

func (uu *authUsecase) SignUp() error {
	return nil
}
