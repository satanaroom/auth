package model

import "time"

const (
	_Roles = iota
	RoleAdmin
	RoleUser
)

type User struct {
	Username string
	Email    string
	Password string
	Role     int
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
