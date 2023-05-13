package errs

import "errors"

var (
	ErrUsernameEmpty       = errors.New("username is empty")
	ErrPasswordEmpty       = errors.New("password is empty")
	ErrPasswordMismatch    = errors.New("password and password confirm do not match")
	ErrRoleInvalid         = errors.New("role is invalid")
	ErrEmailInvalid        = errors.New("email is invalid")
	ErrUserNotFound        = errors.New("user not found")
	ErrDSNNotFound         = errors.New("pg dsn not found")
	ErrGRPCPortNotFound    = errors.New("grpc port not found")
	ErrHTTPPortNotFound    = errors.New("http port not found")
	ErrSwaggerPortNotFound = errors.New("swagger port not found")
)
