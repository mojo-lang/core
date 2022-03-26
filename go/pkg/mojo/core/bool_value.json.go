package core

import (
    jsoniter "github.com/json-iterator/go"
    "strconv"
    "unsafe"
)

func init() {
    jsoniter.RegisterTypeDecoder("core.BoolValue", &BoolValueCodec{})
    jsoniter.RegisterTypeEncoder("core.BoolValue", &BoolValueCodec{})
}

type BoolValueCodec struct {
}

func (codec *BoolValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    (*BoolValue)(ptr).Val = any.ToBool()
}

func (codec *BoolValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    e := (*BoolValue)(ptr)
    stream.WriteString(strconv.FormatBool(e.Val))
}

func (codec *BoolValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    e := (*BoolValue)(ptr)
    return e == nil
}
