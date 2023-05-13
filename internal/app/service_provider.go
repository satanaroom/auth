package app

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	authV1 "github.com/satanaroom/auth/internal/api/auth_v1"
	"github.com/satanaroom/auth/internal/client/pg"
	"github.com/satanaroom/auth/internal/closer"
	"github.com/satanaroom/auth/internal/config"
	authRepository "github.com/satanaroom/auth/internal/repository/auth"
	authService "github.com/satanaroom/auth/internal/service/auth"
	"github.com/satanaroom/auth/pkg/logger"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig

	pgClient       pg.Client
	authRepository authRepository.Repository
	authService    authService.Service

	authImpl *authV1.Implementation
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

func (s *serviceProvider) AuthRepository(ctx context.Context) authRepository.Repository {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepository(s.PGClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) authService.Service {
	if s.authService == nil {
		s.authService = authService.NewService(s.AuthRepository(ctx))
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *authV1.Implementation {
	if s.authImpl == nil {
		s.authImpl = authV1.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}
