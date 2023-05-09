package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/satanaroom/auth/internal/client/pg"
	converter "github.com/satanaroom/auth/internal/converter/auth"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
)

var _ Repository = (*repository)(nil)

const tableName = "users"

type Repository interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, username string) (*model.UserService, error)
	Update(ctx context.Context, username string, user *model.UpdateUser) (int64, error)
	Delete(ctx context.Context, username string) (int64, error)
}

type repository struct {
	pgClient pg.Client
}

func NewRepository(pgClient pg.Client) *repository {
	return &repository{
		pgClient: pgClient,
	}
}

func (r *repository) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	t := time.Now().UTC()

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("username", "email", "password", "role", "created_at", "updated_at").
		Values(info.User.Username, info.User.Email, info.User.Password, info.User.Role, t, t).
		Suffix("RETURNING id")

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	q := pg.Query{
		Name:     "auth.Create",
		QueryRaw: query,
	}

	var id int64
	if err = r.pgClient.PG().QueryRowContext(ctx, q, v...).Scan(&id); err != nil {
		return 0, fmt.Errorf("query row ctx: %w", err)
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, username string) (*model.UserService, error) {
	builder := sq.Select("username", "email", "password", "role", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{
			"username": username,
		}).
		Limit(1)

	query, v, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("to sql: %w", err)
	}
	var (
		name      sql.NullString
		email     sql.NullString
		password  sql.NullString
		role      sql.NullInt32
		createdAt sql.NullTime
		updatedAt sql.NullTime
	)
	q := pg.Query{
		Name:     "auth.Get",
		QueryRaw: query,
	}

	if err = r.pgClient.PG().QueryRowContext(ctx, q, v...).Scan(&name, &email, &password, &role, &createdAt, &updatedAt); err != nil {
		return nil, fmt.Errorf("query row ctx: %w", err)
	}

	return converter.ToGetService(&model.UserRepo{
		User: model.UpdateUser{
			Username: name,
			Email:    email,
			Password: password,
			Role:     role,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}), nil
}

func (r *repository) Delete(ctx context.Context, username string) (int64, error) {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{
			"username": username,
		})

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	q := pg.Query{
		Name:     "auth.Delete",
		QueryRaw: query,
	}

	res, err := r.pgClient.PG().ExecContext(ctx, q, v...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errs.ErrUserNotFound
		}
		return 0, fmt.Errorf("query row ctx: %w", err)
	}

	return res.RowsAffected(), nil
}

func (r *repository) Update(ctx context.Context, username string, user *model.UpdateUser) (int64, error) {
	t := time.Now().UTC()

	selectBuilder := sq.Select("id").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{
			"username": username,
		}).
		Limit(1)

	query, v, err := selectBuilder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	q := pg.Query{
		Name:     "auth.UpdateGetId",
		QueryRaw: query,
	}

	var id int64
	if err = r.pgClient.PG().QueryRowContext(ctx, q, v...).Scan(&id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("query row ctx: %w", errs.ErrUserNotFound)
		}
		return 0, fmt.Errorf("query row ctx: %w", err)
	}

	updateBuilder := sq.Update("users").
		PlaceholderFormat(sq.Dollar)
	updateBuilder = updateBuilder.Set("username", user.Username)
	updateBuilder = updateBuilder.Set("email", user.Email)
	updateBuilder = updateBuilder.Set("password", user.Password)
	updateBuilder = updateBuilder.Set("role", user.Role)
	updateBuilder = updateBuilder.Set("updated_at", t).
		Where(sq.Eq{
			"username": username,
		})

	query, v, err = updateBuilder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	q = pg.Query{
		Name:     "auth.Update",
		QueryRaw: query,
	}

	if _, err = r.pgClient.PG().ExecContext(ctx, q, v...); err != nil {
		return 0, fmt.Errorf("exec ctx: %w", err)
	}

	return id, nil
}
