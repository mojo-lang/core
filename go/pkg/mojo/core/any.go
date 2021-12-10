package core

import (
	"bytes"
	"errors"
	"github.com/golang/protobuf/proto"
	"reflect"
)

const AnyTypeName = "mojo.core.Any"

func NewAny(i interface{}) *Any {
	any := &Any{}
	any.Set(i)
	return any
}

func (m *Any) Get() interface{} {
	return m.typeVal
}

func (m *Any) Set(v interface{}) error {
	if m != nil {
		m.Type = GetMojoTypeName(v)
		if len(m.Type) == 0 {
			return errors.New("unsupported type")
		}
		m.typeVal = v
	}
	return nil
}

func (m *Any) Empty() bool {
	if m != nil {
		return len(m.Type) == 0 && m.typeVal == nil
	}
	return true
}

func (m *Any) Message() proto.Message {
	return m.typeVal.(proto.Message)
}

func (m *Any) Unmarshal(bytes []byte) error {
	err := m.XXX_Unmarshal(bytes)
	if err != nil {
		return err
	}

	t := proto.MessageType(m.Type)
	msg := reflect.New(t).Interface().(proto.Message)
	err = proto.Unmarshal(m.Val, msg)
	if err != nil {
		return err
	}

	m.typeVal = msg
	return nil
}

func (m *Any) Marshal() ([]byte, error) {
	if m != nil {
		switch msg := m.typeVal.(type) {
		case proto.Message:
			b, err := proto.Marshal(msg)
			if err != nil {
				return nil, err
			}
			m.Val = b
		}

		size := m.XXX_Size()
		buffer := bytes.Buffer{}
		buffer.Grow(size)
		return m.XXX_Marshal(buffer.Bytes(), false)
	}
	return nil, errors.New("`Any` is null")
}

func GetMojoTypeName(i interface{}) string {
	switch v := i.(type) {
	case nil:
		return "mojo.core.Null"
	case int8:
		return "mojo.core.Int8"
	case uint8:
		return "mojo.core.UInt8"
	case int16:
		return "mojo.core.Int16"
	case uint16:
		return "mojo.core.UInt16"
	case int32:
		return "mojo.core.Int32"
	case uint32:
		return "mojo.core.UInt32"
	case int64:
		return "mojo.core.Int64"
	case uint64:
		return "mojo.core.UInt64"
	case int:
		return "mojo.core.Int"
	case uint:
		return "mojo.core.UInt"
	case float32:
		return "mojo.core.Float32"
	case float64:
		return "mojo.core.Float64"
	case string:
		return "mojo.core.String"
	case proto.Message:
		return string(proto.MessageReflect(v).Descriptor().FullName())
	default:
		return ""
	}
}
