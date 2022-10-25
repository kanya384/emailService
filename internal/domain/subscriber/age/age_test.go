package age

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAge(t *testing.T) {
	req := require.New(t)
	tests := map[string]struct {
		input int
		err   error
	}{
		"success":         {input: 25, err: nil},
		"error low value": {input: 5, err: ErrNotValidAge},
	}

	for testEmail, testCase := range tests {
		t.Run(testEmail, func(t *testing.T) {
			_, err := NewAge(testCase.input)
			req.Equal(testCase.err, err)
		})
	}
}
