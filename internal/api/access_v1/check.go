package auth_v1

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/satanaroom/auth/pkg/access_v1"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*empty.Empty, error) {
	if err := i.accessService.Check(ctx, req.GetEndpointAddress()); err != nil {
		return nil, fmt.Errorf("accessService.Check: %w", err)
	}

	return &empty.Empty{}, nil
}
