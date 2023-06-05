package user_v1

import (
	"context"

	converter "github.com/satanaroom/auth/internal/converter/user"
	desc "github.com/satanaroom/auth/pkg/user_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	id, err := i.userService.Update(ctx, req.GetUsername(), converter.ToUpdateUser(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &desc.UpdateResponse{
		Id: id,
	}, nil
}
