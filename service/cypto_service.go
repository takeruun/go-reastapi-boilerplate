package service

import "golang.org/x/crypto/bcrypt"

type CyptoService interface {
	HashAndSalt(pwd []byte) (string, error)
	ComparePasswords(hashedPwd string, plainPwd []byte) bool
}

type cyptoService struct{}

func NewCyptoService() CyptoService {
	return &cyptoService{}
}

func (c *cyptoService) HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (c *cyptoService) ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
