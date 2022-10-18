package config

import (
	"errors"
	"fmt"
	"net/mail"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

// LoginAuth creates new loginAuth strcut as smtp.Auth interface.
// loginAuth handle LOGIN type authentication in SMTP.
func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	if !server.TLS {
		advertised := false
		for _, mechanism := range server.Auth {
			if mechanism == "LOGIN" {
				advertised = true
				break
			}
		}
		if !advertised {
			return "", nil, errors.New("unencrypted connection")
		}
	}
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, fmt.Errorf("unkown message from server %s", fromServer)
		}
	}
	return nil, nil
}

type MailServer struct {
	Auth smtp.Auth
	Addr string
	From mail.Address
}

func NewMailServer() *MailServer {
	c := NewConfig()
	return &MailServer{
		Auth: LoginAuth(c.Mail.Auth.Email, c.Mail.Auth.Password),
		Addr: c.Mail.Addr,
		From: mail.Address{Name: c.Mail.FromName, Address: c.Mail.FromEmail},
	}
}
