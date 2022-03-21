package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestOptions_SetValue(t *testing.T) {
    options := NewOptions().SetValue("k1", 1).SetValue("k2", 2)
    assert.Equal(t, 2, len(options))
    assert.Equal(t, 4, len(options.KeyValues()))
}

func TestOptions_SetValues(t *testing.T) {
    options := NewOptions("k1", 1, "k2", "k2")
    assert.Equal(t, 2, len(options))

    options.SetValues("k1", 2, "k3", 3)
    assert.Equal(t, 3, len(options))
}

func TestOptions_Merge(t *testing.T) {
    options := NewOptions("k1", 1, "k2", "k2").Merge(NewOptions("k1", 2, "k3", 3))
    assert.Equal(t, 3, len(options))
    assert.Equal(t, 2, options["k1"].(int))
}
