package name

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testName, _ := NewName("Василий")
	tests := map[string]struct {
		input *Name
		want  string
	}{
		"success": {input: testName, want: "Василий"},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}

}

func TestNewName(t *testing.T) {
	req := require.New(t)
	testName := Name("Василий")
	tests := map[string]struct {
		input string
		want  *Name
		err   error
	}{
		"success":      {input: "Василий", want: &testName, err: nil},
		"error length": {input: "вв", want: nil, err: ErrWrongLength},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			res, err := NewName(testCase.input)
			req.Equal(testCase.want, res)
			req.Equal(testCase.err, err)
		})
	}
}
