package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBreakError(t *testing.T) {
	err := &BreakError{}
	assert.True(t, IsBreakError(err))
}