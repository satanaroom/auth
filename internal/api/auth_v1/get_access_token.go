package auth_v1

import (
	"context"
	"fmt"

	desc "github.com/satanaroom/auth/pkg/auth_v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	accessToken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, fmt.Errorf("authService.GetAccessToken: %w", err)
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
