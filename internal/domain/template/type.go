package template

import (
	"time"

	"emailservice/internal/domain/template/path"

	"github.com/google/uuid"
)

type Template struct {
	id         uuid.UUID
	createdAt  time.Time
	modifiedAt time.Time

	path path.Path
}

func NewWithID(
	id uuid.UUID,
	path path.Path,
	createdAt time.Time,
	modifiedAt time.Time,
) (*Template, error) {
	return &Template{
		id:         id,
		path:       path,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
	}, nil
}

func New(
	path path.Path,
) (*Template, error) {
	return &Template{
		id:         uuid.New(),
		path:       path,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}, nil
}

func (t *Template) ID() uuid.UUID {
	return t.id
}

func (t *Template) Path() path.Path {
	return t.path
}

func (t *Template) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Template) ModifiedAt() time.Time {
	return t.modifiedAt
}
