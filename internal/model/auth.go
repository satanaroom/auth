package model

import "time"

type Role int64

const (
	RoleAdmin Role = 1
	RoleUser  Role = 2
)

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
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Username string
