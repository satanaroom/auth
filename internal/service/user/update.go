package user

import (
	"context"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
)

func (s *service) Update(ctx context.Context, username string, user *model.UserRepo) (int64, error) {
	id, err := s.userRepository.Update(ctx, username, user)
	if err != nil {
		return 0, sys.NewCommonError("failed to update user", codes.Internal)
	}

	return id, nil
}
