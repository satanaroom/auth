package user

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Get(ctx context.Context, username model.Username) (*model.UserService, error) {
	user, err := s.authRepository.Get(ctx, string(username))
	if err != nil {
		logger.Errorf("authRepository.Get: %s", err.Error())
		return nil, fmt.Errorf("authRepository.Get: %w", err)
	}

	return user, nil
}
