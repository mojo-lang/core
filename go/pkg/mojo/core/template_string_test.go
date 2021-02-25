package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTemplateString(t *testing.T) {
	ts := NewTemplateString("/maps/tilestream/v1/layers/{layer}/tiles/{id.level}/{id.x}/{id.y}{.format}")
	assert.Equal(t, "/maps/tilestream/v1/layers/", ts.Segments[0].Content)
	assert.Equal(t, "layer", ts.Segments[1].Content)
	assert.True(t, ts.Segments[1].Templated)
}
