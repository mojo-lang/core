package core

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.Object", &ObjectCodec{})
	jsoniter.RegisterTypeEncoder("core.Object", &ObjectCodec{})

	jsoniter.RegisterTypeDecoder("core.Values", &ValuesCodec{})
	jsoniter.RegisterTypeEncoder("core.Values", &ValuesCodec{})

	jsoniter.RegisterTypeDecoder("core.Value", &ValueCodec{})
	jsoniter.RegisterTypeEncoder("core.Value", &ValueCodec{})
}

type ObjectCodec struct {
}

func (codec *ObjectCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()

	if any.ValueType() == jsoniter.ObjectValue {
		obj := (*Object)(ptr)
		obj.Values = make(map[string]*Value)
		for _, key := range any.Keys() {
			m := any.Get(key)
			switch m.ValueType() {
			case jsoniter.BoolValue:
				obj.Values[key] = NewBoolValue(m.ToBool())
			case jsoniter.NumberValue:
				int64V := m.ToInt64()
				uint64V := m.ToUint64()
				floatV := m.ToFloat64()
				if floatV == float64(int64V) {
					if uint64V == uint64(int64V) {
						obj.Values[key] = NewUint64Value(uint64V)
					} else {
						obj.Values[key] = NewInt64Value(int64V)
					}
				} else {
					obj.Values[key] = NewFloat64Value(floatV)
				}
			case jsoniter.StringValue:
				obj.Values[key] = NewStringValue(m.ToString())
			case jsoniter.ArrayValue:
				values := make([]*Value, 0, m.Size())
				m.ToVal(&values)
				obj.Values[key] = NewArrayValue(values...)
			case jsoniter.ObjectValue:
				object := &Object{}
				m.ToVal(object)
				obj.Values[key] = NewObjectValue(object)
			}
		}
	}
}

func (codec *ObjectCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(((*Object)(ptr)).Values) == 0
}

func (codec *ObjectCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	object := (*Object)(ptr)
	stream.WriteVal(object.Values)
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
		stream.WriteVal(val.ValuesVal.Values)
	case *Value_ObjectVal:
		stream.WriteVal(val.ObjectVal.Values)
	default:
		stream.WriteNil()
	}
}
