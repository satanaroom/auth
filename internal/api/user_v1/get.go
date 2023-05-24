package user_v1

import (
	"context"

	converter "github.com/satanaroom/auth/internal/converter/user"
	desc "github.com/satanaroom/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, converter.ToUsername(req.GetUsername()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %s", err.Error())
	}

	return converter.ToGetDesc(user), nil
}
