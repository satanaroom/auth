package auth_v1

import (
	"context"
	"fmt"

	desc "github.com/satanaroom/auth/pkg/auth_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	refreshToken, err := i.authService.GetRefreshToken(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, fmt.Errorf("authService.GetRefreshToken: %w", err)
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}
