package delivery

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewWithID(t *testing.T) {
	req := require.New(t)
	id := uuid.New()
	timeNow := time.Now()
	templateID := uuid.New()

	t.Run("create delivery with id success", func(t *testing.T) {
		delivery, err := NewWithID(id, timeNow, timeNow, templateID, timeNow, false)
		req.Equal(err, nil)
		req.Equal(delivery.ID(), id)
		req.Equal(delivery.CreatedAt(), timeNow)
		req.Equal(delivery.ModifiedAt(), timeNow)
		req.Equal(delivery.TemplateID(), templateID)
		req.Equal(delivery.SendAt(), timeNow)
		req.Equal(delivery.IsSended(), false)
	})
}

func TestNew(t *testing.T) {
	req := require.New(t)
	timeNow := time.Now()
	templateID := uuid.New()

	t.Run("create delivery with id success", func(t *testing.T) {
		delivery, err := New(templateID, timeNow)
		req.Equal(err, nil)
		req.NotEmpty(delivery.ID())
		req.NotEmpty(delivery.CreatedAt())
		req.NotEmpty(delivery.ModifiedAt())
		req.Equal(delivery.TemplateID(), templateID)
		req.Equal(delivery.SendAt(), timeNow)
		req.Equal(delivery.IsSended(), false)
	})
}
