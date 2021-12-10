package core

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyCodec_Decode(t *testing.T) {
	str := `{"@type": "mojo.core.Checksum", "algorithm": "sha256", "val": "werrdsafadfdfdfff"}`

	any := &Any{}
	err := jsoniter.ConfigFastest.Unmarshal([]byte(str), any)
	assert.NoError(t, err)
}