package delivery

import (
	"time"

	"github.com/google/uuid"
)

type Delivery struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	templateId uuid.UUID
	sendAt     time.Time
	sended     bool
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,

	templateId uuid.UUID,
	sendAt time.Time,
	sended bool,
) (*Delivery, error) {
	return &Delivery{
		id:         id,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
		templateId: templateId,
		sendAt:     sendAt,
		sended:     sended,
	}, nil
}

func New(
	templateId uuid.UUID,
	sendAt time.Time,
) (*Delivery, error) {
	return &Delivery{
		id:         uuid.New(),
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
		sendAt:     sendAt,
		templateId: templateId,
		sended:     false,
	}, nil
}

func (d *Delivery) ID() uuid.UUID {
	return d.id
}

func (d *Delivery) CreatedAt() time.Time {
	return d.createdAt
}

func (d *Delivery) ModifiedAt() time.Time {
	return d.modifiedAt
}

func (d *Delivery) TemplateID() uuid.UUID {
	return d.templateId
}

func (d *Delivery) SendAt() time.Time {
	return d.sendAt
}

func (d *Delivery) IsSended() bool {
	return d.sended
}
