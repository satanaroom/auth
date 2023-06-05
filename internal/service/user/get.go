package user

import (
	"context"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
)

func (s *service) Get(ctx context.Context, username model.Username) (*model.UserService, error) {
	user, err := s.userRepository.Get(ctx, string(username))
	if err != nil {
		return nil, sys.NewCommonError("failed to get user", codes.Internal)
	}

	return user, nil
}
