package auth

import (
	"github.com/satanaroom/auth/internal/model"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
)

func ToInfo(info *desc.UserInfo, passwordConfirm string) *model.UserInfo {
	return &model.UserInfo{
		Username:        info.GetUsername(),
		Email:           info.GetEmail(),
		Password:        info.GetPassword(),
		PasswordConfirm: passwordConfirm,
		Role:            int(info.GetRole().Number()),
	}
}
func ToUsername(username string) model.Username {
	return model.Username(username)
}

func ToUser(info *desc.UserInfo) *model.User {
	return &model.User{
		Username: info.GetUsername(),
		Email:    info.GetEmail(),
		Password: info.GetPassword(),
		Role:     int(info.GetRole().Number()),
	}
}
