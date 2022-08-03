package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestParseArchitecture(t *testing.T) {
    ar, err := ParseArchitecture("AMD64")
    assert.NoError(t, err)
    assert.Equal(t, Architecture_ARCHITECTURE_AMD64, ar)
}
