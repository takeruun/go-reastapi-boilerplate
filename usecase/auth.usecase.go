package usecase

import (
	"app/controller/dto"
	"app/database"
	"app/entity"
	"app/service"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	SignIn(ctx context.Context, signInParams *dto.AuthSignInRequestDto) error
	SignUp(ctx context.Context, signInParams *dto.AuthSignUpRequestDto) error
}

type authUsecase struct {
	userRepo database.UserRepository
	sessionS service.SessionService
	cyptoS   service.CyptoService
}

func NewAuthUsecase(userRepo database.UserRepository, sessionS service.SessionService, cyptoS service.CyptoService) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		sessionS: sessionS,
		cyptoS:   cyptoS,
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

func (uu *authUsecase) SignUp(ctx context.Context, signInParams *dto.AuthSignUpRequestDto) error {
	hashPwd, err := uu.cyptoS.HashAndSalt([]byte(signInParams.Password))
	if err != nil {
		return err
	}

	entity := entity.User{Name: signInParams.Name, Email: signInParams.Email, HashPassword: hashPwd}
	loginUser, err := uu.userRepo.Create(&entity)
	if err != nil {
		return err
	}

	err = uu.sessionS.SaveSession(ctx, "userId", loginUser.ID)
	if err != nil {
		return err
	}

	return nil
}
