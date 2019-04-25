package config

import (
	"fmt"
)

// APIConfig represents the conigure of the API.
type APIConfig struct {
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
	Version string `toml:"version"`
}

// GetAddr returns the API address.
func (c *APIConfig) GetAddr() string {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	return addr
}

// GetRootEndPoint returns the root endpoint.
func (c *APIConfig) GetRootEndPoint() string {
	rootEp := fmt.Sprintf("api/v%s", c.Version)
	return rootEp
}
