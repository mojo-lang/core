package core

import (
    jsoniter "github.com/json-iterator/go"
    "unsafe"
)

func init() {
    jsoniter.RegisterTypeEncoder("core.Null", &NullCodec{})
}

type NullCodec struct {
}

func (codec *NullCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
}

func (codec *NullCodec) IsEmpty(ptr unsafe.Pointer) bool {
    return true
}
