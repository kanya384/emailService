package template

import (
	"testing"
	"time"

	"emailservice/internal/domain/template/path"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewWithID(t *testing.T) {
	req := require.New(t)
	id := uuid.New()
	timeNow := time.Now()
	path := path.Path("/template.html")

	t.Run("create template with id success", func(t *testing.T) {
		template, err := NewWithID(id, path, timeNow, timeNow)
		req.Equal(err, nil)
		req.Equal(template.ID(), id)
		req.Equal(template.CreatedAt(), timeNow)
		req.Equal(template.ModifiedAt(), timeNow)
		req.Equal(template.Path(), path)
	})
}

func TestNew(t *testing.T) {
	req := require.New(t)
	path := path.Path("/template.html")

	t.Run("create template success", func(t *testing.T) {
		template, err := New(path)
		req.Equal(err, nil)
		req.Equal(template.Path(), path)
		req.NotEmpty(template.ID())
		req.NotEmpty(template.CreatedAt())
		req.NotEmpty(template.ModifiedAt())
	})
}
