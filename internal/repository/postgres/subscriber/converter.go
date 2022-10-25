package subscriber

import (
	"emailservice/internal/domain/subscriber"
	"emailservice/internal/domain/subscriber/age"
	"emailservice/internal/domain/subscriber/email"
	"emailservice/internal/domain/subscriber/name"
	"emailservice/internal/domain/subscriber/surname"
	"emailservice/internal/repository/postgres/subscriber/dao"

	"github.com/jackc/pgx/v5"
)

func (r Repository) toCopyFromSource(contacts ...*subscriber.Subscriber) pgx.CopyFromSource {
	rows := make([][]interface{}, len(contacts))

	for i, val := range contacts {
		rows[i] = []interface{}{
			val.ID(),
			val.CreatedAt(),
			val.ModifiedAt(),
			val.Name().String(),
			val.Surname().String(),
			val.Email(),
			val.Age(),
		}
	}
	return pgx.CopyFromRows(rows)
}

func (r Repository) toDomainSubscriber(dao *dao.Subscriber) (result *subscriber.Subscriber, err error) {
	name, err := name.NewName(dao.Name)
	if err != nil {
		return
	}

	surname, err := surname.NewSurname(dao.Surname)
	if err != nil {
		return
	}

	email, err := email.NewEmail(dao.Email)
	if err != nil {
		return
	}

	age, err := age.NewAge(dao.Age)
	if err != nil {
		return
	}

	result, err = subscriber.NewWithID(
		dao.ID,
		dao.CreatedAt,
		dao.ModifiedAt,
		*name,
		*surname,
		*email,
		*age,
	)
	return

}
