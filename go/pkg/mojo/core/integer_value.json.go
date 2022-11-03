package core

import (
	"strconv"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Int32Value", &Int32ValueCodec{})
	RegisterJSONTypeEncoder("core.Int32Value", &Int32ValueCodec{})
	RegisterJSONTypeDecoder("core.UInt32Value", &UInt32ValueCodec{})
	RegisterJSONTypeEncoder("core.UInt32Value", &UInt32ValueCodec{})
	RegisterJSONTypeDecoder("core.Int64Value", &Int64ValueCodec{})
	RegisterJSONTypeEncoder("core.Int64Value", &Int64ValueCodec{})
	RegisterJSONTypeDecoder("core.UInt64Value", &UInt64ValueCodec{})
	RegisterJSONTypeEncoder("core.UInt64Value", &UInt64ValueCodec{})
}

type Int32ValueCodec struct {
}

func (codec *Int32ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	(*Int32Value)(ptr).Val = any.ToInt32()
}

func (codec *Int32ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := (*Int32Value)(ptr)
	stream.WriteRaw(strconv.FormatInt(int64(v.Val), 10))
}

func (codec *Int32ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	v := (*Int32Value)(ptr)
	return v == nil
}

type UInt32ValueCodec struct {
}

func (codec *UInt32ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	(*UInt32Value)(ptr).Val = any.ToUint32()
}

func (codec *UInt32ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := (*UInt32Value)(ptr)
	stream.WriteRaw(strconv.FormatUint(uint64(v.Val), 10))
}

func (codec *UInt32ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	v := (*UInt32Value)(ptr)
	return v == nil
}

type Int64ValueCodec struct {
}

func (codec *Int64ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	(*Int64Value)(ptr).Val = any.ToInt64()
}

func (codec *Int64ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := (*Int64Value)(ptr)
	stream.WriteRaw(strconv.FormatInt(v.Val, 10))
}

func (codec *Int64ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	v := (*Int64Value)(ptr)
	return v == nil
}

type UInt64ValueCodec struct {
}

func (codec *UInt64ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	(*UInt64Value)(ptr).Val = any.ToUint64()
}

func (codec *UInt64ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	v := (*UInt64Value)(ptr)
	stream.WriteRaw(strconv.FormatUint(v.Val, 10))
}

func (codec *UInt64ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
	v := (*UInt64Value)(ptr)
	return v == nil
}
