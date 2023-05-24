package auth_v1

import (
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
