package template

import (
	"context"
	"emailservice/internal/domain/template"
	"emailservice/internal/repository/postgres/template/dao"
	"emailservice/pkg/tools/transaction"
	"errors"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
)

const (
	tableName = `template`
)

var (
	ErrDuplicateKey = errors.New("ERROR: duplicate key value violates unique constraint \"template_pkey\" (SQLSTATE 23505)")
	ErrNotFound     = errors.New("not found")
	ErrUpdate       = errors.New("error updating or no changes")
	ErrEmptyResult  = errors.New("no rows in result set")
)

func (r *Repository) CreateTemplate(ctx context.Context, template *template.Template) (err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	rawQuery := r.Builder.Insert(tableName).Columns(dao.ColumnsTemplate...).Values(template.ID(), template.Path(), template.CreatedAt(), template.ModifiedAt())
	query, args, _ := rawQuery.ToSql()

	_, err = tx.Exec(ctx, query, args...)
	return
}

func (r *Repository) UpdateTemplate(ctx context.Context, ID uuid.UUID, updateFn func(template *template.Template) (*template.Template, error)) (template *template.Template, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	upTemplate, err := r.oneTemplateTX(ctx, tx, ID)
	if err != nil {
		return
	}

	template, err = updateFn(upTemplate)
	if err != nil {
		return
	}

	rawQuery := r.Builder.Update(tableName).Set("path", template.Path()).Set("modified_at", template.ModifiedAt()).Where("id = ?", template.ID())
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

func (r *Repository) DeleteTemplate(ctx context.Context, ID uuid.UUID) (err error) {
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

func (r *Repository) ReadTemplateById(ctx context.Context, ID uuid.UUID) (template *template.Template, err error) {
	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	template, err = r.oneTemplateTX(ctx, tx, ID)

	if template == nil {
		err = ErrNotFound
	}

	return
}

func (r *Repository) oneTemplateTX(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (response *template.Template, err error) {
	rawQuery := r.Builder.Select(dao.ColumnsTemplate...).From(tableName).Where("id = ?", ID)
	query, args, _ := rawQuery.ToSql()

	row, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	daoTemplate, err := pgx.CollectOneRow(row, pgx.RowToStructByPos[dao.Template])
	if err != nil {
		return nil, err
	}

	return r.toDomainTemplate(&daoTemplate)

}
