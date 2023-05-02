package auth_v1

import (
	"github.com/sirupsen/logrus"
	
	"github.com/satanaroom/auth/internal/service/auth"
	desc "github.com/satanaroom/auth/pkg/auth_v1"
	"github.com/satanaroom/auth/pkg/errs"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server

	log         *logrus.Logger
	authService auth.Service
}

func NewImplementation(log *logrus.Logger, authService auth.Service) *Implementation {
	return &Implementation{
		log:         log,
		authService: authService,
	}
}

func validateUsernameRequest(username string) error {
	if username == "" {
		return errs.ErrUsernameEmpty
	}

	return nil
}
