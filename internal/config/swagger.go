package config

import (
	"os"

	"github.com/satanaroom/auth/internal/errs"
)

var _ SwaggerConfig = (*swaggerConfig)(nil)

const swaggerPortEnvName = "SWAGGER_PORT"

type SwaggerConfig interface {
	Port() string
}

type swaggerConfig struct {
	port string
}

func NewSwaggerConfig() (*swaggerConfig, error) {
	port := os.Getenv(swaggerPortEnvName)
	if port == "" {
		return nil, errs.ErrSwaggerPortNotFound
	}

	return &swaggerConfig{
		port: port,
	}, nil
}

func (c *swaggerConfig) Port() string {
	return c.port
}
