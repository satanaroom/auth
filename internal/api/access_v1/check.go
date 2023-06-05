package auth_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/satanaroom/auth/pkg/access_v1"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*empty.Empty, error) {
	if err := i.accessService.Check(ctx, req.GetEndpointAddress()); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
