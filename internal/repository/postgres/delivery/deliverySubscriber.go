package delivery

import (
	"context"
	"emailservice/internal/domain/delivery"

	"emailservice/internal/domain/subscriber"
	"emailservice/pkg/tools/transaction"

	"emailservice/internal/repository/postgres/subscriber/dao"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateDeliveryWithSubscribers(ctx context.Context, delivery *delivery.Delivery, subscribers ...*subscriber.Subscriber) (err error) {
	tx, err := r.createDeliveryTx(ctx, delivery)
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	response, err := r.repoSubscriber.CreateSubscriberTx(ctx, tx, subscribers...)
	if err != nil {
		return err
	}

	var subscribersIDs = make([]uuid.UUID, len(response))
	for i, c := range response {
		subscribersIDs[i] = c.ID()
	}

	if err = r.fillDeliverySubscribersTx(ctx, tx, delivery.ID(), subscribersIDs...); err != nil {
		return
	}

	return

}

func (r *Repository) MarkAsReadedBySubscriber(ctx context.Context, deliveryID uuid.UUID, subscriberID uuid.UUID) (err error) {

	tx, err := r.pg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.pg.Builder.Update("subscriber_in_delivery").Set("opened", true).Where("delivery_id = ? and subscriber_id = ?", deliveryID, subscriberID)
	query, args, _ := rawQuery.ToSql()

	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	if res.RowsAffected() == 0 {
		return ErrUpdate
	}

	return
}

func (r *Repository) ReadDeliverySubscribers(ctx context.Context, deliveryID uuid.UUID) (list []*subscriber.Subscriber, err error) {
	tx, err := r.pg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.pg.Builder.Select("subscriber.id", "subscriber.created_at", "subscriber.modified_at", "subscriber.name", "subscriber.surname", "subscriber.email", "subscriber.age").From("subscribers as subscriber").Join("subscriber_in_delivery as subInDel on subInDel.subscriber_id = subscriber.id").Where("subInDel.delivery_id = ?", deliveryID.String())
	query, args, _ := rawQuery.ToSql()

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return
	}

	for rows.Next() {
		subscriber := dao.Subscriber{}
		err = rows.Scan(&subscriber.ID, &subscriber.CreatedAt, &subscriber.ModifiedAt, &subscriber.Name, &subscriber.Surname, &subscriber.Email, &subscriber.Age)
		if err != nil {
			break
		}
		dSubscriber, err := r.toDomainSubscriber(&subscriber)
		if err != nil {
			break
		}
		list = append(list, dSubscriber)
	}

	if list == nil && err == nil {
		err = ErrNotFound
	}

	return
}

func (r *Repository) fillDeliverySubscribersTx(ctx context.Context, tx pgx.Tx, deliveryID uuid.UUID, subscribersIDs ...uuid.UUID) error {

	var rows [][]interface{}
	for _, subscriberID := range subscribersIDs {
		rows = append(rows, []interface{}{
			uuid.New(),
			deliveryID,
			subscriberID,
			false,
		})
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{"subscriber_in_delivery"},
		dao.CreateColumnSubscriberInDelivery,
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return err
	}

	return nil
}
