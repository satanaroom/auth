package user

import (
	"context"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
	"github.com/satanaroom/auth/internal/sys/validate"
	"github.com/satanaroom/auth/internal/utils"
)

func (s *service) Create(ctx context.Context, info *model.UserInfo) (int64, error) {
	if err := validate.Validate(
		ctx,
		validatePassword(info.User.Password, info.PasswordConfirm),
	); err != nil {
		return 0, err
	}

	passwordHash, err := utils.GeneratePasswordHash(info.User.Password)
	if err != nil {
		return 0, sys.NewCommonError("failed to generate password hash", codes.Internal)
	}
	info.User.Password = passwordHash

	id, err := s.userRepository.Create(ctx, info)
	if err != nil {
		return 0, sys.NewCommonError("failed to create user", codes.Internal)
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
