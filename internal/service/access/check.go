package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/satanaroom/auth/internal/errs"
	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/internal/utils"
	"google.golang.org/grpc/metadata"
)

const authPrefix = "Bearer "

var rolesStorage map[string][]int

func (s *service) Check(ctx context.Context, endpointAddress string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errs.ErrMetadataIsNotProvided
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return errs.ErrAuthorizationHeaderIsNotProvided
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return errs.ErrAuthorizationHeaderFormat
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, s.config.AccessTokenSecretKey())
	if err != nil {
		return fmt.Errorf("utils.VerifyToken: %w", err)
	}

	accessibleRoles, err := s.accessibleRoles(ctx)
	if err != nil {
		return fmt.Errorf("accessibleRoles: %w", err)
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

	return errs.ErrAccessDenied
}

func (s *service) accessibleRoles(ctx context.Context) (map[string][]int, error) {
	if rolesStorage == nil {
		rolesStorage = make(map[string][]int)

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
