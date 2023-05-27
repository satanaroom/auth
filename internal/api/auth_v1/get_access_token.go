package auth_v1

import (
	"context"

	desc "github.com/satanaroom/auth/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	accessToken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get access token: %s", err.Error())
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
