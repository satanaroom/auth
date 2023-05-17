package model

import (
	"database/sql"
	"time"
)

type Role int64
type Rate int64
type Dept int64

const (
	RoleAdmin Role = 1
	RoleUser  Role = 2

	RateHalf Rate = 1
	RateFull Rate = 2

	DevelopmentDept Dept = 1
	AnalyticsDept   Dept = 2
)

type UserRepo struct {
	Username       sql.NullString `db:"username"`
	Email          sql.NullString `db:"email"`
	Password       sql.NullString `db:"password"`
	Role           sql.NullInt32  `db:"role"`
	Department     []byte         `db:"department"`
	DepartmentType sql.NullInt32  `db:"department_type"`
}

type Development struct {
	Grade    int32  `json:"grade"`
	Language string `json:"language"`
	Rate     Rate   `json:"rate"`
}

type Analytics struct {
	Specialization string `json:"specialization"`
	Rate           Rate   `json:"rate"`
}

type User struct {
	Username       string
	Email          string
	Password       string
	Role           Role
	Department     []byte
	DepartmentType Dept
}

type UserInfo struct {
	User            User
	PasswordConfirm string
}

type GetUser struct {
	User      UserRepo     `db:""`
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
