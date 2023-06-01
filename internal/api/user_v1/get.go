package user_v1

import (
	"context"

	converter "github.com/satanaroom/auth/internal/converter/user"
	desc "github.com/satanaroom/auth/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, converter.ToUsername(req.GetUsername()))
	if err != nil {
		return nil, err
	}

	return converter.ToGetDesc(user), nil
}
