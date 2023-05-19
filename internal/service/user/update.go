package user

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Update(ctx context.Context, username string, user *model.UserRepo) (int64, error) {
	if user == nil {
		logger.Errorf("user repo: %s", errs.ErrUserConverting.Error())
		return 0, fmt.Errorf("user repo: %w", errs.ErrUserConverting)
	}

	id, err := s.authRepository.Update(ctx, username, user)
	if err != nil {
		logger.Errorf("authRepository.Update: %s", err.Error())
		return 0, fmt.Errorf("authRepository.Update: %w", err)
	}

	return id, nil
}
