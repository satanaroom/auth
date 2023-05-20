package auth

import (
	"context"

	"github.com/satanaroom/auth/internal/config"
	"github.com/satanaroom/auth/internal/repository/user"
)

var _ Service = (*service)(nil)

type Service interface {
	GetRefreshToken(ctx context.Context, username, password string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}

type service struct {
	config config.AuthConfig

	userRepository user.Repository
}

func NewService(config config.AuthConfig, userRepository user.Repository) *service {
	return &service{
		config:         config,
		userRepository: userRepository,
	}
}
