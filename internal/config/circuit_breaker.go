package config

import (
	"time"

	"github.com/satanaroom/auth/pkg/env"
)

var _ CircuitBreakerConfig = (*circuitBreakerConfig)(nil)

const (
	circuitBreakerNameEnvName              = "CB_AUTH_NAME"
	circuitBreakerMaxRequestsEnvName       = "CB_AUTH_MAX_REQUESTS"
	circuitBreakerIntervalEnvName          = "CB_AUTH_INTERVAL"
	circuitBreakerTimeoutEnvName           = "CB_AUTH_TIMEOUT"
	circuitBreakerFailureRatioLimitEnvName = "CB_AUTH_FAILURE_RATION_LIMIT"
)

type CircuitBreakerConfig interface {
	Name() string
	MaxRequests() uint32
	Interval() time.Duration
	Timeout() time.Duration
	FailureRatioLimit() float64
}

type circuitBreakerConfig struct {
	name              string
	maxRequests       uint32
	interval          time.Duration
	timeout           time.Duration
	failureRatioLimit float64
}

func NewCircuitBreakerConfig() (*circuitBreakerConfig, error) {
	var (
		name              string
		maxRequests       int
		interval          int
		timeout           int
		failureRatioLimit float64
	)
	env.ToString(&name, circuitBreakerNameEnvName, "auth-service")
	env.ToInt(&maxRequests, circuitBreakerMaxRequestsEnvName, 3)
	env.ToInt(&interval, circuitBreakerIntervalEnvName, 10)
	env.ToInt(&timeout, circuitBreakerTimeoutEnvName, 5)
	env.ToFloat(&failureRatioLimit, circuitBreakerFailureRatioLimitEnvName, 0.6)

	return &circuitBreakerConfig{
		name:              name,
		maxRequests:       uint32(maxRequests),
		interval:          time.Duration(interval) * time.Second,
		timeout:           time.Duration(timeout) * time.Second,
		failureRatioLimit: failureRatioLimit,
	}, nil
}

func (c *circuitBreakerConfig) Name() string {
	return c.name
}

func (c *circuitBreakerConfig) MaxRequests() uint32 {
	return c.maxRequests
}

func (c *circuitBreakerConfig) Interval() time.Duration {
	return c.interval
}

func (c *circuitBreakerConfig) Timeout() time.Duration {
	return c.timeout
}

func (c *circuitBreakerConfig) FailureRatioLimit() float64 {
	return c.failureRatioLimit
}
