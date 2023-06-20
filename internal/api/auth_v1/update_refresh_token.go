package auth_v1

import (
	"context"

	desc "github.com/satanaroom/auth/pkg/auth_v1"
)

func (i *Implementation) UpdateRefreshToken(ctx context.Context, req *desc.UpdateRefreshTokenRequest) (*desc.UpdateRefreshTokenResponse, error) {
	newToken, err := i.authService.UpdateRefreshToken(ctx, req.GetOldToken())
	if err != nil {
		return nil, err
	}

	return &desc.UpdateRefreshTokenResponse{
		NewToken: newToken,
	}, nil
}
