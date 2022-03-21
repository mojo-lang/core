package core

import (
    "errors"
    "unsafe"

    jsoniter "github.com/json-iterator/go"
)

func init() {
    jsoniter.RegisterTypeDecoder("core.Value", &ValueCodec{})
    jsoniter.RegisterTypeEncoder("core.Value", &ValueCodec{})
}

type ValueCodec struct {
}

func NewValueCodec() *ValueCodec {
    return &ValueCodec{}
}

func (codec *ValueCodec) DecodeAny(any jsoniter.Any) (*Value, error) {
    switch any.ValueType() {
    case jsoniter.NilValue:
        return &Value{}, nil
    case jsoniter.BoolValue:
        return NewBoolValue(any.ToBool()), nil
    case jsoniter.NumberValue:
        floatVal := any.ToFloat64()
        intVal := any.ToInt64() // not support uint64 (the highest bit is 1)

        if floatVal != float64(intVal) {
            return NewFloat64Value(floatVal), nil
        } else {
            return NewInt64Value(intVal), nil
        }
    case jsoniter.StringValue:
        return NewStringValue(any.ToString()), nil
    case jsoniter.ObjectValue:
        val := make(map[string]*Value)
        any.ToVal(&val)
        return NewMapValue(val), nil
    case jsoniter.ArrayValue:
        val := make([]*Value, 0)
        any.ToVal(&val)
        return NewArrayValue(val...), nil
    }
    return nil, errors.New("type is invalid")
}

func (codec *ValueCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    v, _ := codec.DecodeAny(any)
    (*Value)(ptr).Value = v.Value
}

func (codec *ValueCodec) IsEmpty(ptr unsafe.Pointer) bool {
    value := (*Value)(ptr)
    return value == nil || value.Value == nil
}

func (codec *ValueCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    value := (*Value)(ptr)
    switch val := value.Value.(type) {
    case *Value_BoolVal:
        stream.WriteBool(val.BoolVal)
    case *Value_Int64Val:
        stream.WriteInt64(val.Int64Val)
    case *Value_DoubleVal:
        stream.WriteFloat64Lossy(val.DoubleVal)
    case *Value_StringVal:
        stream.WriteString(val.StringVal)
    case *Value_ValuesVal:
        stream.WriteVal(val.ValuesVal.Vals)
    case *Value_ObjectVal:
        stream.WriteVal(val.ObjectVal.Vals)
    default:
        stream.WriteNil()
    }
}
