package config

import (
	"github.com/satanaroom/auth/pkg/env"
)

var _ TLSConfig = (*tlsConfig)(nil)

const (
	tlsCertFileEnvName = "TLS_AUTH_CERT_FILE"
	tlsKeyFileEnvName  = "TLS_AUTH_KEY_FILE"
)

type TLSConfig interface {
	CertFile() string
	KeyFile() string
}

type tlsConfig struct {
	certFile string
	keyFile  string
}

func NewTLSConfig() (*tlsConfig, error) {
	var (
		certFile string
		keyFile  string
	)
	env.ToString(&certFile, tlsCertFileEnvName, "service.pem")
	env.ToString(&keyFile, tlsKeyFileEnvName, "service.key")

	return &tlsConfig{
		certFile: certFile,
		keyFile:  keyFile,
	}, nil
}

func (c *tlsConfig) CertFile() string {
	return c.certFile
}

func (c *tlsConfig) KeyFile() string {
	return c.keyFile
}
