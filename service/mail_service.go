package service

import (
	"app/config"
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

type MailService interface {
	Send(to string, subject string, body string) error
}

type mailService struct {
	Mail *config.MailServer
}

func NewMailService(Mail *config.MailServer) MailService {
	return &mailService{
		Mail: Mail,
	}
}

func (m *mailService) writeString(b *bytes.Buffer, s string) *bytes.Buffer {
	_, err := b.WriteString(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	return b
}

// サブジェクトを MIME エンコードする
func (m *mailService) encodeSubject(subject string) string {
	// UTF8 文字列を指定文字数で分割する
	b := bytes.NewBuffer([]byte(""))
	strs := []string{}
	length := 13
	for k, c := range strings.Split(subject, "") {
		b.WriteString(c)
		if k%length == length-1 {
			strs = append(strs, b.String())
			b.Reset()
		}
	}
	if b.Len() > 0 {
		strs = append(strs, b.String())
	}
	// MIME エンコードする
	b2 := bytes.NewBuffer([]byte(""))
	b2.WriteString("Subject:")
	for _, line := range strs {
		b2.WriteString(" =?utf-8?B?")
		b2.WriteString(base64.StdEncoding.EncodeToString([]byte(line)))
		b2.WriteString("?=\r\n")
	}
	return b2.String()
}

// 本文を 76 バイト毎に CRLF を挿入して返す
func (m *mailService) encodeBody(body string) string {
	b := bytes.NewBufferString(body)
	s := base64.StdEncoding.EncodeToString(b.Bytes())
	b2 := bytes.NewBuffer([]byte(""))
	for k, c := range strings.Split(s, "") {
		b2.WriteString(c)
		if k%76 == 75 {
			b2.WriteString("\r\n")
		}
	}
	return b2.String()
}

func (m *mailService) Send(to string, subject string, body string) error {
	msg := bytes.NewBuffer([]byte(""))
	msg = m.writeString(msg, "From: "+m.Mail.From.String()+"\r\n")
	msg = m.writeString(msg, "To: "+to+"\r\n")
	// msg = m.writeString(msg, "Subject: " + subject + "\r\n")
	msg = m.writeString(msg, m.encodeSubject(subject))
	msg = m.writeString(msg, "MIME-Version: 1.0\r\n")
	msg = m.writeString(msg, "Content-Type: text/plain; charset=\"utf-8\"\r\n")
	msg = m.writeString(msg, "Content-Transfer-Encoding: base64\r\n")
	msg = m.writeString(msg, "\r\n")

	// msg = m.writeString(msg, base64.StdEncoding.EncodeToString([]byte(body)) + "\r\n")
	msg = m.writeString(msg, m.encodeBody(body))

	if err := smtp.SendMail(m.Mail.Addr, m.Mail.Auth, m.Mail.From.Address, []string{to}, msg.Bytes()); err != nil {
		return err
	}

	fmt.Print(msg, "\n")
	fmt.Print(body, "\n")

	return nil
}
