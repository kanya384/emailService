package path

import (
	"fmt"
	"os"
)

const (
	MinLength = 5
)

var (
	ErrWrongLength = fmt.Errorf("number of characters in the path must be at least %d characters", MinLength)
)

type Path string

func (n Path) String() string {
	return string(n)
}

func NewPath(path string) (*Path, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	res := Path(path)

	return &res, nil
}
