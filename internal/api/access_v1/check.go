package auth_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/satanaroom/auth/pkg/access_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*empty.Empty, error) {
	if err := i.accessService.Check(ctx, req.GetEndpointAddress()); err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "access denied: %s", err.Error())
	}

	return &empty.Empty{}, nil
}
