package auth_v1

import (
	"context"
	"fmt"

	desc "github.com/satanaroom/auth/pkg/access_v1"
)

func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*desc.CheckResponse, error) {
	accessToken, err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, fmt.Errorf("accessService.Check: %w", err)
	}

	return &desc.CheckResponse{
		IsAllowed: accessToken,
	}, nil
}
