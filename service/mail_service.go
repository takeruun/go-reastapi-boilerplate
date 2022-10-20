package service

import (
	"app/config"
)

type MailService interface {
	SendMail(to, subject, body string) error
}

type mailService struct {
	Mail *config.Mail
}

func NewMailService(Mail *config.Mail) MailService {
	return &mailService{
		Mail: Mail,
	}
}

func (m *mailService) SendMail(to, subject, body string) error {
	if err := m.Mail.Send(to, subject, body); err != nil {
		return err
	}

	return nil
}
