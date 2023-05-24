package config

import (
	"encoding/base64"
	"time"

	"github.com/satanaroom/auth/pkg/env"
)

var _ AuthConfig = (*authConfig)(nil)

const (
	refreshTokenSecretKeyEnvName = "REFRESH_TOKEN_SECRET_KEY"
	accessTokenSecretKeyEnvName  = "ACCESS_TOKEN_SECRET_KEY"

	refreshTokenExpirationEnvName = "REFRESH_TOKEN_EXPIRATION_MINUTES"
	accessTokenExpirationEnvName  = "ACCESS_TOKEN_EXPIRATION_MINUTES"
)

type AuthConfig interface {
	RefreshTokenSecretKey() []byte
	AccessTokenSecretKey() []byte
	RefreshTokenExpiration() time.Duration
	AccessTokenExpiration() time.Duration
}

type authConfig struct {
	refreshTokenSecretKey  []byte
	accessTokenSecretKey   []byte
	refreshTokenExpiration time.Duration
	accessTokenExpiration  time.Duration
}

func NewAuthConfig() (*authConfig, error) {
	var (
		refreshTokenSecretKey  string
		accessTokenSecretKey   string
		refreshTokenExpiration int
		accessTokenExpiration  int
	)
	env.ToString(&refreshTokenSecretKey, refreshTokenSecretKeyEnvName, "localhost:8080")
	env.ToString(&accessTokenSecretKey, accessTokenSecretKeyEnvName, "localhost:8080")
	env.ToInt(&refreshTokenExpiration, refreshTokenExpirationEnvName, 60)
	env.ToInt(&accessTokenExpiration, accessTokenExpirationEnvName, 5)

	return &authConfig{
		refreshTokenSecretKey:  decode(refreshTokenSecretKey),
		accessTokenSecretKey:   decode(accessTokenSecretKey),
		refreshTokenExpiration: time.Minute * time.Duration(refreshTokenExpiration),
		accessTokenExpiration:  time.Minute * time.Duration(accessTokenExpiration),
	}, nil
}

func decode(key string) []byte {
	return []byte(base64.StdEncoding.EncodeToString([]byte(key)))
}

func (c *authConfig) RefreshTokenSecretKey() []byte {
	return c.refreshTokenSecretKey

}
func (c *authConfig) AccessTokenSecretKey() []byte {
	return c.accessTokenSecretKey
}
func (c *authConfig) RefreshTokenExpiration() time.Duration {
	return c.refreshTokenExpiration
}
func (c *authConfig) AccessTokenExpiration() time.Duration {
	return c.accessTokenExpiration
}
