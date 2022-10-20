package config

import "github.com/takeruun/gomail"

type Mail struct {
	*gomail.Mail
}

func NewMail() *Mail {
	config := NewConfig()
	gomailConfig := &gomail.Config{
		Auth: struct {
			Host     string
			Email    string
			Password string
		}{
			Host:     config.Mail.Auth.Host,
			Email:    config.Mail.Auth.Email,
			Password: config.Mail.Auth.Password,
		},
		From: struct {
			Name  string
			Email string
		}{
			Name:  config.Mail.From.Name,
			Email: config.Mail.From.Email,
		},
		Addr: config.Mail.Addr,
	}

	return &Mail{
		gomail.New(gomailConfig),
	}
}
