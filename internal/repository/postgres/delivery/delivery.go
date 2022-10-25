package delivery

import (
	"context"
	"emailservice/pkg/tools/transaction"
	"errors"
	"time"

	"emailservice/internal/domain/delivery"
	"emailservice/internal/repository/postgres/delivery/dao"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	tableName = `delivery`
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"delivery_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrUpdate       = errors.New("error updating or no changes")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateDelivery(ctx context.Context, delivery *delivery.Delivery) (err error) {
	tx, err := r.pg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.pg.Builder.Insert(tableName).Columns(dao.ColumnsDelivery...).Values(delivery.ID(), delivery.TemplateID(), delivery.SendAt(), delivery.IsSended(), delivery.CreatedAt(), delivery.ModifiedAt())
	query, args, _ := rawQuery.ToSql()

	_, err = tx.Exec(ctx, query, args...)
	return
}

func (r *Repository) createDeliveryTx(ctx context.Context, delivery *delivery.Delivery) (tx pgx.Tx, err error) {
	tx, err = r.pg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	rawQuery := r.pg.Builder.Insert(tableName).Columns(dao.ColumnsDelivery...).Values(delivery.ID(), delivery.TemplateID(), delivery.SendAt(), delivery.IsSended(), delivery.CreatedAt(), delivery.ModifiedAt())
	query, args, _ := rawQuery.ToSql()

	_, err = tx.Exec(ctx, query, args...)
	return
}

func (r *Repository) UpdateDelivery(ctx context.Context, ID uuid.UUID, updateFn func(delivery *delivery.Delivery) (*delivery.Delivery, error)) (delivery *delivery.Delivery, err error) {
	tx, err := r.pg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	upDelivery, err := r.oneDeliveryTX(ctx, tx, ID)
	if err != nil {
		return
	}

	delivery, err = updateFn(upDelivery)
	if err != nil {
		return
	}

	rawQuery := r.pg.Builder.Update(tableName).Set("template_id", delivery.TemplateID()).Set("send_at", delivery.SendAt()).Set("sended", delivery.IsSended()).Set("modified_at", delivery.ModifiedAt()).Where("id = ?", delivery.ID())
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

func (r *Repository) DeleteDelivery(ctx context.Context, ID uuid.UUID) (err error) {
	tx, err := r.pg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.pg.Builder.Delete(tableName).Where("id = ?", ID)
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

func (r *Repository) ReadDeliveryById(ctx context.Context, ID uuid.UUID) (delivery *delivery.Delivery, err error) {
	tx, err := r.pg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	delivery, err = r.oneDeliveryTX(ctx, tx, ID)

	if delivery == nil {
		err = ErrNotFound
	}

	return
}

func (r *Repository) ReadDeliveryTasks(ctx context.Context) (list []*delivery.Delivery, err error) {
	tx, err := r.pg.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.pg.Builder.Select(dao.ColumnsDelivery...).From(tableName).Where("sended = ? and send_at < ?", false, time.Now())
	query, args, _ := rawQuery.ToSql()

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return
	}

	for rows.Next() {
		delivery := dao.Delivery{}
		err = rows.Scan(&delivery.ID, &delivery.TemplateId, &delivery.SendAt, &delivery.Sended, &delivery.CreatedAt, &delivery.ModifiedAt)
		if err != nil {
			break
		}
		dDelivery, err := r.toDomainDelivery(&delivery)
		if err != nil {
			break
		}
		list = append(list, dDelivery)
	}

	if list == nil && err == nil {
		err = ErrNotFound
	}

	return
}

func (r *Repository) oneDeliveryTX(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (response *delivery.Delivery, err error) {
	rawQuery := r.pg.Builder.Select(dao.ColumnsDelivery...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoDelivery, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Delivery])
	if err != nil {
		return nil, err
	}

	return r.toDomainDelivery(&daoDelivery)

}
