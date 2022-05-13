package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestIsError(t *testing.T) {
    assert.True(t, IsError(NewErrorFrom(400, "msg")))
}
