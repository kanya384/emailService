package subscriber

import (
	"context"
	"emailservice/internal/domain/subscriber"
	"emailservice/internal/repository/postgres/subscriber/dao"
	"errors"

	"github.com/google/uuid"

	"emailservice/pkg/tools/transaction"

	"github.com/jackc/pgx/v5"
)

const (
	tableName = `subscribers`
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"subscriber_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrUpdate       = errors.New("error updating or no changes")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateSubscriber(ctx context.Context, subscriber *subscriber.Subscriber) (err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Insert(tableName).Columns(dao.ColumnsSubscriber...).Values(subscriber.ID(), subscriber.Name(), subscriber.Surname(), subscriber.Email(), subscriber.Age(), subscriber.CreatedAt(), subscriber.ModifiedAt())
	query, args, _ := rawQuery.ToSql()

	_, err = tx.Exec(ctx, query, args...)
	return
}

func (r *Repository) CreateSubscriberTx(ctx context.Context, tx pgx.Tx, subscribers ...*subscriber.Subscriber) ([]*subscriber.Subscriber, error) {
	if len(subscribers) == 0 {
		return []*subscriber.Subscriber{}, nil
	}

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{tableName},
		dao.ColumnsSubscriber,
		r.toCopyFromSource(subscribers...))
	if err != nil {
		return nil, err
	}

	return subscribers, nil
}

func (r *Repository) UpdateSubscriber(ctx context.Context, ID uuid.UUID, updateFn func(subscriber *subscriber.Subscriber) (*subscriber.Subscriber, error)) (subscriber *subscriber.Subscriber, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	upSubscriber, err := r.oneSubscriberTX(ctx, tx, ID)
	if err != nil {
		return
	}

	subscriber, err = updateFn(upSubscriber)
	if err != nil {
		return
	}

	rawQuery := r.Builder.Update(tableName).Set("name", subscriber.Name()).Set("surname", subscriber.Surname()).Set("email", subscriber.Email()).Set("age", subscriber.Age()).Set("modified_at", subscriber.ModifiedAt()).Where("id = ?", subscriber.ID())
	query, args, _ := rawQuery.ToSql()

	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	if res.RowsAffected() == 0 {
		return nil, ErrUpdate
	}

	return
}

func (r *Repository) DeleteSubscriber(ctx context.Context, ID uuid.UUID) (err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Delete(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	res, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	if res.RowsAffected() == 0 {
		return ErrNotFound
	}

	return
}

func (r *Repository) ReadeSubscriberById(ctx context.Context, ID uuid.UUID) (subscriber *subscriber.Subscriber, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	subscriber, err = r.oneSubscriberTX(ctx, tx, ID)

	if subscriber == nil {
		err = ErrNotFound
	}

	return
}

func (r *Repository) oneSubscriberTX(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (response *subscriber.Subscriber, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsSubscriber...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoSubscriber, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Subscriber])
	if err != nil {
		return nil, err
	}

	return r.toDomainSubscriber(&daoSubscriber)

}
