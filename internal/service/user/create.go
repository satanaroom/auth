package user

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/utils"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	if !isValidPassword(info.User.Password, info.PasswordConfirm) {
		logger.Errorf("password is invalid: %s", errs.ErrPasswordMismatch.Error())
		return 0, errs.ErrPasswordMismatch
	}

	passwordHash, err := utils.GeneratePasswordHash(info.User.Password)
	if err != nil {
		logger.Errorf("generate password hash: %s", err.Error())
		return 0, fmt.Errorf("generate password hash: %w", err)
	}
	info.User.Password = passwordHash

	id, err := s.userRepository.Create(ctx, info)
	if err != nil {
		logger.Errorf("authRepository.Create: %s", err.Error())
		return 0, fmt.Errorf("authRepository.Create: %w", err)
	}

	return id, nil
}

func isValidPassword(password, confirm string) bool {
	return password == confirm
}
