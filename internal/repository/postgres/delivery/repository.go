package delivery

import (
	"emailservice/internal/repository/postgres/subscriber"
	"emailservice/pkg/psql"
)

type Repository struct {
	pg             *psql.Postgres
	repoSubscriber subscriber.Subscriber
}

func New(pg *psql.Postgres, repoSubscriber subscriber.Subscriber) (*Repository, error) {
	var r = &Repository{
		pg:             pg,
		repoSubscriber: repoSubscriber,
	}
	return r, nil
}
