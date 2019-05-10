package config

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/kelseyhightower/envconfig"
)

// CorrelatorConfig represents the configure of the correlator.
type CorrelatorConfig struct {
	// Host is the correlator host.
	Host string `required:"true"`

	// Port is the port number the correlator listen on.
	Port int `required:"true"`
}

// LoadCorrelatorConfig loads the correlator configuration from environment variables.
func LoadCorrelatorConfig() (*CorrelatorConfig, error) {
	conf := new(CorrelatorConfig)
	if err := envconfig.Process("correlator", conf); err != nil {
		return nil, xerrors.Errorf("failed to load the correlator configuration: %w", err)
	}
	return conf, nil
}

// GetAddr returns the correlator address.
func (c *CorrelatorConfig) GetAddr() string {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return addr
}

// FGConfig represents the configure of the FG.
type FGConfig struct {
	// Host is the FG host.
	Host string `required:"true"`

	// Port is the port number the FG listen on.
	Port int `required:"true"`

	// Order is the switching order.
	Order []int `default:"10,9,13,8,0,80,16,32"`
}

// LoadFGConfig loads the FG configuration from environment variables.
func LoadFGConfig() (*FGConfig, error) {
	conf := new(FGConfig)
	if err := envconfig.Process("fg", conf); err != nil {
		return nil, xerrors.Errorf("failed to load the FG configuration: %w", err)
	}
	return conf, nil
}

// GetAddr returns the FG address.
func (c *FGConfig) GetAddr() string {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return addr
}

// GetOrder returns the FGing order.
func (c *FGConfig) GetOrder() []int {
	return c.Order
}
