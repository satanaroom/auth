package auth

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Update(ctx context.Context, username string, user *model.UpdateUser) (int64, error) {
	if !user.Username.Valid {
		logger.Errorf("username is required: %s", errs.ErrUsernameRequired.Error())
		return 0, errs.ErrUsernameRequired
	}

	if !user.Role.Valid {
		logger.Errorf("role is required: %s", errs.ErrRoleRequired.Error())
		return 0, errs.ErrRoleRequired
	}

	if user.Role.Int32 != 0 && !isValidRole(model.Role(user.Role.Int32)) {
		logger.Errorf("role is invalid: %s", errs.ErrRoleInvalid.Error())
		return 0, errs.ErrRoleInvalid
	}

	if user.Email.Valid {
		if user.Email.String != "" && !isValidEmail(user.Email.String) {
			logger.Errorf("email is invalid: %s", errs.ErrEmailInvalid.Error())
			return 0, errs.ErrEmailInvalid
		}
	}

	id, err := s.authRepository.Update(ctx, username, user)
	if err != nil {
		logger.Errorf("authRepository.Update: %s", err.Error())
		return 0, fmt.Errorf("authRepository.Update: %w", err)
	}

	return id, nil
}
