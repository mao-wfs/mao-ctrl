package configs

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

// CorrelatorConfig is the configure of the correlator.
type CorrelatorConfig struct {
	// Host is the correlator host.
	Host string `required:"true"`

	// Port is the port number the correlator listen on.
	Port int16 `required:"true"`
}

// LoadCorrelatorConfig loads the correlator configuration from environment variables.
func LoadCorrelatorConfig() (*CorrelatorConfig, error) {
	conf := new(CorrelatorConfig)
	if err := envconfig.Process("correlator", conf); err != nil {
		return nil, xerrors.Errorf("failed to load the correlator configuration: %w", err)
	}
	return conf, nil
}

// Addr returns the correlator address.
func (c *CorrelatorConfig) Addr() string {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return addr
}
