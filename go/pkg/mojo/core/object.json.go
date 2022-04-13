package core

import (
    "unsafe"

    jsoniter "github.com/json-iterator/go"
)

func init() {
    RegisterJSONTypeDecoder("core.Object", &ObjectCodec{})
    RegisterJSONTypeEncoder("core.Object", &ObjectCodec{})
}

type ObjectCodec struct {
}

func (codec *ObjectCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    obj := (*Object)(ptr)
    if any.ValueType() == jsoniter.ObjectValue && any.Size() > 0 {
        obj.Vals = make(map[string]*Value)
        for _, key := range any.Keys() {
            x := any.Get(key)
            switch x.ValueType() {
            case jsoniter.BoolValue:
                obj.Vals[key] = NewBoolValue(x.ToBool())
            case jsoniter.NumberValue:
                int64V := x.ToInt64()
                uint64V := x.ToUint64()
                floatV := x.ToFloat64()
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
                obj.Vals[key] = NewStringValue(x.ToString())
            case jsoniter.ArrayValue:
                values := make([]*Value, 0, x.Size())
                x.ToVal(&values)
                obj.Vals[key] = NewArrayValue(values...)
            case jsoniter.ObjectValue:
                object := &Object{}
                x.ToVal(object)
                obj.Vals[key] = NewObjectValue(object)
            }
        }
    }
}

func (codec *ObjectCodec) IsEmpty(ptr unsafe.Pointer) bool {
    object := (*Object)(ptr)
    return object == nil || len(object.Vals) == 0
}

func (codec *ObjectCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    object := (*Object)(ptr)
    stream.WriteVal(object.Vals)
}
