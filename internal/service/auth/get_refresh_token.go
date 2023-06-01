package auth

import (
	"context"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
	"github.com/satanaroom/auth/internal/utils"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) GetRefreshToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepository.Get(ctx, username)
	if err != nil {
		logger.Errorf("userRepository.Get: %s", err.Error())
		return "", sys.NewCommonError(errs.ErrNoUserByUsername.Error(), codes.InvalidArgument)
	}

	if !utils.HashPassword(user.User.Password, password) {
		logger.Error("passwords isn't valid")
		return "", sys.NewCommonError(errs.ErrInvalidPassword.Error(), codes.InvalidArgument)
	}

	token, err := utils.GenerateToken(&user.User, s.config.RefreshTokenSecretKey(), s.config.RefreshTokenExpiration())
	if err != nil {
		logger.Errorf("utils.GenerateToken: %s", err.Error())
		return "", sys.NewCommonError("failed to generate token", codes.Internal)
	}

	return token, nil
}
