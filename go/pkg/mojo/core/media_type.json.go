package core

import (
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.MediaType", &MediaTypeCodec{})
	jsoniter.RegisterTypeEncoder("core.MediaType", &MediaTypeCodec{})
}

type MediaTypeCodec struct {
}

func (codec *MediaTypeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadString()
	mediaType := (*MediaType)(ptr)

	if len(value) > 0 {
		err := mediaType.Parse(value)
		if err != nil {
			iter.ReportError("parse MediaType", err.Error())
		}
	}
}

func (codec *MediaTypeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	mediaType := (*MediaType)(ptr)
	return mediaType == nil || (len(mediaType.Type) == 0 || len(mediaType.Subtype) == 0)
}

func (codec *MediaTypeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	mediaType := (*MediaType)(ptr)
	stream.WriteRaw(mediaType.Type)
	stream.WriteRaw("/")
	stream.WriteRaw(mediaType.Subtype)

	if mediaType.Parameter != nil {
		stream.WriteRaw(";")
		stream.WriteRaw(mediaType.Parameter.Key)
		stream.WriteRaw("=")
		stream.WriteVal(mediaType.Parameter.Value)
	}
}
