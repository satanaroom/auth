package auth

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/model"
)

func (s *service) Delete(ctx context.Context, username model.Username) (int64, error) {
	id, err := s.authRepository.Delete(ctx, string(username))
	if err != nil {
		return 0, fmt.Errorf("authRepository.Delete: %w", err)
	}

	return id, nil
}
