package main

import "time"

// ServerConfig is an HTTP-related options set.
type ServerConfig struct {
	// ReadTimeout defines read_timeout from the client.
	ReadTimeout time.Duration `yaml:"read_timeout"`
	// Listen is the network address the app should be running on.
	Listen string `yaml:"listen"`
}

// config is the application config.
type config struct {
	// ServerConfig defines http server settings.
	ServerConfig ServerConfig `yaml:"http"`
}
