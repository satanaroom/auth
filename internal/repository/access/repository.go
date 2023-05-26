package access

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/satanaroom/auth/internal/client/pg"
	"github.com/satanaroom/auth/internal/model"
)

var _ Repository = (*repository)(nil)

const tableName = "accesses"

type Repository interface {
	GetList(ctx context.Context) ([]model.AccessInfo, error)
}

type repository struct {
	pgClient pg.Client
}

func NewRepository(pgClient pg.Client) *repository {
	return &repository{
		pgClient: pgClient,
	}
}

func (r *repository) GetList(ctx context.Context) ([]model.AccessInfo, error) {
	builder := sq.Select("id", "endpoint_address", "role", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql: %w", err)
	}

	q := pg.Query{
		Name:     "access.GetList",
		QueryRaw: query,
	}

	var accessInfo []model.AccessInfo
	if err = r.pgClient.PG().ScanAll(ctx, &accessInfo, q, v...); err != nil {
		return nil, fmt.Errorf("scan one: %w", err)
	}

	return accessInfo, nil
}
