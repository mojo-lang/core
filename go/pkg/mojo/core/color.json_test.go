package core

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/stretchr/testify/assert"
    "testing"
)

const colorJsonStr = `"#fff0f5"`

func TestColorStringCodec_Decode(t *testing.T) {
    c := &Color{}
    jsoniter.UnmarshalFromString(colorJsonStr, c)
    assert.Equal(t, color.Red, c.Red)
}

func TestColorStringCodec_Encode(t *testing.T) {
    str, err := jsoniter.MarshalToString(color)
    assert.NoError(t, err)
    assert.Equal(t, colorJsonStr, str)
}

type ColorWrap struct {
    Color *Color `json:"color"`
}

func init() {
    jsoniter.RegisterFieldEncoder("core.ColorWrap", "Color", &ColorStructCodec{IsFieldPointer: true})
    jsoniter.RegisterFieldDecoder("core.ColorWrap", "Color", &ColorStructCodec{IsFieldPointer: true})
}

const colorWrapJson = `{"color":{"red":255,"green":240,"blue":245,"alpha":0.1}}`

func TestColorStructCodec_Encode(t *testing.T) {
    json, err := jsoniter.ConfigDefault.MarshalToString(&ColorWrap{Color: colorAlpha})
    assert.NoError(t, err)
    assert.Equal(t, colorWrapJson, json)
}

func TestColorStructCodec_Decode(t *testing.T) {
    v := &ColorWrap{}
    err := jsoniter.ConfigDefault.UnmarshalFromString(colorWrapJson, v)
    assert.NoError(t, err)
    assert.Equal(t, colorAlpha.Red, v.Color.Red)
    assert.Equal(t, colorAlpha.GetAlpha(), v.Color.GetAlpha())
}
