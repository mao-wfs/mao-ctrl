package config

import (
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var confPath, _ = filepath.Abs("./config/config-sample.toml")

// Config represents the configure of MAO-WFS controller.
type Config struct {
	API    APIConfig    `toml:"api"`
	Device DeviceConfig `toml:"device"`
}

// LoadConfig loads the configure of MAO-WFS controller.
// TODO: Refactor
func LoadConfig() (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(confPath, &conf); err != nil {
		return conf, err
	}
	return conf, nil
}

// GetAPIConfig returns the configure of the API.
func (c *Config) GetAPIConfig() APIConfig {
	return c.API
}

// GetDeviceConfig returns the configure of devices.
func (c *Config) GetDeviceConfig() DeviceConfig {
	return c.Device
}
