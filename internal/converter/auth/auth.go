package auth

import (
	"database/sql"
	"time"

	"github.com/satanaroom/auth/internal/model"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToInfo(info *desc.UserInfo, passwordConfirm string) *model.UserInfo {
	return &model.UserInfo{
		User: model.User{
			Username: info.GetUsername(),
			Email:    info.GetEmail(),
			Password: info.GetPassword(),
			Role:     model.Role(info.GetRole().Number()),
		},
		PasswordConfirm: passwordConfirm,
	}
}
func ToUsername(username string) model.Username {
	return model.Username(username)
}

func ToUpdateUser(info *desc.UpdateUser) *model.UpdateUser {
	return &model.UpdateUser{
		Username: sql.NullString{
			String: info.Username.GetValue(),
			Valid:  info.Username.ProtoReflect().IsValid(),
		},
		Email: sql.NullString{
			String: info.Email.GetValue(),
			Valid:  info.Email.ProtoReflect().IsValid(),
		},
		Password: sql.NullString{
			String: info.Password.GetValue(),
			Valid:  info.Password.ProtoReflect().IsValid(),
		},
		Role: sql.NullInt32{
			Int32: info.Role.GetValue(),
			Valid: info.Role.ProtoReflect().IsValid(),
		},
	}
}

func ToGetService(user *model.UserRepo) *model.UserService {
	var (
		username, email, password string
		role                      model.Role
		createdAt                 time.Time
		updatedAt                 time.Time
	)

	if user.User.Username.Valid {
		username = user.User.Username.String
	}
	if user.User.Email.Valid {
		email = user.User.Email.String
	}
	if user.User.Password.Valid {
		password = user.User.Password.String
	}
	if user.User.Role.Valid {
		role = model.Role(user.User.Role.Int32)
	}
	if user.CreatedAt.Valid {
		createdAt = user.CreatedAt.Time
	}
	if user.UpdatedAt.Valid {
		updatedAt = user.UpdatedAt.Time
	}
	return &model.UserService{
		User: model.User{
			Username: username,
			Email:    email,
			Password: password,
			Role:     role,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func ToGetDesc(user *model.UserService) *desc.GetResponse {
	return &desc.GetResponse{
		Info: &desc.UserInfo{
			Username: user.User.Username,
			Email:    user.User.Email,
			Password: user.User.Password,
			Role:     desc.Role(user.User.Role),
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
