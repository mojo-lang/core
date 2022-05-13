package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestIsSkipError(t *testing.T) {
    err := &SkipError{}
    assert.True(t, IsSkipError(err))
}
