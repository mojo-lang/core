package core

import (
    jsoniter "github.com/json-iterator/go"
    "unsafe"
)

func init() {
    RegisterJSONTypeDecoder("core.Ordering", &OrderingCodec{})
    RegisterJSONTypeEncoder("core.Ordering", &OrderingCodec{})
}

type OrderingCodec struct {
}

func (codec *OrderingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    ordering := (*Ordering)(ptr)
    if any.ValueType() == jsoniter.StringValue {
        if err := ordering.Parse(any.ToString()); err != nil {
            iter.ReportError("Decode Ordering", err.Error())
        }
    }
}

func (codec *OrderingCodec) IsEmpty(ptr unsafe.Pointer) bool {
    ordering := (*Ordering)(ptr)
    return ordering == nil || len(ordering.Vals) == 0
}

func (codec *OrderingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    stream.WriteString((*Ordering)(ptr).Format())
}
