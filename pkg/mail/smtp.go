package mail

import (
	"fmt"
	"net/smtp"
)

type Sender interface {
	SendEmail(m *message) (err error)
}

type sender struct {
	host string
	port int
	auth smtp.Auth
}

func NewSender(host string, port int, username, pass string) *sender {

	auth := smtp.PlainAuth("", username, pass, host)

	return &sender{
		auth: auth,
		host: host,
		port: port,
	}
}

func (s *sender) address() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func (s *sender) SendEmail(m *message) (err error) {

	err = smtp.SendMail(s.address(), s.auth, m.From(), []string{m.To()}, m.GenerateMessage())
	return
}
