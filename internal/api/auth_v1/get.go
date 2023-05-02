package auth_v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	converter "github.com/satanaroom/auth/internal/converter/auth"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	if err := validateUsernameRequest(req.GetUsername()); err != nil {
		i.log.Errorf("validate request: %s", err.Error())
		return &desc.GetResponse{}, status.Errorf(codes.InvalidArgument, "validate request: %s", err.Error())
	}

	user, err := i.authService.Get(ctx, converter.ToUsername(req.GetUsername()))
	if err != nil {
		i.log.Errorf("failed to get user: %s", err.Error())
		return &desc.GetResponse{}, status.Errorf(codes.Internal, "failed to get user: %s", err.Error())
	}

	return &desc.GetResponse{
		Info: &desc.UserInfo{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
			Role:     desc.Role(user.Role),
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}
