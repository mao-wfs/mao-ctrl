package configs

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

// APIConfig is the configuration of the API.
type APIConfig struct {
	// Port is the port number the API listen on.
	Port int16 `default:"3000"`
}

// LoadAPIConfig loads the API configuration from environment variables.
func LoadAPIConfig() (*APIConfig, error) {
	conf := new(APIConfig)
	if err := envconfig.Process("api", conf); err != nil {
		return nil, xerrors.Errorf("failed to load the API configuration: %w", err)
	}
	return conf, nil
}

// Addr returns the API address.
func (c APIConfig) Addr() string {
	addr := fmt.Sprintf(":%d", c.Port)
	return addr
}
