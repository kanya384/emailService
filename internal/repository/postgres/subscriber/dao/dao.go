package dao

import (
	"time"

	"github.com/google/uuid"
)

type Subscriber struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Surname    string    `db:"surname"`
	Email      string    `db:"email"`
	Age        int       `db:"age"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}

var ColumnsSubscriber = []string{
	"id",
	"created_at",
	"modified_at",
	"name",
	"surname",
	"email",
	"age",
}

var CreateColumnSubscriberInDelivery = []string{
	"id",
	"delivery_id",
	"subscriber_id",
	"opened",
}
