package core

import (
    "errors"
    "google.golang.org/protobuf/proto"
)

const AnyTypeName = "Any"
const AnyTypeFullName = "mojo.core.Any"

func NewAny(i interface{}) *Any {
    any := &Any{}
    any.Set(i)
    return any
}

func (x *Any) Get() interface{} {
    return x.typedVal
}

func (x *Any) Set(v interface{}) error {
    if x != nil {
        x.Type = GetMojoTypeName(v)
        if len(x.Type) == 0 {
            return errors.New("unsupported type")
        }
        x.typedVal = v
    }
    return nil
}

func (x *Any) Empty() bool {
    if x != nil {
        return len(x.Type) == 0 && x.typedVal == nil
    }
    return true
}

func (x *Any) GetMessage() proto.Message {
    if msg, ok := x.typedVal.(proto.Message); ok {
        return msg
    }
    return nil
}

func GetMojoTypeName(i interface{}) string {
    switch v := i.(type) {
    case nil:
        return NullTypeName
    case int8:
        return Int8TypeName
    case uint8:
        return UInt8TypeName
    case int16:
        return Int16TypeName
    case uint16:
        return UInt16TypeName
    case int32:
        return Int32TypeName
    case uint32:
        return UInt32TypeName
    case int64:
        return Int64TypeName
    case uint64:
        return UInt64TypeName
    case int:
        return IntTypeName
    case uint:
        return UIntTypeName
    case float32:
        return Float32TypeName
    case float64:
        return Float64TypeName
    case string:
        return StringTypeName
    case proto.Message:
        return string(proto.MessageName(v))
    default:
        return ""
    }
}
