package auth

import (
	"context"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
	"github.com/satanaroom/auth/internal/utils"
)

func (s *service) UpdateRefreshToken(_ context.Context, oldToken string) (string, error) {
	claims, err := utils.VerifyToken(oldToken, s.config.RefreshTokenSecretKey())
	if err != nil {
		return "", sys.NewCommonError(errs.ErrInvalidToken.Error(), codes.InvalidArgument)
	}

	user := &model.User{
		Username: claims.Username,
		Role:     claims.Role,
	}

	newToken, err := utils.GenerateToken(user, s.config.RefreshTokenSecretKey(), s.config.RefreshTokenExpiration())
	if err != nil {
		return "", sys.NewCommonError(errs.ErrGenerateToken.Error(), codes.Internal)
	}

	return newToken, nil
}
