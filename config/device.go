package config

import (
	"fmt"
)

// DeviceConfig represents the configure of MAO-WFS devices.
type DeviceConfig struct {
	Correlator CorrelatorConfig `toml:"correlator"`
	FG         FGConfig         `toml:"fg"`
}

// GetCorrelatorConfig returns the correlator's configure.
func (c *DeviceConfig) GetCorrelatorConfig() CorrelatorConfig {
	return c.Correlator
}

// GetFGConfig returns the FG's configure.
func (c *DeviceConfig) GetFGConfig() FGConfig {
	return c.FG
}

// CorrelatorConfig represents the configure of the correlator.
type CorrelatorConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

// GetAddr returns the correlator address.
func (c *CorrelatorConfig) GetAddr() string {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return addr
}

// FGConfig represents the configure of the FG.
type FGConfig struct {
	Host  string `toml:"host"`
	Port  int    `toml:"port"`
	Order []int  `toml:"switch-order"`
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
