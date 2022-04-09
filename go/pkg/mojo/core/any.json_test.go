package core

import (
    "encoding/json"
    jsoniter "github.com/json-iterator/go"
    "github.com/stretchr/testify/assert"
    "testing"
)

const anyStr = `{"@type":"mojo.core.Error","code":"404","message":"something wrong"}`

var anyError = &Error{Code: NotFound, Message: "something wrong"}

func TestAnyCodec_Decode(t *testing.T) {
    any := &Any{}
    err := jsoniter.ConfigFastest.UnmarshalFromString(anyStr, any)

    assert.NoError(t, err)
    assert.Equal(t, anyError.Code.Val, any.Get().(*Error).Code.Val)
    assert.Equal(t, anyError.Message, any.Get().(*Error).Message)
}

func TestAnyCodec_Decode2(t *testing.T) {
    any := &Any{}
    err := json.Unmarshal([]byte(anyStr), any)

    assert.NoError(t, err)
    assert.Equal(t, anyError.Code.Val, any.Get().(*Error).Code.Val)
    assert.Equal(t, anyError.Message, any.Get().(*Error).Message)
}

func TestAnyCodec_Encode(t *testing.T) {
    any := NewAny(anyError)
    out, err := jsoniter.ConfigFastest.MarshalToString(any)

    assert.NoError(t, err)
    assert.Equal(t, anyStr, out)
}

func TestAnyCodec_Encode2(t *testing.T) {
    any := NewAny(anyError)
    out, err := json.Marshal(any)

    assert.NoError(t, err)
    assert.Equal(t, anyStr, string(out))
}
