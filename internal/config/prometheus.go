package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ PrometheusConfig = (*prometheusConfig)(nil)

const prometheusHostEnvName = "PROMETHEUS_AUTH_HOST"

type PrometheusConfig interface {
	Host() string
}

type prometheusConfig struct {
	host string
}

func NewPrometheusConfig() (*prometheusConfig, error) {
	var host string
	env.ToString(&host, prometheusHostEnvName, "localhost:9090")

	return &prometheusConfig{
		host: host,
	}, nil
}

func (c *prometheusConfig) Host() string {
	return c.host
}
