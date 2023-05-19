package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/satanaroom/auth/internal/client/pg"
	converter "github.com/satanaroom/auth/internal/converter/user"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
)

var _ Repository = (*repository)(nil)

const tableName = "users"

type Repository interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, username string) (*model.UserService, error)
	Update(ctx context.Context, username string, user *model.UserRepo) (int64, error)
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
		Columns("username", "email", "password", "role",
			"created_at", "updated_at", "department", "department_type").
		Values(info.User.Username, info.User.Email, info.User.Password,
			info.User.Role, t, t, info.User.Department, info.User.DepartmentType).
		Suffix("RETURNING id")

	query, v, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	q := pg.Query{
		Name:     "user.Create",
		QueryRaw: query,
	}

	var uId model.UserId
	if err = r.pgClient.PG().ScanOne(ctx, &uId, q, v...); err != nil {
		return 0, fmt.Errorf("query row ctx: %w", err)
	}

	return uId.Id, nil
}

func (r *repository) Get(ctx context.Context, username string) (*model.UserService, error) {
	builder := sq.Select("username", "email", "password",
		"role", "created_at", "updated_at", "department", "department_type").
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

	q := pg.Query{
		Name:     "user.Get",
		QueryRaw: query,
	}

	var user model.GetUser
	if err = r.pgClient.PG().ScanOne(ctx, &user, q, v...); err != nil {
		return nil, fmt.Errorf("scan one: %w", err)
	}

	return converter.ToGetService(&user), nil
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
		Name:     "user.Delete",
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

func (r *repository) Update(ctx context.Context, username string, user *model.UserRepo) (int64, error) {
	t := time.Now().UTC()

	updateBuilder := sq.Update("users").
		PlaceholderFormat(sq.Dollar)

	if user.Username.Valid {
		updateBuilder = updateBuilder.Set("username", user.Username)
	}
	if user.Email.Valid {
		updateBuilder = updateBuilder.Set("email", user.Email)
	}
	if user.Password.Valid {
		updateBuilder = updateBuilder.Set("password", user.Password)
	}
	if user.Role.Valid {
		updateBuilder = updateBuilder.Set("role", user.Role)
	}
	if user.DepartmentType.Valid {
		updateBuilder = updateBuilder.Set("department", user.Department)
		updateBuilder = updateBuilder.Set("department_type", user.DepartmentType.Int32)
	}
	updateBuilder = updateBuilder.Set("updated_at", t).
		Where(sq.Eq{
			"username": username,
		}).
		Suffix("RETURNING id")

	query, v, err := updateBuilder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("to sql: %w", err)
	}

	q := pg.Query{
		Name:     "user.Update",
		QueryRaw: query,
	}

	var uId model.UserId
	if err = r.pgClient.PG().ScanOne(ctx, &uId, q, v...); err != nil {
		return 0, fmt.Errorf("exec ctx: %w", err)
	}

	return uId.Id, nil
}
