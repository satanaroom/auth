package user_v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	converter "github.com/satanaroom/auth/internal/converter/user"
	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	desc "github.com/satanaroom/auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if err := validateCreateRequest(req.GetInfo(), req.GetPasswordConfirm()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validate request: %s", err.Error())
	}

	id, err := i.authService.Create(ctx, converter.ToInfo(req.GetInfo(), req.GetPasswordConfirm()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err.Error())
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func validateCreateRequest(info *desc.UserInfo, passwordConfirm string) error {
	if info.GetUsername() == "" {
		return errs.ErrUsernameEmpty
	}
	role := model.Role(info.GetRole().Number())
	if role != model.RoleAdmin && role != model.RoleUser {
		return errs.ErrRoleInvalid
	}
	if info.GetPassword() == "" || passwordConfirm == "" {
		return errs.ErrPasswordEmpty
	}

	return nil
}
