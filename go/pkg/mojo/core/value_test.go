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

func TestValue_GetKind(t *testing.T) {
    val := NewInt32Value(32)
    assert.Equal(t, ValueKind_VALUE_KIND_INTEGER, val.GetKind())

    val = NewStringValue("foo")
    assert.Equal(t, ValueKind_VALUE_KIND_STRING, val.GetKind())

    val = NewStringArrayValue("foo")
    assert.Equal(t, ValueKind_VALUE_KIND_ARRAY, val.GetKind())
}
