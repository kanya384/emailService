package surname

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testSurname, _ := NewSurname("Иванов")
	tests := map[string]struct {
		input *Surname
		want  string
	}{
		"success": {input: testSurname, want: "Иванов"},
	}

	for surname, testCase := range tests {
		t.Run(surname, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewSurname(t *testing.T) {
	req := require.New(t)
	testSurname, _ := NewSurname("Иванов")
	tests := map[string]struct {
		input string
		want  *Surname
		err   error
	}{
		"success":      {input: "Иванов", want: testSurname, err: nil},
		"error length": {input: "И", want: nil, err: ErrWrongLength},
	}

	for testSurname, testCase := range tests {
		t.Run(testSurname, func(t *testing.T) {
			res, err := NewSurname(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
