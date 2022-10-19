package usecase_test

import (
	"app/controller/dto"
	"app/entity"
	"app/test/mock_database"
	"app/test/mock_service"
	"app/usecase"
	"errors"

	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockUserRepository *mock_database.MockUserRepository
var mockCyptoService *mock_service.MockCyptoService
var mockSessionService *mock_service.MockSessionService
var mockMailServive *mock_service.MockMailService
var authUsecase usecase.AuthUsecase

func setUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository = mock_database.NewMockUserRepository(ctrl)
	mockCyptoService = mock_service.NewMockCyptoService(ctrl)
	mockSessionService = mock_service.NewMockSessionService(ctrl)
	mockMailServive = mock_service.NewMockMailService(ctrl)

	return func() {}
}

func TestSignIn(t *testing.T) {
	setup := setUp(t)
	defer setup()

	var (
		email        = "test@example.com"
		expectedUser = entity.User{ID: 1, Email: email, Name: "test", HashPassword: "$2a$10$dipu0jX8IYvy.r.XsecEt.gF4XYsYmBIheZotlxLekYC9UKQWHKe2"}
		password     = "password"
	)

	t.Run("success", func(t *testing.T) {
		mockUserRepository.EXPECT().FindByEmail(email).Return(&expectedUser, nil)
		mockCyptoService.EXPECT().ComparePasswords(expectedUser.HashPassword, []byte(password)).Return(true)
		mockSessionService.EXPECT().SaveSession(gomock.Any(), "userId", expectedUser.ID)

		authUsecase = usecase.NewAuthUsecase(
			mockUserRepository,
			mockSessionService,
			mockCyptoService,
			mockMailServive,
		)

		err := authUsecase.SignIn(context.TODO(), &dto.AuthSignInRequestDto{Email: email, Password: password})

		assert.NoError(t, err)
	})

	t.Run("If the authentication fails", func(t *testing.T) {
		mockUserRepository.EXPECT().FindByEmail(email).Return(&expectedUser, nil)
		mockCyptoService.EXPECT().ComparePasswords(expectedUser.HashPassword, []byte(password)).Return(false)
		mockSessionService.EXPECT().SaveSession(gomock.Any(), "userId", expectedUser.ID)

		authUsecase = usecase.NewAuthUsecase(
			mockUserRepository,
			mockSessionService,
			mockCyptoService,
			mockMailServive,
		)

		err := authUsecase.SignIn(context.TODO(), &dto.AuthSignInRequestDto{Email: email, Password: password})

		assert.Error(t, err)
		assert.EqualError(t, err, "Authentication Failure")
	})
}

func TestSignUp(t *testing.T) {
	setup := setUp(t)
	defer setup()

	var (
		email        = "test@example.com"
		password     = "password"
		name         = "test"
		hashPassword = "$2a$10$dipu0jX8IYvy.r.XsecEt.gF4XYsYmBIheZotlxLekYC9UKQWHKe2"
		expectedUser = entity.User{ID: 1, Email: email, Name: name, HashPassword: hashPassword}
	)

	t.Run("success", func(t *testing.T) {
		mockCyptoService.EXPECT().HashAndSalt([]byte(password)).Return(hashPassword, nil)
		mockUserRepository.EXPECT().Create(&entity.User{Email: email, HashPassword: hashPassword, Name: name}).Return(&expectedUser, nil)
		mockSessionService.EXPECT().SaveSession(gomock.Any(), "userId", expectedUser.ID)

		authUsecase = usecase.NewAuthUsecase(
			mockUserRepository,
			mockSessionService,
			mockCyptoService,
			mockMailServive,
		)

		err := authUsecase.SignUp(context.TODO(), &dto.AuthSignUpRequestDto{Email: email, Password: password, Name: name})

		assert.NoError(t, err)
	})

	t.Run("If hashing fails", func(t *testing.T) {
		err := errors.New("")

		mockCyptoService.EXPECT().HashAndSalt([]byte(password)).Return(hashPassword, err)
		mockUserRepository.EXPECT().Create(&entity.User{Email: email, HashPassword: hashPassword, Name: name}).Return(&expectedUser, nil)
		mockSessionService.EXPECT().SaveSession(gomock.Any(), "userId", expectedUser.ID)

		authUsecase = usecase.NewAuthUsecase(
			mockUserRepository,
			mockSessionService,
			mockCyptoService,
			mockMailServive,
		)

		err = authUsecase.SignUp(context.TODO(), &dto.AuthSignUpRequestDto{Email: email, Password: password, Name: name})

		assert.Error(t, err)
	})
}
