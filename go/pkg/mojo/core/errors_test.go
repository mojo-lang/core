package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewAbortedError(t *testing.T) {
    err := NewAbortedError("test %s", "value")
    assert.Equal(t, Aborted.Val, err.Code.Val)
    assert.Equal(t, "test value", err.Message)
}

func TestNewAbortedError2(t *testing.T) {
    err := NewAbortedError("test value")
    assert.Equal(t, Aborted.Val, err.Code.Val)
    assert.Equal(t, "test value", err.Message)
}

func TestIsAbortedError(t *testing.T) {
    err := NewAbortedError("aborted")
    assert.True(t, IsAbortedError(err))
}

func TestIsCancelledError(t *testing.T) {
    assert.True(t, IsCancelledError(NewCancelledError("cancel")))
}

func TestNewBadRequestError(t *testing.T) {
    err := NewBadRequestError("bad request")
    assert.Equal(t, BadRequest, err.Code)
}
