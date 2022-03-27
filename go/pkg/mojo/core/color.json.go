package core

import (
    jsoniter "github.com/json-iterator/go"
    "unsafe"
)

func init() {
    jsoniter.RegisterTypeDecoder("core.Color", &ColorStringCodec{})
    jsoniter.RegisterTypeEncoder("core.Color", &ColorStringCodec{})
}

// BareColor will be jsonify to raw, without any codec
type BareColor Color

type ColorStringCodec struct {
    IsFieldPointer bool
}

func (codec *ColorStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    value := iter.ReadString()
    color := codec.color(ptr)
    if color == nil {
        color = &Color{}
        *(**Color)(ptr) = color
    }

    if err := color.Parse(value); err != nil {
        iter.ReportError("ColorStringCodec", err.Error())
    }
}

func (codec *ColorStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
    return codec.color(ptr) != nil
}

func (codec *ColorStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    color := codec.color(ptr)
    stream.WriteString(color.Format())
}

func (codec *ColorStringCodec) color(ptr unsafe.Pointer) *Color {
    if codec.IsFieldPointer {
        return *(**Color)(ptr)
    }
    return (*Color)(ptr)
}

type ColorStructCodec struct {
    IsFieldPointer bool
}

func (codec *ColorStructCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    color := codec.bareColor(ptr)
    if value := iter.ReadAny(); value.ValueType() == jsoniter.ObjectValue {
        if color == nil {
            color = &BareColor{}
            *(**BareColor)(ptr) = color
        }
        value.ToVal(color)
    }
}

func (codec *ColorStructCodec) IsEmpty(ptr unsafe.Pointer) bool {
    color := codec.bareColor(ptr)
    return color == nil
}

func (codec *ColorStructCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    c := codec.bareColor(ptr)
    stream.WriteVal(c)
}

func (codec *ColorStructCodec) bareColor(ptr unsafe.Pointer) *BareColor {
    if codec.IsFieldPointer {
        return *(**BareColor)(ptr)
    }
    return (*BareColor)(ptr)
}
