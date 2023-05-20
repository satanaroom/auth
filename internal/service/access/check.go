package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/utils"
	"google.golang.org/grpc/metadata"
)

const authPrefix = "Bearer "

var rolesStorage = map[string]string{}

func (s *service) Check(ctx context.Context, endpointAddress string) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, errs.ErrMetadataIsNotProvided
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return false, errs.ErrAuthorizationHeaderIsNotProvided
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return false, errs.ErrAuthorizationHeaderFormat
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, s.config.AccessTokenSecretKey())
	if err != nil {
		return false, fmt.Errorf("utils.VerifyToken: %w", err)
	}

	accessibleRoles, err := s.accessibleRoles(ctx)
	if err != nil {
		return false, fmt.Errorf("accessibleRoles: %w", err)
	}

	role, ok := accessibleRoles[endpointAddress]
	if !ok {
		return true, nil
	}

	if claims.Role != role {
		return false, errs.ErrAccessDenied
	}

	return true, nil
}

func (s *service) accessibleRoles(ctx context.Context) (map[string]string, error) {
	if rolesStorage == nil {
		rolesStorage = make(map[string]string)

		accessInfo, err := s.accessRepository.GetList(ctx)
		if err != nil {
			return nil, fmt.Errorf("accessRepository.GetList: %w", err)
		}

		for _, info := range accessInfo {
			rolesStorage[info.EndpointAddress] = info.Role
		}
	}

	return rolesStorage, nil
}
