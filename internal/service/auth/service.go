package auth

import (
	"context"

	"github.com/satanaroom/auth/internal/repository/user"
)

var _ Service = (*service)(nil)

type Service interface {
	GetRefreshToken(ctx context.Context, username, password string) (string, error)
	GetAccessToken() (string, error)
}

type service struct {
	userRepository user.Repository
}

func NewService(userRepository user.Repository) *service {
	return &service{
		userRepository: userRepository,
	}
}
