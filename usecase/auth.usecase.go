package usecase

import (
	"app/controller/dto"
	"app/database"
	"app/entity"
	"app/service"
	"context"
	"errors"
)

type AuthUsecase interface {
	SignIn(ctx context.Context, signInParams *dto.AuthSignInRequestDto) error
	SignUp(ctx context.Context, signInParams *dto.AuthSignUpRequestDto) error
	Show(ctx context.Context) (user *entity.User, err error)
	Edit(ctx context.Context, userUpdateParams *dto.AuthUserUpdateRequestDto) (user *entity.User, err error)
	Delete(ctx context.Context) error
}

type authUsecase struct {
	userRepo database.UserRepository
	sessionS service.SessionService
	cyptoS   service.CyptoService
	mailS    service.MailService
}

func NewAuthUsecase(userRepo database.UserRepository, sessionS service.SessionService, cyptoS service.CyptoService, mailS service.MailService) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		sessionS: sessionS,
		cyptoS:   cyptoS,
		mailS:    mailS,
	}
}

func (uu *authUsecase) SignIn(ctx context.Context, signInParams *dto.AuthSignInRequestDto) error {
	loginUser, err := uu.userRepo.FindByEmail(signInParams.Email)
	if err != nil {
		return err
	}

	if !uu.cyptoS.ComparePasswords(loginUser.HashPassword, []byte(signInParams.Password)) {
		return errors.New("Authentication Failure")
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

	u := entity.User{Name: signInParams.Name, Email: signInParams.Email, HashPassword: hashPwd}
	loginUser, err := uu.userRepo.Create(&u)
	if err != nil {
		return err
	}

	err = uu.sessionS.SaveSession(ctx, "userId", loginUser.ID)
	if err != nil {
		return err
	}

	return nil
}

func (uu *authUsecase) Show(ctx context.Context) (user *entity.User, err error) {
	session, _ := uu.sessionS.GetSession(ctx, "_goreset_session")
	userId := session.Values["userId"].(uint64)

	user, err = uu.userRepo.Find(userId)
	if err != nil {
		return nil, err
	}

	return
}

func (uu *authUsecase) Edit(ctx context.Context, userUpdateParams *dto.AuthUserUpdateRequestDto) (user *entity.User, err error) {
	session, _ := uu.sessionS.GetSession(ctx, "_goreset_session")
	userId := session.Values["userId"].(uint64)

	u := &entity.User{
		ID:    userId,
		Name:  userUpdateParams.Name,
		Email: userUpdateParams.Email,
	}

	user, err = uu.userRepo.Update(u)
	if err != nil {
		return nil, err
	}

	return
}

func (uu *authUsecase) Delete(ctx context.Context) error {
	session, _ := uu.sessionS.GetSession(ctx, "_goreset_session")
	userId := session.Values["userId"].(uint64)

	if err := uu.userRepo.Delete(userId); err != nil {
		return nil
	}

	uu.sessionS.DeleteSession(ctx)

	return nil
}
