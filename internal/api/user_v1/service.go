package user_v1

import (
	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/service/user"
	desc "github.com/satanaroom/auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server

	authService user.Service
}

func NewImplementation(authService user.Service) *Implementation {
	return &Implementation{
		authService: authService,
	}
}

func validateUsernameRequest(username string) error {
	if username == "" {
		return errs.ErrUsernameEmpty
	}

	return nil
}
