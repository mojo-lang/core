package core

import (
    jsoniter "github.com/json-iterator/go"
    "strconv"
    "unsafe"
)

func init() {
    jsoniter.RegisterTypeDecoder("core.Int32Value", &Int32ValueCodec{})
    jsoniter.RegisterTypeEncoder("core.Int32Value", &Int32ValueCodec{})
    jsoniter.RegisterTypeDecoder("core.UInt32Value", &UInt32ValueCodec{})
    jsoniter.RegisterTypeEncoder("core.UInt32Value", &UInt32ValueCodec{})
    jsoniter.RegisterTypeDecoder("core.Int64Value", &Int64ValueCodec{})
    jsoniter.RegisterTypeEncoder("core.Int64Value", &Int64ValueCodec{})
    jsoniter.RegisterTypeDecoder("core.UInt64Value", &UInt64ValueCodec{})
    jsoniter.RegisterTypeEncoder("core.UInt64Value", &UInt64ValueCodec{})
}

type Int32ValueCodec struct {
}

func (codec *Int32ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    (*Int32Value)(ptr).Val = any.ToInt32()
}

func (codec *Int32ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    e := (*Int32Value)(ptr)
    stream.WriteString(strconv.FormatInt(int64(e.Val), 10))
}

func (codec *Int32ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    e := (*Int32Value)(ptr)
    return e == nil
}

type UInt32ValueCodec struct {
}

func (codec *UInt32ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    (*UInt32Value)(ptr).Val = any.ToUint32()
}

func (codec *UInt32ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    e := (*UInt32Value)(ptr)
    stream.WriteString(strconv.FormatUint(uint64(e.Val), 10))
}

func (codec *UInt32ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    e := (*UInt32Value)(ptr)
    return e == nil
}

type Int64ValueCodec struct {
}

func (codec *Int64ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    (*Int64Value)(ptr).Val = any.ToInt64()
}

func (codec *Int64ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    e := (*Int64Value)(ptr)
    stream.WriteString(strconv.FormatInt(e.Val, 10))
}

func (codec *Int64ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    e := (*Int64Value)(ptr)
    return e == nil
}

type UInt64ValueCodec struct {
}

func (codec *UInt64ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    (*UInt64Value)(ptr).Val = any.ToUint64()
}

func (codec *UInt64ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    e := (*UInt64Value)(ptr)
    stream.WriteString(strconv.FormatUint(e.Val, 10))
}

func (codec *UInt64ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    e := (*UInt64Value)(ptr)
    return e == nil
}
