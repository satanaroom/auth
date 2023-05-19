package user_v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	converter "github.com/satanaroom/auth/internal/converter/user"
	desc "github.com/satanaroom/auth/pkg/user_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	affectedRows, err := i.authService.Delete(ctx, converter.ToUsername(req.GetUsername()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %s", err.Error())
	}

	return &desc.DeleteResponse{
		AffectedRows: affectedRows,
	}, nil
}
