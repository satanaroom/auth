package auth

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/utils"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, s.config.RefreshTokenSecretKey())
	if err != nil {
		return "", fmt.Errorf("utils.VerifyToken: %w", err)
	}

	user, err := s.userRepository.Get(ctx, claims.Username)
	if err != nil {
		return "", fmt.Errorf("userRepository.Get: %w", err)
	}

	accessToken, err := utils.GenerateToken(&user.User, s.config.AccessTokenSecretKey(), s.config.AccessTokenExpiration())
	if err != nil {
		return "", fmt.Errorf("utils.GenerateToken: %w", err)
	}

	return accessToken, nil
}
