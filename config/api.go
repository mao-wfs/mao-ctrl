package config

import (
	"fmt"
)

// APIConfig represents the conigure of the API.
type APIConfig struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

// GetAddr returns the API address.
func (c *APIConfig) GetAddr() string {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return addr
}
