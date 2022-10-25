package template

import (
	"emailservice/pkg/psql"
)

type Repository struct {
	*psql.Postgres
}

func New(pg *psql.Postgres) (*Repository, error) {
	var r = &Repository{pg}
	return r, nil
}
