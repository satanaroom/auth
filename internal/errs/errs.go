package errs

import "errors"

var (
	ErrPasswordMismatch = errors.New("password and password confirm do not match")
	ErrUserToResponse   = errors.New("converting user to response error")
	ErrUserConverting   = errors.New("converting error")
	ErrUserNotFound     = errors.New("user not found")
)
