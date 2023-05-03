package auth

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Update(ctx context.Context, username string, user *model.User) (int64, error) {
	if user.Role != 0 && !isValidRole(user.Role) {
		logger.Errorf("role is invalid: %s", errs.ErrRoleInvalid.Error())
		return 0, errs.ErrRoleInvalid
	}

	if user.Email != "" && !isValidEmail(user.Email) {
		logger.Errorf("email is invalid: %s", errs.ErrEmailInvalid.Error())
		return 0, errs.ErrEmailInvalid
	}

	id, err := s.authRepository.Update(ctx, username, user)
	if err != nil {
		logger.Errorf("authRepository.Update: %s", err.Error())
		return 0, fmt.Errorf("authRepository.Update: %w", err)
	}

	return id, nil
}
