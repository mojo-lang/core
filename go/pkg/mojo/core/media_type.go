package core

import (
	"bytes"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"unsafe"
)

const (
	ApplicationJson              = "application/json"
	ApplicationPdf               = "application/pdf"
	ApplicationOctetStream       = "application/octet-stream"
	ApplicationWwwFormUrlencoded = "application/x-www-form-urlencoded"
	ApplicationXml               = "application/xml"
	ApplicationGeoJson           = "pplication/geo+json"

	ImagePng  = "image/png"
	ImageJpeg = "image/jpeg"
	ImageSvg  = "image/svg+xml"

	TextPlain      = "text/plain"
	TextCsv        = "text/csv"
	TextCss        = "text/css"
	TextHtml       = "text/html"
	TextJavascript = "text/javascript"

	VideoMp4 = "video/mp4"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.MediaType", &MediaTypeCodec{})
	jsoniter.RegisterTypeEncoder("core.MediaType", &MediaTypeCodec{})
}

func ParseMediaType(mediaType string) (*MediaType, error) {
	mt := &MediaType{}
	err := mt.Parse(mediaType)
	if err != nil {
		return nil, err
	}
	return mt, nil
}

func (m *MediaType) Parse(mediaType string) error {
	if m != nil && len(mediaType) > 0 {
		segments := strings.Split(mediaType, ";")
		if len(segments) > 1 {
			//TODO
		}

		segments = strings.Split(segments[0], "/")
		if len(segments) > 1 {
			m.Type = segments[0]
			m.Subtype = segments[1]
		} else {
			return errors.New("\"has no subtype for the MediaType\"")
		}
	}
	return nil
}

func (m *MediaType) Format() string {
	if m != nil {
		buffer := bytes.Buffer{}
		buffer.WriteString(m.Type)
		buffer.WriteByte('/')
		buffer.WriteString(m.Subtype)

		if m.Parameter != nil {
			buffer.WriteByte(';')
			buffer.WriteString(m.Parameter.Key)
			buffer.WriteByte('=')
			buffer.WriteString(m.Parameter.Value.GetString()) //FIXME: support other types
		}
		return buffer.String()
	}
	return ""
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
