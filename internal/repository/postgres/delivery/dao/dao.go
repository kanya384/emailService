package dao

import (
	"time"
)

type Delivery struct {
	ID         string    `db:"id"`
	TemplateId string    `db:"template_id"`
	SendAt     time.Time `db:"send_at"`
	Sended     bool      `db:"sended"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}

var ColumnsDelivery = []string{
	"id",
	"template_id",
	"send_at",
	"sended",
	"created_at",
	"modified_at",
}
