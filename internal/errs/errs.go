package errs

import "errors"

var (
	ErrPasswordMismatch                 = errors.New("password and password confirm do not match")
	ErrUserToResponse                   = errors.New("converting user to response error")
	ErrUserConverting                   = errors.New("converting error")
	ErrUserNotFound                     = errors.New("user not found")
	ErrMetadataIsNotProvided            = errors.New("metadata is not provided")
	ErrAuthorizationHeaderIsNotProvided = errors.New("authorization header is not provided")
	ErrAuthorizationHeaderFormat        = errors.New("authorization header format invalid")
	ErrAccessDenied                     = errors.New("access denied")
)
