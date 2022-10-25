package email

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testEmail, _ := NewEmail("test01@mail.ru")
	tests := map[string]struct {
		input *Email
		want  string
	}{
		"success": {input: testEmail, want: "test01@mail.ru"},
	}

	for email, testCase := range tests {
		t.Run(email, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewEmail(t *testing.T) {
	req := require.New(t)
	testEmail, _ := NewEmail("test01@mail.ru")
	tests := map[string]struct {
		input string
		want  *Email
		err   error
	}{
		"success":      {input: "test01@mail.ru", want: testEmail, err: nil},
		"error length": {input: "tt ", want: nil, err: ErrNotValidEmail},
	}

	for testEmail, testCase := range tests {
		t.Run(testEmail, func(t *testing.T) {
			res, err := NewEmail(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
