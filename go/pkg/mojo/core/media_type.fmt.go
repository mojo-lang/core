package core

import (
	"bytes"
	"errors"
	"strings"
)

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
