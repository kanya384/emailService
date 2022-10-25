package surname

import (
	"fmt"
)

const (
	MaxLength = 300
	MinLength = 5
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the surname must be between %d and %d", MinLength, MaxLength)
)

type Surname string

func (n Surname) String() string {
	return string(n)
}

func NewSurname(surname string) (*Surname, error) {
	if (len([]rune(surname))) < MinLength || (len([]rune(surname))) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Surname(surname)

	return &n, nil
}
