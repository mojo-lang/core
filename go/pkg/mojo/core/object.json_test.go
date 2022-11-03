package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestObjectCodec_Decode(t *testing.T) {
	const Val = "{\"key\":\"value\"}"
	object := &Object{}
	err := jsoniter.ConfigFastest.Unmarshal([]byte(Val), object)
	assert.NoError(t, err)
	assert.Equal(t, "value", object.GetString("key"))
}

func TestObjectCodec_Decode_Empty(t *testing.T) {
	const Val = "{}"
	object := &Object{}
	err := jsoniter.ConfigFastest.Unmarshal([]byte(Val), object)
	assert.NoError(t, err)
}

func TestObjectCodec_Encode_Empty(t *testing.T) {
	type Foo struct {
		Name   string  `json:"name"`
		Object *Object `json:"object,omitempty"`
	}

	str, err := jsoniter.ConfigFastest.MarshalToString(&Foo{Name: "foo"})
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"foo"}`, str)

	// str, err := jsoniter.ConfigFastest.MarshalToString(&Foo{Name: "foo", Object: NewObject()})
	// assert.NoError(t, err)
	// assert.Equal(t, `{"name":"foo"}`, str)
}
