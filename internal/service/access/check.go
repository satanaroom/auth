package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/sys"
	"github.com/satanaroom/auth/internal/sys/codes"
	"github.com/satanaroom/auth/internal/utils"
	"github.com/satanaroom/auth/pkg/logger"
	"google.golang.org/grpc/metadata"
)

const authPrefix = "Bearer "

var rolesStorage map[string][]int

func (s *service) Check(ctx context.Context, endpointAddress string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Error("metadata.FromIncomingContext: not provided")
		return sys.NewCommonError(errs.ErrMetadataIsNotProvided.Error(), codes.DataLoss)
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		logger.Error("authorization header: not provided")
		return sys.NewCommonError(errs.ErrAuthorizationHeaderIsNotProvided.Error(), codes.DataLoss)
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		logger.Error("authPrefix: not provided")
		return sys.NewCommonError(errs.ErrAuthorizationHeaderFormat.Error(), codes.DataLoss)
	}

	claims, err := utils.VerifyToken(strings.TrimPrefix(authHeader[0], authPrefix),
		s.config.AccessTokenSecretKey())
	if err != nil {
		logger.Errorf("utils.VerifyToken: %s", err.Error())
		return sys.NewCommonError(errs.ErrInvalidToken.Error(), codes.InvalidArgument)
	}

	accessibleRoles, err := s.accessibleRoles(ctx)
	if err != nil {
		logger.Errorf("accessibleRoles: %s", err.Error())
		return sys.NewCommonError(errs.ErrAccessDenied.Error(), codes.PermissionDenied)
	}

	role, ok := accessibleRoles[endpointAddress]
	if !ok {
		return nil
	}

	for _, r := range role {
		if claims.Role == model.Role(r) {
			return nil
		}
	}

	return sys.NewCommonError(errs.ErrAccessDenied.Error(), codes.PermissionDenied)
}

func (s *service) accessibleRoles(ctx context.Context) (map[string][]int, error) {
	if rolesStorage == nil {
		rolesStorage = make(map[string][]int)

		accessInfo, err := s.accessRepository.GetList(ctx)
		if err != nil {
			return nil, fmt.Errorf("get accesses list: %w", err)
		}

		for _, info := range accessInfo {
			rolesStorage[info.EndpointAddress] = info.Role
		}
	}

	return rolesStorage, nil
}
