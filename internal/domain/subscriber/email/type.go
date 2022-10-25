package email

import (
	"errors"
	"regexp"
)

var (
	ErrNotValidEmail = errors.New("provided email is not valid")
)

type Email string

func (n Email) String() string {
	return string(n)
}

func NewEmail(email string) (*Email, error) {

	if !isEmailValid(email) {
		return nil, ErrNotValidEmail
	}

	n := Email(email)

	return &n, nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
