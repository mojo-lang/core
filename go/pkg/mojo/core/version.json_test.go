package core

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestVersionJsonCodec_Decode(t *testing.T) {
    version := &Version{}
    jsoniter.UnmarshalFromString(`"1.2.3"`, version)
    assert.Equal(t, uint64(1), version.Major)
}

func TestVersionJsonCodec_Encode(t *testing.T) {
    version := &Version{
        Major: 1,
        Minor: 2,
        Patch: 3,
    }

    str, err := jsoniter.MarshalToString(version)
    assert.NoError(t, err)
    assert.Equal(t, "1.2.3", str)
}
