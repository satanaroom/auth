package user_v1

import (
	"context"

	converter "github.com/satanaroom/auth/internal/converter/user"
	desc "github.com/satanaroom/auth/pkg/user_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	affectedRows, err := i.userService.Delete(ctx, converter.ToUsername(req.GetUsername()))
	if err != nil {
		return nil, err
	}

	return &desc.DeleteResponse{
		AffectedRows: affectedRows,
	}, nil
}
