package config

import (
	"os"

	"github.com/satanaroom/auth/internal/errs"
)

var _ HTTPConfig = (*httpConfig)(nil)

const httpPortEnvName = "HTTP_PORT"

type HTTPConfig interface {
	Port() string
}

type httpConfig struct {
	port string
}

func NewHTTPConfig() (*httpConfig, error) {
	port := os.Getenv(httpPortEnvName)
	if port == "" {
		return nil, errs.ErrHTTPPortNotFound
	}

	return &httpConfig{
		port: port,
	}, nil
}

func (c *httpConfig) Port() string {
	return c.port
}
