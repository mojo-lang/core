package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionTag(t *testing.T)  {
	assert.True(t, IsVersionTag("v1"))
	assert.True(t, IsVersionTag("v0"))
	assert.True(t, IsVersionTag("v1beta"))
	assert.True(t, IsVersionTag("v1beta2"))
	assert.False(t, IsVersionTag("vbeta"))
}