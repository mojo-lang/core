package core

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/stretchr/testify/assert"
    "testing"
)

const anyStr = `{"@type":"mojo.core.Checksum","algorithm":"SHA256","value":"sha256-checksum-value"}`

func TestAnyCodec_Decode(t *testing.T) {
    any := &Any{}
    err := jsoniter.ConfigFastest.Unmarshal([]byte(anyStr), any)
    assert.NoError(t, err)
    assert.Equal(t, Checksum_ALGORITHM_SHA256, any.Get().(*Checksum).Algorithm)
    assert.Equal(t, "sha256-checksum-value", any.Get().(*Checksum).Value)
}

func TestAnyCodec_Encode(t *testing.T) {
    any := NewAny(&Checksum{
        Algorithm: Checksum_ALGORITHM_SHA256,
        Value:     "sha256-checksum-value",
    })

    out, err := jsoniter.ConfigFastest.MarshalToString(any)
    assert.NoError(t, err)
    assert.Equal(t, anyStr, out)
}
