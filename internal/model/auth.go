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
	Username sql.NullString
	Email    sql.NullString
	Password sql.NullString
	Role     sql.NullInt32
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
	User      UpdateUser
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}

type UserService struct {
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Username string
