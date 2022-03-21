package core

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestObjectCodec_Decode(t *testing.T) {
    const Val = "{\"key\":\"value\"}"
    object := &Object{}
    err := jsoniter.ConfigFastest.Unmarshal([]byte(Val), object)
    assert.NoError(t, err)
    assert.Equal(t, "value", object.GetString("key"))
}
