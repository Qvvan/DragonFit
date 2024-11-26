package repository

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	errors "github.com/qvvan/dragonfit/pkg/client/postgresql/utils"
)

type BaseRepo struct {
	db    *pgxpool.Pool
	table string
}

// Метод создания записи
func (r *BaseRepo) create(ctx *gin.Context, newModel interface{}) (string, *errors.CustomError) {
	query, _, err := goqu.Insert(r.table).Rows(newModel).Returning("id").ToSQL()
	if err != nil {
		return "", errors.NewQueryError(r.table, "query build", err)
	}

	var id string
	if err := r.db.QueryRow(ctx, query).Scan(&id); err != nil {
		parsedErr := errors.ParsePostgresError(err)
		return "", errors.NewCreateError(r.table, "scan result", parsedErr)
	}

	return id, nil
}

// Метод получения SQL-запроса для записи по ID
func (r *BaseRepo) getQuery(id string) (string, *errors.CustomError) {
	qb := goqu.From(r.table).Where(goqu.I("id").Eq(id))
	query, _, err := qb.ToSQL()
	if err != nil {
		return "", errors.NewQueryError(r.table, "query build for get", err)
	}
	return query, nil
}

// Метод удаления записи
func (r *BaseRepo) Delete(ctx context.Context, id string) *errors.CustomError {
	ds := goqu.From(r.table).Delete().Where(goqu.I("id").Eq(id))

	query, _, err := ds.ToSQL()
	if err != nil {
		return errors.NewQueryError(r.table, "query build for delete", err)
	}

	if _, err := r.db.Exec(ctx, query); err != nil {
		parsedErr := errors.ParsePostgresError(err)
		return errors.NewDeleteError(r.table, id, parsedErr)
	}

	return nil
}

// Метод обновления записи
func (r *BaseRepo) update(ctx context.Context, updateModel interface{}, id string) *errors.CustomError {
	query, _, err := goqu.Update(r.table).
		Set(updateModel).
		Returning("id").
		Where(goqu.I("id").Eq(id)).
		ToSQL()

	if err != nil {
		return errors.NewQueryError(r.table, "query build for update", err)
	}

	if _, err := r.db.Exec(ctx, query); err != nil {
		parsedErr := errors.ParsePostgresError(err)
		return errors.NewUpdateError(r.table, id, parsedErr)
	}

	return nil
}
