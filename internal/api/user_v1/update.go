package user_v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	converter "github.com/satanaroom/auth/internal/converter/user"
	desc "github.com/satanaroom/auth/pkg/user_v1"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	if err := validateUsernameRequest(req.GetUsername()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate request: %s", err.Error())
	}

	info, err := converter.ToUpdateUser(req.GetInfo())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "converting: %s", err.Error())
	}
	id, err := i.authService.Update(ctx, req.GetUsername(), info)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err.Error())
	}

	return &desc.UpdateResponse{
		Id: id,
	}, nil
}
