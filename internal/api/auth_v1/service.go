package auth_v1

import (
	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/service/auth"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server

	authService auth.Service
}

func NewImplementation(authService auth.Service) *Implementation {
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
