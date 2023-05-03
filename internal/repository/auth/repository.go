package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
)

var _ Repository = (*repository)(nil)

const tableName = "users"

type Repository interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, username string) (*model.UserRepo, error)
	Update(ctx context.Context, username string, user *model.User) (int64, error)
	Delete(ctx context.Context, username string) (int64, error)
}

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
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

	var id int64
	if err = r.pool.QueryRow(ctx, query, v...).Scan(&id); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, username string) (*model.UserRepo, error) {
	builder := sq.Select("username", "email", "password", "role", "created_at", "updated_at").
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
	if err = r.pool.QueryRow(ctx, query, v...).Scan(&name, &email, &password, &role, &createdAt, &updatedAt); err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}

	return &model.UserRepo{
		User: model.User{
			Username: name.String,
			Email:    email.String,
			Password: password.String,
			Role:     int(role.Int32),
		},
		CreatedAt: createdAt.Time,
		UpdatedAt: updatedAt.Time,
	}, nil
}

func (r *repository) Delete(ctx context.Context, username string) (int64, error) {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{
			"username": username,
		}).
		Suffix("RETURNING id")

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	res, err := r.pool.Exec(ctx, query, v...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errs.ErrUserNotFound
		}
		return 0, fmt.Errorf("query row: %w", err)
	}

	return res.RowsAffected(), nil
}

func (r *repository) Update(ctx context.Context, username string, user *model.User) (int64, error) {
	t := time.Now().UTC()

	var id int64
	if err := r.pool.QueryRow(ctx, "SELECT id FROM users WHERE username = $1", username).Scan(&id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("query row: %w", errs.ErrUserNotFound)
		}
		return 0, fmt.Errorf("query row: %w", err)
	}

	builder := sq.Update("users").
		PlaceholderFormat(sq.Dollar)
	if user.Username != "" {
		builder = builder.Set("username", user.Username)
	}
	if user.Email != "" {
		builder = builder.Set("email", user.Email)
	}
	if user.Password != "" {
		builder = builder.Set("password", user.Password)
	}
	if user.Role != 0 {
		builder = builder.Set("role", user.Role)
	}
	builder = builder.Set("updated_at", t).
		Where(sq.Eq{
			"username": username,
		})

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	if _, err = r.pool.Exec(ctx, query, v...); err != nil {
		return 0, fmt.Errorf("exec: %w", err)
	}

	return id, nil
}
