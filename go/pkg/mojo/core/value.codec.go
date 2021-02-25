package core

import (
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.Object", &ObjectJsonCodec{})
	jsoniter.RegisterTypeEncoder("core.Object", &ObjectJsonCodec{})

	jsoniter.RegisterTypeDecoder("core.Values", &ValuesJsonCodec{})
	jsoniter.RegisterTypeEncoder("core.Values", &ValuesJsonCodec{})

	jsoniter.RegisterTypeDecoder("core.Value", &ValueJsonCodec{})
	jsoniter.RegisterTypeEncoder("core.Value", &ValueJsonCodec{})
}

type ObjectJsonCodec struct {
}

func (codec *ObjectJsonCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
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

func (codec *ObjectJsonCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return len(((*Object)(ptr)).Values) == 0
}

func (codec *ObjectJsonCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	object := (*Object)(ptr)
	stream.WriteVal(object.Values)
}

type ValuesJsonCodec struct {
}

func (codec *ValuesJsonCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
}

func (codec *ValuesJsonCodec) IsEmpty(ptr unsafe.Pointer) bool {
	values := (*Values)(ptr)
	return values == nil || len(values.Values) == 0
}

func (codec *ValuesJsonCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
}

type ValueJsonCodec struct {
}

func (codec *ValueJsonCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	val := (*Value)(ptr)
	switch any.ValueType() {
	case jsoniter.BoolValue:
		val.Value = &Value_BoolVal{BoolVal: any.ToBool()}
	case jsoniter.NumberValue:
		int64V := any.ToInt64() // if integer overflow will change pos to neg int
		floatV := any.ToFloat64()
		if floatV == float64(int64V) { // [-2^53, 2^53]
			val.Value = &Value_Int64Val{Int64Val: int64V}
		} else {
			val.Value = &Value_DoubleVal{DoubleVal: floatV}
		}
	case jsoniter.StringValue:
		val.Value = &Value_StringVal{StringVal: any.ToString()}
	case jsoniter.ArrayValue:
		values := make([]*Value, 0, any.Size())
		any.ToVal(&values)
		val.Value = &Value_ValuesVal{ValuesVal: &Values{Values: values}}
	case jsoniter.ObjectValue:
		obj := &Object{}
		any.ToVal(obj)
		val.Value = &Value_ObjectVal{ObjectVal: obj}
	}
}

func (codec *ValueJsonCodec) IsEmpty(ptr unsafe.Pointer) bool {
	value := (*Value)(ptr)
	return value == nil || value.Value == nil
}

func (codec *ValueJsonCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
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
	}
}
