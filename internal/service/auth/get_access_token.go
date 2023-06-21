package auth

import (
	"context"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
	"github.com/satanaroom/auth/internal/utils"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, s.config.RefreshTokenSecretKey())
	if err != nil {
		return "", sys.NewCommonError(errs.ErrInvalidToken.Error(), codes.InvalidArgument)
	}

	user, err := s.userRepository.Get(ctx, claims.Username)
	if err != nil {
		return "", sys.NewCommonError(errs.ErrNoUserByUsername.Error(), codes.InvalidArgument)
	}

	accessToken, err := utils.GenerateToken(&user.User, s.config.AccessTokenSecretKey(), s.config.AccessTokenExpiration())
	if err != nil {
		return "", sys.NewCommonError(errs.ErrGenerateToken.Error(), codes.Internal)
	}

	return accessToken, nil
}
