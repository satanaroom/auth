package auth

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/utils"
)

func (s *service) GetRefreshToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepository.Get(ctx, username)
	if err != nil {
		return "", fmt.Errorf("userRepository.Get: %w", err)
	}

	if !utils.HashPassword(user.User.Password, password) {
		return "", fmt.Errorf("password is invalid")
	}

	token, err := utils.GenerateToken(&user.User, s.config.RefreshTokenSecretKey(), s.config.RefreshTokenExpiration())
	if err != nil {
		return "", fmt.Errorf("utils.GenerateToken: %w", err)
	}

	return token, nil
}
