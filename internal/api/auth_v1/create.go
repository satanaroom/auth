package auth_v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	converter "github.com/satanaroom/auth/internal/converter/auth"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/auth/pkg/errs"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if err := validateCreateRequest(req.GetInfo(), req.GetPasswordConfirm()); err != nil {
		i.log.Errorf("validate request: %s", err.Error())
		return &desc.CreateResponse{}, status.Errorf(codes.InvalidArgument, "validate request: %s", err.Error())
	}

	id, err := i.authService.Create(ctx, converter.ToInfo(req.GetInfo(), req.GetPasswordConfirm()))
	if err != nil {
		i.log.Errorf("failed to create user: %s", err.Error())
		return &desc.CreateResponse{}, status.Errorf(codes.Internal, "failed to create user: %s", err.Error())
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func validateCreateRequest(info *desc.UserInfo, passwordConfirm string) error {
	if info.GetUsername() == "" {
		return errs.ErrUsernameEmpty
	}
	if int(info.GetRole().Number()) == 0 {
		return errs.ErrRoleEmpty
	}
	if info.GetPassword() == "" || passwordConfirm == "" {
		return errs.ErrPasswordEmpty
	}

	return nil
}
