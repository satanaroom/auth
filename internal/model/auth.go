package model

import "time"

const (
	RoleAdmin = iota + 1
	RoleUser
)

type UserInfo struct {
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
	Role            int
}

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Email     string
	Password  string
	Role      int
}

type Username string
