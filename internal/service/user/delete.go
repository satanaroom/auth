package user

import (
	"context"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
)

func (s *service) Delete(ctx context.Context, username model.Username) (int64, error) {
	id, err := s.userRepository.Delete(ctx, string(username))
	if err != nil {
		return 0, sys.NewCommonError("failed to delete user", codes.Internal)
	}

	return id, nil
}
