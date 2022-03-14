package core

import (
    "encoding/base64"
    "fmt"
    jsoniter "github.com/json-iterator/go"
    "strconv"
    "unsafe"
)

func init() {
    jsoniter.RegisterTypeDecoder("core.BoolValue", &BoolValueCodec{})
    jsoniter.RegisterTypeEncoder("core.BoolValue", &BoolValueCodec{})
    jsoniter.RegisterTypeDecoder("core.Int64Value", &Int64ValueCodec{})
    jsoniter.RegisterTypeEncoder("core.Int64Value", &Int64ValueCodec{})
    jsoniter.RegisterTypeDecoder("core.UInt64Value", &UInt64ValueCodec{})
    jsoniter.RegisterTypeEncoder("core.UInt64Value", &UInt64ValueCodec{})
    jsoniter.RegisterTypeDecoder("core.Int32Value", &Int32ValueCodec{})
    jsoniter.RegisterTypeEncoder("core.Int32Value", &Int32ValueCodec{})
    jsoniter.RegisterTypeDecoder("core.UInt32Value", &UInt32ValueCodec{})
    jsoniter.RegisterTypeEncoder("core.UInt32Value", &UInt32ValueCodec{})
    jsoniter.RegisterTypeDecoder("core.FloatValue", &FloatValueCodec{})
    jsoniter.RegisterTypeEncoder("core.FloatValue", &FloatValueCodec{})
    jsoniter.RegisterTypeDecoder("core.DoubleValue", &DoubleValueCodec{})
    jsoniter.RegisterTypeEncoder("core.DoubleValue", &DoubleValueCodec{})
    jsoniter.RegisterTypeDecoder("core.StringValue", &StringValueCodec{})
    jsoniter.RegisterTypeEncoder("core.StringValue", &StringValueCodec{})
    jsoniter.RegisterTypeDecoder("core.BytesValue", &BytesValueCodec{})
    jsoniter.RegisterTypeEncoder("core.BytesValue", &BytesValueCodec{})
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

type FloatValueCodec struct {
}

func (codec *FloatValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    (*FloatValue)(ptr).Val = iter.ReadAny().ToFloat32()
}

func (codec *FloatValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    stream.WriteString(fmt.Sprint((*FloatValue)(ptr).Val))
}

func (codec *FloatValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    e := (*FloatValue)(ptr)
    return e == nil
}

type DoubleValueCodec struct {
}

func (codec *DoubleValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    (*DoubleValue)(ptr).Val = iter.ReadAny().ToFloat64()
}

func (codec *DoubleValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    stream.WriteString(fmt.Sprint((*DoubleValue)(ptr).Val))
}

func (codec *DoubleValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    e := (*DoubleValue)(ptr)
    return e == nil
}

type StringValueCodec struct {
}

func (codec *StringValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    (*StringValue)(ptr).Val = iter.ReadAny().ToString()
}

func (codec *StringValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    stream.WriteString((*StringValue)(ptr).Val)
}

func (codec *StringValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    e := (*StringValue)(ptr)
    return e == nil && len(e.Val) > 0
}

type BytesValueCodec struct {
}

func (codec *BytesValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    if bytes, err := base64.StdEncoding.DecodeString(iter.ReadAny().ToString()); err != nil {
        iter.ReportError("BytesValueCodec:Decode", err.Error())
    } else {
        (*BytesValue)(ptr).Val = bytes
    }
}

func (codec *BytesValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    stream.WriteString(base64.StdEncoding.EncodeToString((*BytesValue)(ptr).Val))
}

func (codec *BytesValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    e := (*BytesValue)(ptr)
    return e == nil && len(e.Val) > 0
}
