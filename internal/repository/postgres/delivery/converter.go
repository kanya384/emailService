package delivery

import (
	"emailservice/internal/domain/delivery"
	"emailservice/internal/repository/postgres/delivery/dao"
	daoSub "emailservice/internal/repository/postgres/subscriber/dao"

	"github.com/google/uuid"

	"emailservice/internal/domain/subscriber"
	"emailservice/internal/domain/subscriber/age"
	"emailservice/internal/domain/subscriber/email"
	"emailservice/internal/domain/subscriber/name"
	"emailservice/internal/domain/subscriber/surname"
)

func (r Repository) toDomainDelivery(dao *dao.Delivery) (result *delivery.Delivery, err error) {
	id, err := uuid.Parse(dao.ID)
	if err != nil {
		return
	}
	templateID, err := uuid.Parse(dao.TemplateId)
	if err != nil {
		return
	}
	result, err = delivery.NewWithID(
		id,
		dao.CreatedAt,
		dao.ModifiedAt,
		templateID,
		dao.SendAt,
		dao.Sended,
	)
	return
}

func (r Repository) toDomainSubscriber(dao *daoSub.Subscriber) (result *subscriber.Subscriber, err error) {
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
