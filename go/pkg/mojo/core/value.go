package core

import (
    "google.golang.org/protobuf/runtime/protoimpl"
    "math"
    "unicode/utf8"
)

const ValueTypeName = "Value"
const ValueTypeFullName = "mojo.core.Value"

// NewValue constructs a Value from a general-purpose Go interface.
//
//	╔════════════════════════╤════════════════════════════════════════════╗
//	║ Go type                │ Conversion                                 ║
//	╠════════════════════════╪════════════════════════════════════════════╣
//	║ nil                    │ stored as NullVal                          ║
//	║ bool                   │ stored as BoolVal                          ║
//	║ int, int32, int64      │ stored as NegativeVal/PositiveVal          ║
//	║ uint, uint32, uint64   │ stored as NegativeVal/PositiveVal          ║
//	║ float32, float64       │ stored as DoubleVal                        ║
//	║ string                 │ stored as StringVal; must be valid UTF-8   ║
//	║ []byte                 │ stored as BytesVal; base64-encoded to json ║
//	║ map[string]interface{} │ stored as ObjectVal                        ║
//	║ []interface{}          │ stored as ValuesVal                        ║
//	╚════════════════════════╧════════════════════════════════════════════╝
func NewValue(v interface{}) (*Value, error) {
    switch v := v.(type) {
    case nil:
        return NewNullValue(), nil
    case bool:
        return NewBoolValue(v), nil
    case int:
        return NewIntValue(v), nil
    case int32:
        return NewInt32Value(v), nil
    case int64:
        return NewInt64Value(v), nil
    case uint:
        return NewUIntValue(v), nil
    case uint32:
        return NewUInt32Value(v), nil
    case uint64:
        return NewUInt64Value(v), nil
    case float32:
        return NewFloat32Value(v), nil
    case float64:
        return NewFloat64Value(v), nil
    case string:
        if !utf8.ValidString(v) {
            return nil, protoimpl.X.NewError("invalid UTF-8 in string: %q", v)
        }
        return NewStringValue(v), nil
    case []byte:
        return NewBytesValue(v), nil
    case map[string]interface{}:
        obj, err := NewObjectFromMap(v)
        if err != nil {
            return nil, err
        }
        return NewObjectValue(obj), nil
    case []interface{}:
        array, err := NewValues(v...)
        if err != nil {
            return nil, err
        }
        return NewValuesValue(array), nil
    case *Value:
        return v, nil
    case *Object:
        return NewObjectValue(v), nil
    case *Values:
        return NewValuesValue(v), nil
    default:
        return nil, protoimpl.X.NewError("invalid type: %T", v)
    }
}

func NewNullValue() *Value {
    return &Value{Val: &Value_NullVal{NullVal: &Null{}}}
}

func NewBoolValue(val bool) *Value {
    return &Value{Val: &Value_BoolVal{BoolVal: val}}
}

func NewIntValue(val int) *Value {
    return NewInt64Value(int64(val))
}

func NewInt32Value(val int32) *Value {
    return NewInt64Value(int64(val))
}

func NewInt64Value(val int64) *Value {
    if val >= 0 {
        return &Value{Val: &Value_PositiveVal{PositiveVal: uint64(val)}}
    }
    return &Value{Val: &Value_NegativeVal{NegativeVal: -val}}
}

func NewUIntValue(val uint) *Value {
    return NewUInt64Value(uint64(val))
}

func NewUInt32Value(val uint32) *Value {
    return NewUInt64Value(uint64(val))
}

func NewUInt64Value(val uint64) *Value {
    return &Value{Val: &Value_PositiveVal{PositiveVal: val}}
}

func NewFloat32Value(val float32) *Value {
    return NewFloat64Value(float64(val))
}

func NewFloat64Value(val float64) *Value {
    return &Value{Val: &Value_DoubleVal{DoubleVal: val}}
}

func NewStringValue(val string) *Value {
    return &Value{Val: &Value_StringVal{StringVal: val}}
}

func NewBytesValue(val []byte) *Value {
    return &Value{Val: &Value_BytesVal{BytesVal: val}}
}

func NewObjectValue(obj *Object) *Value {
    return &Value{Val: &Value_ObjectVal{ObjectVal: obj}}
}

func NewValuesValue(array *Values) *Value {
    return &Value{Val: &Value_ValuesVal{ValuesVal: array}}
}

func NewArrayValue(values ...*Value) *Value {
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewMapValue(values map[string]*Value) *Value {
    return &Value{Val: &Value_ObjectVal{ObjectVal: &Object{Vals: values}}}
}

func NewIntArrayValue(ints ...int) *Value {
    values := make([]*Value, 0, len(ints))
    for _, v := range ints {
        values = append(values, NewIntValue(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewInt32ArrayValue(int32s ...int32) *Value {
    values := make([]*Value, 0, len(int32s))
    for _, v := range int32s {
        values = append(values, NewInt32Value(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewInt64ArrayValue(int64s ...int64) *Value {
    values := make([]*Value, 0, len(int64s))
    for _, v := range int64s {
        values = append(values, NewInt64Value(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewUInt32ArrayValue(uint32s ...uint32) *Value {
    values := make([]*Value, 0, len(uint32s))
    for _, v := range uint32s {
        values = append(values, NewUInt32Value(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewUInt64ArrayValue(uint64s ...uint64) *Value {
    values := make([]*Value, 0, len(uint64s))
    for _, v := range uint64s {
        values = append(values, NewUInt64Value(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewFloat32ArrayValue(float32s ...float32) *Value {
    values := make([]*Value, 0, len(float32s))
    for _, v := range float32s {
        values = append(values, NewFloat32Value(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewFloat64ArrayValue(float64s ...float64) *Value {
    values := make([]*Value, 0, len(float64s))
    for _, v := range float64s {
        values = append(values, NewFloat64Value(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewStringArrayValue(strings ...string) *Value {
    values := make([]*Value, 0, len(strings))
    for _, v := range strings {
        values = append(values, NewStringValue(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewObjectArrayValue(objects ...*Object) *Value {
    values := make([]*Value, 0, len(objects))
    for _, v := range objects {
        values = append(values, NewObjectValue(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

// ToInterface converts x to a general-purpose Go interface.
//
// Calling Value.MarshalJSON and "encoding/json".Marshal on this output produce
// semantically equivalent JSON (assuming no errors occur).
//
// Floating-point values (i.e., "NaN", "Infinity", and "-Infinity") are
// converted as strings to remain compatible with MarshalJSON.
func (x *Value) ToInterface() interface{} {
    switch v := x.GetVal().(type) {
    case *Value_BoolVal:
        if v != nil {
            return v.BoolVal
        }
    case *Value_NegativeVal:
        if v != nil {
            return -v.NegativeVal
        }
    case *Value_PositiveVal:
        if v != nil {
            return v.PositiveVal
        }
    case *Value_DoubleVal:
        if v != nil {
            switch {
            case math.IsNaN(v.DoubleVal):
                return "NaN"
            case math.IsInf(v.DoubleVal, +1):
                return "Infinity"
            case math.IsInf(v.DoubleVal, -1):
                return "-Infinity"
            default:
                return v.DoubleVal
            }
        }
    case *Value_StringVal:
        if v != nil {
            return v.StringVal
        }
    case *Value_BytesVal:
        if v != nil {
            return v.BytesVal
        }
    case *Value_ObjectVal:
        if v != nil {
            return v.ObjectVal.ToMap()
        }
    case *Value_ValuesVal:
        if v != nil {
            return v.ValuesVal.ToSlice()
        }
    }
    return nil
}

func (x *Value) GetBool() bool {
    return x.GetBoolVal()
}

func (x *Value) GetInt() int {
    return int(x.GetInt64())
}

func (x *Value) GetInt32() int32 {
    return int32(x.GetInt64())
}

func (x *Value) GetInt64() int64 {
    if negative := x.GetNegativeVal(); negative > 0 {
        return -negative
    }
    return int64(x.GetPositiveVal())
}

func (x *Value) GetUint() uint {
    return uint(x.GetPositiveVal())
}

func (x *Value) GetUint32() uint32 {
    return uint32(x.GetPositiveVal())
}

func (x *Value) GetUint64() uint64 {
    return x.GetPositiveVal()
}

func (x *Value) GetFloat32() float32 {
    return float32(x.GetFloat64())
}

func (x *Value) GetFloat64() float64 {
    return x.GetDoubleVal()
}

func (x *Value) GetDouble() float64 {
    return x.GetDoubleVal()
}

func (x *Value) GetString() string {
    return x.GetStringVal()
}

func (x *Value) GetBytes() []byte {
    return x.GetBytesVal()
}

func (x *Value) GetObject() *Object {
    return x.GetObjectVal()
}

func (x *Value) GetValues() []*Value {
    if values := x.GetValuesVal(); values != nil {
        return values.Vals
    }
    return nil
}

func (x *Value) GetObjectArray() []*Object {
    if values := x.GetValues(); len(values) > 0 {
        objects := make([]*Object, 0, len(values))
        for _, value := range values {
            objects = append(objects, value.GetObject())
        }
        return objects
    }
    return nil
}

func (x *Value) GetStringArray() []string {
    values := x.GetValues()
    array := make([]string, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetString())
    }
    return array
}

func (x *Value) GetIntArray() []int {
    values := x.GetValues()
    array := make([]int, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetInt())
    }
    return array
}

func (x *Value) GetInt32Array() []int32 {
    values := x.GetValues()
    array := make([]int32, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetInt32())
    }
    return array
}

func (x *Value) GetInt64Array() []int64 {
    values := x.GetValues()
    array := make([]int64, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetInt64())
    }
    return array
}

func (x *Value) GetUintArray() []uint {
    values := x.GetValues()
    array := make([]uint, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetUint())
    }
    return array
}

func (x *Value) GetUint32Array() []uint32 {
    values := x.GetValues()
    array := make([]uint32, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetUint32())
    }
    return array
}

func (x *Value) GetUint64Array() []uint64 {
    values := x.GetValues()
    array := make([]uint64, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetUint64())
    }
    return array
}

func (x *Value) GetFloat32Array() []float32 {
    values := x.GetValues()
    array := make([]float32, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetFloat32())
    }
    return array
}

func (x *Value) GetFloat64Array() []float64 {
    values := x.GetValues()
    array := make([]float64, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetFloat64())
    }
    return array
}
