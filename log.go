package balabol

import (
	"github.com/rs/zerolog"
)

// Logger is a common interface for logging libraries.
type Logger interface {
	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...interface{})
	// Error prints out an application error.
	Error(message string)
}

// AppLogger is the app's logger.
type AppLogger struct {
	provider *zerolog.Logger
}

// NewAppLogger initializes the app's logger.
func NewAppLogger(provider *zerolog.Logger) *AppLogger {
	return &AppLogger{
		provider: provider,
	}
}

// Printf prints out a message in a certain format.
func (l *AppLogger) Printf(format string, args ...interface{}) {
	l.provider.Info().Msgf(format, args...)
}

// Print prints out a message a log message.
func (l *AppLogger) Println(message string) {
	l.provider.Info().Msg(message)
}

// Print prints out a message a log message.
func (l *AppLogger) Error(message string) {
	l.provider.Error().Msg(message)
}
