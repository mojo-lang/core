package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatJson(t *testing.T) {
	const input = `{"key": "test123", "value": {"foo": 1, "bar": "baz"}}`

	output, err := FormatJson([]byte(input))
	assert.NoError(t, err)

	const expect = `{
    "key": "test123",
    "value": {
        "foo": 1,
        "bar": "baz"
    }
}`
	assert.Equal(t, expect, string(output))
}
