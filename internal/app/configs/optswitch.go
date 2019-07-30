package configs

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

// OptSwitchConfig is the configuration of the optical switch.
type OptSwitchConfig struct {
	// PG is the configuration of the PG.
	PG *PGConfig

	// FG is the configuration of the FG.
	FG *FGConfig
}

// LoadOptSwitchConfig loads the optical switch configuration from environment variables.
func LoadOptSwitchConfig() (*OptSwitchConfig, error) {
	pgConf, err := loadPGConfig()
	if err != nil {
		return nil, xerrors.Errorf("failed to load the optical switch configuration: %w", err)
	}
	fgConf, err := loadFGConfig()
	if err != nil {
		return nil, xerrors.Errorf("failed to load the optical switch configuration: %w", err)
	}

	conf := &OptSwitchConfig{
		PG: pgConf,
		FG: fgConf,
	}
	return conf, nil
}

// PGAddr returns the PG address.
func (c *OptSwitchConfig) PGAddr() string {
	addr := fmt.Sprintf("%s:%d", c.PG.Host, c.PG.Port)
	return addr
}

// Order returns the switching order.
func (c *OptSwitchConfig) Order() []int16 {
	return c.PG.Order
}

// FGAddr returns the FG address.
func (c *OptSwitchConfig) FGAddr() string {
	addr := fmt.Sprintf("%s:%d", c.FG.Host, c.FG.Port)
	return addr
}

// PGConfig is the configuration of the pattern generator (Model 3390 Arbitrary Waveform Generator).
type PGConfig struct {
	// Host is the PG host.
	Host string `required:"true"`

	// Port is the port number the PG listen on.
	Port int16 `required:"true"`

	// Order is the switching order.
	Order []int16 `default:"10,9,13,8,0,80,16,32"`
}

// FGConfig is the configuration of the function generator (Agilent 33500B Seriese).
type FGConfig struct {
	// Host is the FG host.
	Host string `required:"true"`

	// Port is the port number the FG listen on.
	Port int16 `required:"true"`
}

func loadPGConfig() (*PGConfig, error) {
	conf := new(PGConfig)
	if err := envconfig.Process("pg", conf); err != nil {
		return nil, xerrors.Errorf("failed to load the PG configuration: %w", err)
	}
	return conf, nil
}

func loadFGConfig() (*FGConfig, error) {
	conf := new(FGConfig)
	if err := envconfig.Process("fg", conf); err != nil {
		return nil, xerrors.Errorf("failed to load the PG configuration: %w", err)
	}
	return conf, nil
}
