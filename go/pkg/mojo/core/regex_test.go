package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex_ToRegexp(t *testing.T) {
	m := NewRegex(".+").ToRegexp().MatchString("test")
	assert.True(t, m)
}
