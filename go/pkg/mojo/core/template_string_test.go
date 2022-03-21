package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewTemplateString(t *testing.T) {
    ts := NewTemplateString(urlTemplate)
    assert.Equal(t, "/maps/tiles/v1/layers/", ts.Segments[0].Content)
    assert.Equal(t, "layer", ts.Segments[1].Content)
    assert.True(t, ts.Segments[1].Templated)
}
