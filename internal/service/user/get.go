package user

import (
	"context"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
	"github.com/satanaroom/auth/pkg/logger"
)

func (s *service) Get(ctx context.Context, username model.Username) (*model.UserService, error) {
	user, err := s.userRepository.Get(ctx, string(username))
	if err != nil {
		logger.Errorf("authRepository.Get: %s", err.Error())
		return nil, sys.NewCommonError("failed to get user", codes.Internal)
	}

	return user, nil
}
