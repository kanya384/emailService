package repository

import (
	"emailservice/internal/repository/postgres/delivery"
	"emailservice/internal/repository/postgres/subscriber"
	"emailservice/internal/repository/postgres/template"
	"emailservice/pkg/psql"
)

type Repository struct {
	Delivery   *delivery.Repository
	Subscriber *subscriber.Repository
	Template   *template.Repository
}

func NewRepository(pg *psql.Postgres) (*Repository, error) {

	subscriber, err := subscriber.New(pg)
	if err != nil {
		return nil, err
	}

	delivery, err := delivery.New(pg, subscriber)
	if err != nil {
		return nil, err
	}

	template, err := template.New(pg)
	if err != nil {
		return nil, err
	}

	return &Repository{
		delivery,
		subscriber,
		template,
	}, nil

}
