package subscriber

import (
	"emailservice/internal/domain/subscriber/age"
	"emailservice/internal/domain/subscriber/email"
	"emailservice/internal/domain/subscriber/name"
	"emailservice/internal/domain/subscriber/surname"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewWithID(t *testing.T) {
	req := require.New(t)
	id := uuid.New()
	timeNow := time.Now()
	name, _ := name.NewName("tests")
	surname, _ := surname.NewSurname("tests")
	email, _ := email.NewEmail("test01@mail.ru")
	age, _ := age.NewAge(55)

	t.Run("create subscriber with id success", func(t *testing.T) {
		subscriber, err := NewWithID(id, timeNow, timeNow, *name, *surname, *email, *age)
		req.Equal(err, nil)
		req.Equal(subscriber.ID(), id)
		req.Equal(subscriber.CreatedAt(), timeNow)
		req.Equal(subscriber.ModifiedAt(), timeNow)
		req.Equal(subscriber.Name(), *name)
		req.Equal(subscriber.Email(), *email)
		req.Equal(subscriber.Surname(), *surname)
		req.Equal(subscriber.Age(), *age)
	})
}

func TestNew(t *testing.T) {
	req := require.New(t)

	name, _ := name.NewName("tests")
	surname, _ := surname.NewSurname("tests")
	email, _ := email.NewEmail("test01@mail.ru")
	age, _ := age.NewAge(55)

	t.Run("create subscriber success", func(t *testing.T) {
		subscriber, err := New(*name, *surname, *email, *age)
		req.Equal(err, nil)
		req.Equal(subscriber.Name(), *name)
		req.Equal(subscriber.Email(), *email)
		req.Equal(subscriber.Surname(), *surname)
		req.Equal(subscriber.Age(), *age)
	})
}
