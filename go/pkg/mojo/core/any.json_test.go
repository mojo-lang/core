package core

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAnyCodec_Decode(t *testing.T) {
    str := `{"@type": "mojo.core.Checksum", "algorithm": "sha256", "value": "sha256-checksum-value"}`

    any := &Any{}
    err := jsoniter.ConfigFastest.Unmarshal([]byte(str), any)
    assert.NoError(t, err)
    assert.Equal(t, Checksum_ALGORITHM_SHA256, any.Get().(*Checksum).Algorithm)
    assert.Equal(t, "sha256-checksum-value", any.Get().(*Checksum).Value)
}
