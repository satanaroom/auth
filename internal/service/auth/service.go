package auth

import (
	"context"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/repository/auth"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(ctx context.Context, info *model.UserInfo) (int64, error)
	Get(ctx context.Context, username model.Username) (*model.UserService, error)
	Update(ctx context.Context, username string, user *model.UpdateUser) (int64, error)
	Delete(ctx context.Context, username model.Username) (int64, error)
}

type service struct {
	authRepository auth.Repository
}

func NewService(authRepository auth.Repository) *service {
	return &service{
		authRepository: authRepository,
	}
}
