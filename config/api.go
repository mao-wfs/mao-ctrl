package config

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/kelseyhightower/envconfig"
)

// APIConfig represents the conigure of the API.
type APIConfig struct {
	// Host is the API host.
	Host string

	// Port is the port number the API listen on.
	Port int `default:"3000"`
}

// LoadAPIConfig loads the API configuration from environment variables.
func LoadAPIConfig() (*APIConfig, error) {
	conf := new(APIConfig)
	if err := envconfig.Process("api", conf); err != nil {
		return nil, xerrors.Errorf("failed to load the API configuration: %w", err)
	}
	return conf, nil
}

// GetAddr returns the API address.
func (c *APIConfig) GetAddr() string {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return addr
}
