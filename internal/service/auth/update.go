package auth

import (
	"context"
	"fmt"
	"github.com/satanaroom/auth/pkg/errs"

	"github.com/satanaroom/auth/internal/model"
)

func (s *service) Update(ctx context.Context, username string, user *model.User) (int64, error) {
	if user.Role != 0 && !isValidRole(user.Role) {
		return 0, errs.ErrRoleInvalid
	}

	if user.Email != "" && !isValidEmail(user.Email) {
		return 0, errs.ErrEmailInvalid
	}

	id, err := s.authRepository.Update(ctx, username, user)
	if err != nil {
		return 0, fmt.Errorf("authRepository.Update: %w", err)
	}

	return id, nil
}
