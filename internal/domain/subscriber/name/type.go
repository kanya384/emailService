package name

import (
	"fmt"
)

const (
	MaxLength = 100
	MinLength = 3
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the name must be between %d and %d", MinLength, MaxLength)
)

type Name string

func (n Name) String() string {
	return string(n)
}

func NewName(name string) (*Name, error) {
	if (len([]rune(name))) < MinLength || (len([]rune(name))) > MaxLength {
		return nil, ErrWrongLength
	}

	n := Name(name)

	return &n, nil
}
