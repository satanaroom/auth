package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ GRPCConfig = (*grpcConfig)(nil)

const grpcHostEnvName = "GRPC_AUTH_HOST"

type GRPCConfig interface {
	Host() string
}

type grpcConfig struct {
	host string
}

func NewGRPCConfig() (*grpcConfig, error) {
	var host string
	env.ToString(&host, grpcHostEnvName, "localhost:50051")

	return &grpcConfig{
		host: host,
	}, nil
}

func (c *grpcConfig) Host() string {
	return c.host
}
