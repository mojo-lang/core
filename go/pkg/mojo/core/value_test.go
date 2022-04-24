package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewArrayValue(t *testing.T) {
    val := NewInt32ArrayValue(1, 2, 3, 4)

    vi, ok := val.ToInterface().([]interface{})
    assert.True(t, ok)
    assert.Equal(t, 4, len(vi))
}
