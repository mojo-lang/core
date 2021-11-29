package core

import (
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.Uuid", &UuidCodec{})
	jsoniter.RegisterTypeEncoder("core.Uuid", &UuidCodec{})
}

type UuidCodec struct {
}

func (codec *UuidCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadString()
	id := (*Uuid)(ptr)
	err := id.Parse(value)
	if err != nil {
		iter.ReportError("Uuid Decode", err.Error())
	}
}

func (codec *UuidCodec) IsEmpty(ptr unsafe.Pointer) bool {
	id := (*Uuid)(ptr)
	return id == nil || len(id.Val) == 0
}

func (codec *UuidCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	id := (*Uuid)(ptr)
	stream.WriteVal(id.Format())
}
