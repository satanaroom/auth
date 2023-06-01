package errs

import "errors"

var (
	ErrUserNotFound                     = errors.New("user not found")
	ErrMetadataIsNotProvided            = errors.New("metadata is not provided")
	ErrAuthorizationHeaderIsNotProvided = errors.New("authorization header is not provided")
	ErrAuthorizationHeaderFormat        = errors.New("authorization header format invalid")
	ErrAccessDenied                     = errors.New("access denied")
	ErrInvalidToken                     = errors.New("token isn't valid")
	ErrNoUserByUsername                 = errors.New("no such user")
	ErrGenerateToken                    = errors.New("generate token unavailable")
	ErrInvalidPassword                  = errors.New("password isn't valid")
	ErrPasswordHashGenerating           = errors.New("password hash generating unavailable")
)
