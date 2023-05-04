package auth_v1

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	converter "github.com/satanaroom/auth/internal/converter/auth"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	if err := validateUsernameRequest(req.GetUsername()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate request: %s", err.Error())
	}

	affectedRows, err := i.authService.Delete(ctx, converter.ToUsername(req.GetUsername()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %s", err.Error())
	}

	return &desc.DeleteResponse{
		AffectedRows: affectedRows,
	}, nil
}
