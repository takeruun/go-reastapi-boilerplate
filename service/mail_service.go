package service

import (
	"app/config"
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

const (
	charactorLimitForOneLine = 78
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

// 適切な長さにカットしCRLFを挿入
func (m *mailService) cutAndAddCrlf(msg string) string {
	buffer := bytes.Buffer{}
	for k, c := range strings.Split(msg, "") {
		buffer.WriteString(c)
		if (k+1)%charactorLimitForOneLine == 0 {
			buffer.WriteString("\r\n")
		}
	}
	return buffer.String()
}

func (m *mailService) makeMailBody(body string) string {
	encodedBody := base64.StdEncoding.EncodeToString([]byte(body))
	return m.cutAndAddCrlf(encodedBody)
}

// UTF8文字列を指定文字数で分割
func (m *mailService) utf8Split(utf8string string, length int) []string {
	result := []string{}
	buffer := bytes.Buffer{}
	for k, c := range strings.Split(utf8string, "") {
		buffer.WriteString(c)
		if (k+1)%length == 0 {
			result = append(result, buffer.String())
			buffer.Reset()
		}
	}
	if buffer.Len() > 0 {
		result = append(result, buffer.String())
	}
	return result
}

// タイトルをMIMEエンコード
func (m *mailService) encodeSubject(subject string) string {
	buffer := bytes.Buffer{}
	buffer.WriteString("Subject:")
	limit := charactorLimitForOneLine / 6 // Unicodeでは一文字が最大6バイトになるため
	for _, line := range m.utf8Split(subject, limit) {
		buffer.WriteString(" =?utf-8?B?")
		buffer.WriteString(base64.StdEncoding.EncodeToString([]byte(line)))
		buffer.WriteString("?=\r\n")
	}
	return buffer.String()
}

// ヘッダを作る
func (m *mailService) makeMailHeader(to, subject string) bytes.Buffer {
	header := bytes.Buffer{}
	header.WriteString("From: " + m.Mail.From.String() + "\r\n")
	header.WriteString("To: " + to + "\r\n")
	header.WriteString(m.encodeSubject(subject))
	header.WriteString("MIME-Version: 1.0\r\n")
	header.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	header.WriteString("Content-Transfer-Encoding: base64\r\n")

	return header
}

func (m *mailService) Send(to, subject, body string) error {
	mailHeader := m.makeMailHeader(to, subject)
	mailBody := m.makeMailBody(body)

	msg := mailHeader
	msg.WriteString(mailBody)

	if err := smtp.SendMail(m.Mail.Addr, m.Mail.Auth, m.Mail.From.Address, []string{to}, msg.Bytes()); err != nil {
		return err
	}

	fmt.Print(msg, "\n")
	fmt.Print(body, "\n")

	return nil
}
