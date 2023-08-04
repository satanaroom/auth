package user_v1

import (
	"github.com/satanaroom/auth/internal/service/user"
	desc "github.com/satanaroom/auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server

	userService user.Service
}

func NewImplementation(userService user.Service) *Implementation {
	return &Implementation{
		userService: userService,
	}
}

func newMockUserV1(i Implementation) *Implementation {
	return &Implementation{
		desc.UnimplementedUserV1Server{},
		i.userService,
	}
}
