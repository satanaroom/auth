package config

import (
	"os"

	"github.com/satanaroom/auth/internal/errs"
)

var _ GRPCConfig = (*grpcConfig)(nil)

const grpcPortEnvName = "GRPC_PORT"

type GRPCConfig interface {
	Port() string
}

type grpcConfig struct {
	port string
}

func NewGRPCConfig() (*grpcConfig, error) {
	port := os.Getenv(grpcPortEnvName)
	if port == "" {
		return nil, errs.ErrGRPCPortNotFound
	}

	return &grpcConfig{
		port: port,
	}, nil
}

func (c *grpcConfig) Port() string {
	return c.port
}
