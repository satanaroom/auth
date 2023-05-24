package app

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	accessV1 "github.com/satanaroom/auth/internal/api/access_v1"
	authV1 "github.com/satanaroom/auth/internal/api/auth_v1"
	userV1 "github.com/satanaroom/auth/internal/api/user_v1"

	"github.com/satanaroom/auth/internal/client/pg"
	"github.com/satanaroom/auth/internal/closer"
	"github.com/satanaroom/auth/internal/config"
	accessRepository "github.com/satanaroom/auth/internal/repository/access"
	userRepository "github.com/satanaroom/auth/internal/repository/user"
	accessService "github.com/satanaroom/auth/internal/service/access"
	authService "github.com/satanaroom/auth/internal/service/auth"
	userService "github.com/satanaroom/auth/internal/service/user"

	"github.com/satanaroom/auth/pkg/logger"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	authConfig    config.AuthConfig

	pgClient         pg.Client
	userRepository   userRepository.Repository
	accessRepository accessRepository.Repository

	userService   userService.Service
	authService   authService.Service
	accessService accessService.Service

	userImpl   *userV1.Implementation
	authImpl   *authV1.Implementation
	accessImpl *accessV1.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			logger.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			logger.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			logger.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSwaggerConfig()
		if err != nil {
			logger.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) AuthConfig() config.AuthConfig {
	if s.authConfig == nil {
		cfg, err := config.NewAuthConfig()
		if err != nil {
			logger.Fatalf("failed to get auth config: %s", err.Error())
		}

		s.authConfig = cfg
	}

	return s.authConfig
}

func (s *serviceProvider) PGClient(ctx context.Context) pg.Client {
	if s.pgClient == nil {
		pgCfg, err := pgxpool.ParseConfig(s.PGConfig().DSN())
		if err != nil {
			logger.Fatalf("failed to get db config: %s", err.Error())
		}

		client, err := pg.NewClient(ctx, pgCfg)
		if err != nil {
			logger.Fatalf("failed to initialize pg client: %s", err.Error())
		}

		if err = client.PG().Ping(ctx); err != nil {
			logger.Fatalf("failed to ping pg: %s", err.Error())
		}

		closer.Add(client.Close)
		s.pgClient = client
	}
	return s.pgClient
}

func (s *serviceProvider) UserRepository(ctx context.Context) userRepository.Repository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.PGClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) AccessRepository(ctx context.Context) accessRepository.Repository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepository(s.PGClient(ctx))
	}

	return s.accessRepository
}

func (s *serviceProvider) UserService(ctx context.Context) userService.Service {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) authService.Service {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthConfig(), s.UserRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AccessService(ctx context.Context) accessService.Service {
	if s.accessService == nil {
		s.accessService = accessService.NewService(s.AuthConfig(), s.AccessRepository(ctx))
	}

	return s.accessService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *userV1.Implementation {
	if s.userImpl == nil {
		s.userImpl = userV1.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *authV1.Implementation {
	if s.authImpl == nil {
		s.authImpl = authV1.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) AccessImpl(ctx context.Context) *accessV1.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = accessV1.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}
