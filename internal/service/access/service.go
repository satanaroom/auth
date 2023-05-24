package auth

import (
	"context"

	"github.com/satanaroom/auth/internal/config"
	"github.com/satanaroom/auth/internal/repository/access"
)

var _ Service = (*service)(nil)

type Service interface {
	Check(ctx context.Context, endpointAddress string) error
}

type service struct {
	config config.AuthConfig

	accessRepository access.Repository
}

func NewService(config config.AuthConfig, accessRepository access.Repository) *service {
	return &service{
		config:           config,
		accessRepository: accessRepository,
	}
}
