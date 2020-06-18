package main

import "time"

// ServerConfig is an HTTP-related options set.
type ServerConfig struct {
	// Listen is the network address the app should be running on.
	Listen string `yaml:"listen"`
	// ReadTimeout defines read_timeout from the client.
	ReadTimeout time.Duration `yaml:"read_timeout"`
	// IdleTimeout defines how long to wait for the next timeout if the connection has keepalive enabled.
	IdleTimeout time.Duration `yaml:"idle_timeout"`
	// WriteTimeout specifies how long to wait between writes of a response.
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

// config is the application config.
type config struct {
	// ServerConfig defines http server settings.
	ServerConfig ServerConfig `yaml:"http"`
}
