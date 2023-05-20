package auth_v1

import (
	access "github.com/satanaroom/auth/internal/service/access"
	desc "github.com/satanaroom/auth/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server

	accessService access.Service
}

func NewImplementation(accessService access.Service) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
