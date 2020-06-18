package balabol

import "time"

// ServerConfig is an HTTP-related options set.
type ServerConfig struct {
	// Listen is the network address the app should be running on.
	Listen string
	// ReadTimeout defines read_timeout from the client.
	ReadTimeout time.Duration
	// IdleTimeout defines how long to wait for the next timeout if the connection has keepalive enabled.
	IdleTimeout time.Duration
	// WriteTimeout specifies how long to wait between writes of a response.
	WriteTimeout time.Duration
}

// AppConfig is the application config.
type AppConfig struct {
	ServerConfig ServerConfig
}
