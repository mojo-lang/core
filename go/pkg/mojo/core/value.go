package core

const ValueTypeName = "Value"
const ValueTypeFullName = "mojo.core.Value"

//TODO
//func NewValue(v interface{}) *Value {
//    switch value := v.(type) {
//    case int8:
//        return NewInt64Value(int64(value))
//    }
//
//    return nil
//}

func NewObjectValue(obj *Object) *Value {
    return &Value{Val: &Value_ObjectVal{ObjectVal: obj}}
}

func NewArrayValue(values ...*Value) *Value {
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewMapValue(values map[string]*Value) *Value {
    return &Value{Val: &Value_ObjectVal{ObjectVal: &Object{Vals: values}}}
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
    return &Value{Val: &Value_Int64Val{Int64Val: val}}
}

func NewUintValue(val uint) *Value {
    return NewUint64Value(uint64(val))
}

func NewUint32Value(val uint32) *Value {
    return NewUint64Value(uint64(val))
}

func NewUint64Value(val uint64) *Value {
    return &Value{Val: &Value_Int64Val{Int64Val: int64(val)}}
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

func NewUint32ArrayValue(uint32s ...uint32) *Value {
    values := make([]*Value, 0, len(uint32s))
    for _, v := range uint32s {
        values = append(values, NewUint32Value(v))
    }
    return &Value{Val: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewUint64ArrayValue(uint64s ...uint64) *Value {
    values := make([]*Value, 0, len(uint64s))
    for _, v := range uint64s {
        values = append(values, NewUint64Value(v))
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
    return x.GetInt64Val()
}

func (x *Value) GetUint() uint {
    if x == nil {
        return 0
    }

    return uint(x.GetInt64())
}

func (x *Value) GetUint32() uint32 {
    if x == nil {
        return 0
    }

    return uint32(x.GetInt64())
}

func (x *Value) GetUint64() uint64 {
    if x == nil {
        return 0
    }

    return uint64(x.GetInt64())
}

func (x *Value) GetFloat32() float32 {
    if x == nil {
        return 0
    }

    return float32(x.GetFloat64())
}

func (x *Value) GetFloat64() float64 {
    if x == nil {
        return 0
    }

    return x.GetDoubleVal()
}

func (x *Value) GetDouble() float64 {
    if x == nil {
        return 0
    }

    return x.GetDoubleVal()
}

func (x *Value) GetString() string {
    if x == nil {
        return ""
    }

    return x.GetStringVal()
}

func (x *Value) GetObject() *Object {
    if x == nil {
        return nil
    }

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
