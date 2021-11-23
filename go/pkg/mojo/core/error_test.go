package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAbortedError(t *testing.T) {
	err := NewAbortedError("test %s", "value")
	assert.Equal(t, Aborted.Value, err.Code.Value)
	assert.Equal(t, "test value", err.Message)
}

func TestNewAbortedError2(t *testing.T) {
	err := NewAbortedError("test value")
	assert.Equal(t, Aborted.Value, err.Code.Value)
	assert.Equal(t, "test value", err.Message)
}
