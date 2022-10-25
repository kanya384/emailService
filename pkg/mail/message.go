package mail

import (
	"fmt"
)

type message struct {
	from        string
	to          string
	subject     string
	body        string
	contentType string
}

func NewMessage() *message {
	return &message{}
}

func (m *message) SetFrom(from string) {
	m.from = from
}

func (m message) From() string {
	return m.from
}

func (m *message) SetTo(to string) {
	m.to = to
}

func (m message) To() string {
	return m.to
}

func (m *message) SetSubject(subject string) {
	m.subject = subject
}

func (m message) Subject() string {
	return m.subject
}

func (m *message) SetBody(body string) {
	m.body = body
}

func (m message) Body() []byte {
	return []byte(m.body)
}

func (m *message) SetContentType(contentType string) {
	m.contentType = contentType
}

func (m message) ContentType() string {
	return m.contentType
}

func (m message) GenerateMessage() []byte {
	return []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME: MIME-version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\";\r\n"+
			"\n\n"+
			"%s",
		m.From(), m.To(), m.Subject(), m.Body()))
}
