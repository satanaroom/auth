package auth

import (
	"context"
	"fmt"
	"regexp"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	if !isValidPassword(info.User.Password, info.PasswordConfirm) {
		logger.Errorf("password is invalid: %s", errs.ErrPasswordMismatch.Error())
		return 0, errs.ErrPasswordMismatch
	}

	if !isValidRole(info.User.Role) {
		logger.Errorf("role is invalid: %s", errs.ErrRoleInvalid.Error())
		return 0, errs.ErrRoleInvalid
	}

	if info.User.Email != "" && !isValidEmail(info.User.Email) {
		logger.Errorf("email is invalid: %s", errs.ErrEmailInvalid.Error())
		return 0, errs.ErrEmailInvalid
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

func isValidRole(role int) bool {
	if role != model.RoleAdmin && role != model.RoleUser {
		return false
	}
	return true
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)

	return re.MatchString(email)
}
