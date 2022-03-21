package core

import (
    "unsafe"

    jsoniter "github.com/json-iterator/go"
)

func init() {
    jsoniter.RegisterTypeDecoder("core.Object", &ObjectCodec{})
    jsoniter.RegisterTypeEncoder("core.Object", &ObjectCodec{})
}

type ObjectCodec struct {
}

func (codec *ObjectCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()

    if any.ValueType() == jsoniter.ObjectValue {
        obj := (*Object)(ptr)
        obj.Vals = make(map[string]*Value)
        for _, key := range any.Keys() {
            m := any.Get(key)
            switch m.ValueType() {
            case jsoniter.BoolValue:
                obj.Vals[key] = NewBoolValue(m.ToBool())
            case jsoniter.NumberValue:
                int64V := m.ToInt64()
                uint64V := m.ToUint64()
                floatV := m.ToFloat64()
                if floatV == float64(int64V) {
                    if uint64V == uint64(int64V) {
                        obj.Vals[key] = NewUint64Value(uint64V)
                    } else {
                        obj.Vals[key] = NewInt64Value(int64V)
                    }
                } else {
                    obj.Vals[key] = NewFloat64Value(floatV)
                }
            case jsoniter.StringValue:
                obj.Vals[key] = NewStringValue(m.ToString())
            case jsoniter.ArrayValue:
                values := make([]*Value, 0, m.Size())
                m.ToVal(&values)
                obj.Vals[key] = NewArrayValue(values...)
            case jsoniter.ObjectValue:
                object := &Object{}
                m.ToVal(object)
                obj.Vals[key] = NewObjectValue(object)
            }
        }
    }
}

func (codec *ObjectCodec) IsEmpty(ptr unsafe.Pointer) bool {
    return len(((*Object)(ptr)).Vals) == 0
}

func (codec *ObjectCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    object := (*Object)(ptr)
    stream.WriteVal(object.Vals)
}
