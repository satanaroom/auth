package model

import (
	"database/sql"
	"time"
)

type Role int64

const (
	RoleAdmin Role = 1
	RoleUser  Role = 2
)

type UpdateUser struct {
	Username sql.NullString `db:"username"`
	Email    sql.NullString `db:"email"`
	Password sql.NullString `db:"password"`
	Role     sql.NullInt32  `db:"role"`
}

type User struct {
	Username string
	Email    string
	Password string
	Role     Role
}

type UserInfo struct {
	User            User
	PasswordConfirm string
}

type UserRepo struct {
	User      UpdateUser   `db:""`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserService struct {
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Username string

type UserId struct {
	Id int64 `db:"id"`
}
