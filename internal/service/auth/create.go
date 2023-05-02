package auth

import (
	"context"
	"fmt"
	"regexp"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/pkg/errs"
)

func (s *service) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	if !isValidPassword(info.Password, info.PasswordConfirm) {
		return 0, errs.ErrPasswordMismatch
	}

	if !isValidRole(info.Role) {
		return 0, errs.ErrRoleInvalid
	}

	if info.Email != "" && !isValidEmail(info.Email) {
		return 0, errs.ErrEmailInvalid
	}

	id, err := s.authRepository.Create(ctx, info)
	if err != nil {
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
