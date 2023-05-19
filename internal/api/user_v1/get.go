package user_v1

import (
	"context"

	converter "github.com/satanaroom/auth/internal/converter/user"
	"github.com/satanaroom/auth/internal/errs"
	desc "github.com/satanaroom/auth/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.authService.Get(ctx, converter.ToUsername(req.GetUsername()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %s", err.Error())
	}
	resp := converter.ToGetDesc(user)
	if resp == nil {
		return nil, status.Errorf(codes.Internal, "failed converting user to response: %s",
			errs.ErrUserToResponse.Error())
	}

	return resp, nil
}
