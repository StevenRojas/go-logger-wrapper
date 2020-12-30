package gologgerwrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteToken(t *testing.T) {
	logger, err := NewLogger()
	assert.Nil(t, err)
	logger.Debug("This is a debug message")
}
