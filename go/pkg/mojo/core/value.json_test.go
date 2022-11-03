package core

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

type ValueTag struct {
	Tag   string `json:"tag"`
	Value *Value `json:"value"`
}

func TestValuesCodec_Decode(t *testing.T) {
	json := `{"tag":"bool", "value": true}`
	vt := &ValueTag{}
	err := jsoniter.ConfigDefault.UnmarshalFromString(json, vt)
	assert.NoError(t, err)
	assert.Equal(t, "bool", vt.Tag)
	assert.Equal(t, ValueKind_VALUE_KIND_BOOLEAN, vt.Value.GetKind())
}
