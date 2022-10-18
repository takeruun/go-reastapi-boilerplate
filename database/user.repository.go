package database

import (
	"app/config"
	"app/entity"
)

type UserRepository interface {
	FindAll() (uesrs []*entity.User, err error)
	FindByEmail(email string) (user *entity.User, err error)
	Create(u *entity.User) (user *entity.User, err error)
}

type userRepository struct {
	DB *config.DB
}

func NewUserRepository(DB *config.DB) UserRepository {
	return &userRepository{DB: DB}
}

func (userRep *userRepository) FindAll() (users []*entity.User, err error) {
	err = userRep.DB.Model(&entity.User{}).
		Find(&users).
		Error

	if err != nil {
		return nil, err
	}

	return
}

func (userRep *userRepository) FindByEmail(email string) (user *entity.User, err error) {
	err = userRep.DB.First(&user, "email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return
}

func (userRep *userRepository) Create(u *entity.User) (user *entity.User, err error) {
	err = userRep.DB.Create(&u).Error
	if err != nil {
		return nil, err
	}

	err = userRep.DB.Find(&user, &u.ID).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
