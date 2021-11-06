package core

import (
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.Url", &UrlCodec{})
	jsoniter.RegisterTypeEncoder("core.Url", &UrlCodec{})
}

type UrlCodec struct {
}

func (codec *UrlCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadString()
	url := (*Url)(ptr)
	url.Parse(value)
}

func (codec *UrlCodec) IsEmpty(ptr unsafe.Pointer) bool {
	url := (*Url)(ptr)
	return url == nil ||
		(len(url.Path) == 0 && (url.Authority == nil || len(url.Authority.Host) == 0) && len(url.Fragment) == 0)
}

func (codec *UrlCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	url := (*Url)(ptr)
	stream.WriteString(url.Format())
}
