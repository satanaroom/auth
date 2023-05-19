package user

import (
	"context"
	"fmt"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Delete(ctx context.Context, username model.Username) (int64, error) {
	id, err := s.authRepository.Delete(ctx, string(username))
	if err != nil {
		logger.Errorf("authRepository.Delete: %s", err.Error())
		return 0, fmt.Errorf("authRepository.Delete: %w", err)
	}

	return id, nil
}
