package auth_v1

import (
	"context"

	desc "github.com/satanaroom/auth/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	refreshToken, err := i.authService.GetRefreshToken(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get refresh token: %s", err.Error())
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}
