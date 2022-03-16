package core

import (
    "bytes"
    "errors"
    "reflect"

    "github.com/golang/protobuf/proto"
)

const AnyTypeName = "Any"
const AnyTypeFullName = "mojo.core.Any"

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
        m.Type = GetMojoTypeFullName(v)
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

func GetMojoTypeFullName(i interface{}) string {
    switch v := i.(type) {
    case nil:
        return NullTypeFullName
    case int8:
        return Int8TypeFullName
    case uint8:
        return UInt8TypeFullName
    case int16:
        return Int16TypeFullName
    case uint16:
        return UInt16TypeFullName
    case int32:
        return Int32TypeFullName
    case uint32:
        return UInt32TypeFullName
    case int64:
        return Int64TypeFullName
    case uint64:
        return UInt64TypeFullName
    case int:
        return IntTypeFullName
    case uint:
        return UIntTypeFullName
    case float32:
        return FloatTypeFullName
    case float64:
        return DoubleTypeFullName
    case string:
        return StringTypeFullName
    case proto.Message:
        return string(proto.MessageReflect(v).Descriptor().FullName())
    default:
        return ""
    }
}
