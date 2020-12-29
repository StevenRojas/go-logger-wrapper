package gologgerwrapper

import (
	"testing"
)

func TestDeleteToken(t *testing.T) {
	logger := NewLogger(LoggerConfig{
		LogLevel: 7,
	})
	logger.Debug("This is a debug message")
}
