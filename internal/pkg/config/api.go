package config

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/kelseyhightower/envconfig"
)

// APIConfig represents the configuration of the API.
type APIConfig struct {
	// Host is the API host.
	Host string `envconfig:"HOST"`

	// Port is the port number the API listen on.
	Port uint16 `default:"3030"`
}

// LoadAPIConfig loads the API configuration from environment variables.
func LoadAPIConfig() (*APIConfig, error) {
	conf := new(APIConfig)
	if err := envconfig.Process("api", conf); err != nil {
		return nil, xerrors.Errorf("failed to load API the configuration: %w", err)
	}
	return conf, nil
}

// GetAddr returns the API address.
func (c APIConfig) GetAddr() string {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return addr
}
