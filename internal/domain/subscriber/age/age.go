package age

import (
	"fmt"
)

const (
	MinAge = 10
	MaxAge = 100
)

var (
	ErrNotValidAge = fmt.Errorf("provided age is not valid, it must be between %d and %d", MinAge, MaxAge)
)

type Age int

func NewAge(age int) (*Age, error) {
	if age < MinAge || age > MaxAge {
		return nil, ErrNotValidAge
	}

	n := Age(age)

	return &n, nil
}
