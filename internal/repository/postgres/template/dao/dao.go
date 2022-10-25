package dao

import (
	"time"

	"github.com/google/uuid"
)

type Template struct {
	ID         uuid.UUID `db:"id"`
	Path       string    `db:"path"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}

var ColumnsTemplate = []string{
	"id",
	"path",
	"created_at",
	"modified_at",
}
