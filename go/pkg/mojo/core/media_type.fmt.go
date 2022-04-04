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

func (x *MediaType) Parse(mediaType string) error {
    if x != nil && len(mediaType) > 0 {
        segments := strings.Split(mediaType, ";")
        if len(segments) > 1 {
            //TODO
        }

        segments = strings.Split(segments[0], "/")
        if len(segments) > 1 {
            x.Type = segments[0]
            x.Subtype = segments[1]
        } else {
            return errors.New("\"has no subtype for the MediaType\"")
        }
    }
    return nil
}

func (x *MediaType) Format() string {
    if x != nil {
        buffer := bytes.Buffer{}
        buffer.WriteString(x.Type)
        buffer.WriteByte('/')
        buffer.WriteString(x.Subtype)

        if x.Parameter != nil {
            buffer.WriteByte(';')
            buffer.WriteString(x.Parameter.Key)
            buffer.WriteByte('=')
            buffer.WriteString(x.Parameter.Value.GetString()) //FIXME: support other types
        }
        return buffer.String()
    }
    return ""
}
