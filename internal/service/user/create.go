package user

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	if !isValidPassword(info.User.Password, info.PasswordConfirm) {
		logger.Errorf("password is invalid: %s", errs.ErrPasswordMismatch.Error())
		return 0, errs.ErrPasswordMismatch
	}

	id, err := s.authRepository.Create(ctx, info)
	if err != nil {
		logger.Errorf("authRepository.Create: %s", err.Error())
		return 0, fmt.Errorf("authRepository.Create: %w", err)
	}

	return id, nil
}

func isValidPassword(password, confirm string) bool {
	return password == confirm
}
