package auth

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/model"
)

func (s *service) Get(ctx context.Context, username model.Username) (*model.User, error) {
	user, err := s.authRepository.Get(ctx, string(username))
	if err != nil {
		return nil, fmt.Errorf("authRepository.Get: %w", err)
	}

	return user, nil
}
