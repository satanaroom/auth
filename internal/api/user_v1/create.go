package user_v1

import (
	"context"

	converter "github.com/satanaroom/auth/internal/converter/user"
	desc "github.com/satanaroom/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToInfo(req.GetInfo(), req.GetPasswordConfirm()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
