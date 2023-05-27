package user

import (
	"context"
	"fmt"
	"time"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
	"github.com/satanaroom/auth/internal/sys/validate"
	"github.com/satanaroom/auth/internal/utils"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	if err := validate.Validate(
		ctx,
		validatePassword(info.User.Password, info.PasswordConfirm),
	); err != nil {
		logger.Errorf("validate: %s", err.Error())
		return 0, err
	}

	passwordHash, err := utils.GeneratePasswordHash(info.User.Password)
	if err != nil {
		logger.Errorf("generate password hash: %s", err.Error())
		return 0, sys.NewCommonError("failed to generate password hash", codes.Internal)
	}
	info.User.Password = passwordHash

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, err := s.userRepository.Create(ctx, info)
	if err != nil {
		logger.Errorf("authRepository.Create: %s", err.Error())
		return 0, sys.NewCommonError(fmt.Sprintf("failed to create user: %s", err.Error()), codes.Internal)
	}

	return id, nil
}

func validatePassword(password, confirm string) validate.Condition {
	return func(ctx context.Context) error {
		if password != confirm {
			return validate.NewValidationErrors("password and password confirm must be equal")
		}

		return nil
	}
}
