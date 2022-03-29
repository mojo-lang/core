package core

import (
    jsoniter "github.com/json-iterator/go"
)

func RegisterJSONTypeDecoder(typ string, decoder jsoniter.ValDecoder) {
    jsoniter.RegisterTypeDecoder(typ, decoder)
}

func RegisterJSONTypeEncoder(typ string, encoder jsoniter.ValEncoder) {
    jsoniter.RegisterTypeEncoder(typ, encoder)
}

func RegisterJSONFieldDecoder(typ string, field string, decoder jsoniter.ValDecoder) {
    jsoniter.RegisterFieldDecoder(typ, field, decoder)
}

func RegisterJSONFieldEncoder(typ string, field string, encoder jsoniter.ValEncoder) {
    jsoniter.RegisterFieldEncoder(typ, field, encoder)
    RegisterAnyFieldEncoder(typ, field, encoder)
}
