package balabol

import "time"

// ServerConfig is an HTTP-related options set.
type ServerConfig struct {
	ReadTimeout time.Duration
	Listen      string
}

// AppConfig is the application config.
type AppConfig struct {
	ServerConfig ServerConfig
}
