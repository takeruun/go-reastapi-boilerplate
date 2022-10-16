package usecase

import (
	"app/database"
	"app/entity"
)

type UserUsecase interface {
	FindAll() (users *[]entity.User, err error)
}

type userUsecase struct {
	userRepo database.UserRepository
}

func NewUserUsecase(userRepo database.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uu *userUsecase) FindAll() (users *[]entity.User, err error) {
	return uu.userRepo.FindAll()
}
