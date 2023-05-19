package user_v1

import (
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
