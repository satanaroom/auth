package config

import (
	"time"

	"github.com/satanaroom/auth/pkg/env"
)

var _ RateLimiterConfig = (*rateLimiterConfig)(nil)

const (
	rateLimiterLimitEnvName  = "RATE_LIMITER_AUTH_LIMIT"
	rateLimiterPeriodEnvName = "RATE_LIMITER_AUTH_PERIOD"
)

type RateLimiterConfig interface {
	Limit() int
	Period() time.Duration
}

type rateLimiterConfig struct {
	limit  int
	period time.Duration
}

func NewRateLimiterConfig() (*rateLimiterConfig, error) {
	var (
		limit  int
		period int
	)
	env.ToInt(&limit, rateLimiterLimitEnvName, 10)
	env.ToInt(&period, rateLimiterPeriodEnvName, 1000)

	return &rateLimiterConfig{
		limit:  limit,
		period: time.Duration(period) * time.Millisecond,
	}, nil
}

func (c *rateLimiterConfig) Limit() int {
	return c.limit
}

func (c *rateLimiterConfig) Period() time.Duration {
	return c.period
}
