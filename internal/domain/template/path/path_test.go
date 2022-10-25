package path

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	req := require.New(t)
	testPath := Path("./cmd/app/main.go")
	tests := map[string]struct {
		input *Path
		want  string
	}{
		"success": {input: &testPath, want: "./cmd/app/main.go"},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			res := testCase.input.String()
			req.Equal(testCase.want, res)
		})
	}
}

func TestNewPath(t *testing.T) {
	req := require.New(t)
	testPath := "./test_file.html"
	file, err := os.Create(testPath)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		file.Close()
		os.Remove(file.Name())
	}(file)

	t.Run("success", func(t *testing.T) {
		res, err := NewPath(testPath)
		req.Equal(testPath, res.String())
		req.Empty(err)

	})

	t.Run("no such file", func(t *testing.T) {
		res, err := NewPath("not_existing_path")
		req.Empty(res)
		req.NotEmpty(err)
	})

}
