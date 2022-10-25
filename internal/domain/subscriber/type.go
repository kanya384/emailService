package subscriber

import (
	"emailservice/internal/domain/subscriber/age"
	"emailservice/internal/domain/subscriber/email"
	"emailservice/internal/domain/subscriber/name"
	"emailservice/internal/domain/subscriber/surname"
	"time"

	"github.com/google/uuid"
)

type Subscriber struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	name    name.Name
	surname surname.Surname
	email   email.Email
	age     age.Age
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,

	name name.Name,
	surname surname.Surname,
	email email.Email,
	age age.Age,
) (*Subscriber, error) {
	return &Subscriber{
		id:         id,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
		name:       name,
		surname:    surname,
		email:      email,
		age:        age,
	}, nil
}

func New(
	name name.Name,
	surname surname.Surname,
	email email.Email,
	age age.Age,
) (*Subscriber, error) {
	return &Subscriber{
		id:         uuid.New(),
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
		name:       name,
		surname:    surname,
		email:      email,
		age:        age,
	}, nil
}

func (t *Subscriber) ID() uuid.UUID {
	return t.id
}

func (t *Subscriber) Name() name.Name {
	return t.name
}

func (t *Subscriber) Surname() surname.Surname {
	return t.surname
}

func (t *Subscriber) Email() email.Email {
	return t.email
}

func (t *Subscriber) Age() age.Age {
	return t.age
}

func (t *Subscriber) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Subscriber) ModifiedAt() time.Time {
	return t.modifiedAt
}
