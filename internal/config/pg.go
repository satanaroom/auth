package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ PGConfig = (*pgConfig)(nil)

const pgDSNEnvName = "PG_DSN"

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func NewPGConfig() (*pgConfig, error) {
	var dsn string
	env.ToString(&dsn, pgDSNEnvName, "")

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
