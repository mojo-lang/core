package core

const ValueTypeName = "Value"
const ValueTypeFullName = "mojo.core.Value"

//TODO add more convenience new func
//func NewValue(v interface{}) *Value {
//
//}

func NewObjectValue(obj *Object) *Value {
    return &Value{Value: &Value_ObjectVal{ObjectVal: obj}}
}

func NewArrayValue(values ...*Value) *Value {
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewMapValue(values map[string]*Value) *Value {
    return &Value{Value: &Value_ObjectVal{ObjectVal: &Object{Vals: values}}}
}

func NewBoolValue(val bool) *Value {
    return &Value{Value: &Value_BoolVal{BoolVal: val}}
}

func NewIntValue(val int) *Value {
    return NewInt64Value(int64(val))
}

func NewInt32Value(val int32) *Value {
    return NewInt64Value(int64(val))
}

func NewInt64Value(val int64) *Value {
    return &Value{Value: &Value_Int64Val{Int64Val: val}}
}

func NewUintValue(val uint) *Value {
    return NewUint64Value(uint64(val))
}

func NewUint32Value(val uint32) *Value {
    return NewUint64Value(uint64(val))
}

func NewUint64Value(val uint64) *Value {
    return &Value{Value: &Value_Int64Val{Int64Val: int64(val)}}
}

func NewFloat32Value(val float32) *Value {
    return NewFloat64Value(float64(val))
}

func NewFloat64Value(val float64) *Value {
    return &Value{Value: &Value_DoubleVal{DoubleVal: val}}
}

func NewStringValue(val string) *Value {
    return &Value{Value: &Value_StringVal{StringVal: val}}
}

func NewIntArrayValue(vals ...int) *Value {
    values := make([]*Value, 0, len(vals))
    for _, v := range vals {
        values = append(values, NewIntValue(v))
    }
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewInt32ArrayValue(vals ...int32) *Value {
    values := make([]*Value, 0, len(vals))
    for _, v := range vals {
        values = append(values, NewInt32Value(v))
    }
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewInt64ArrayValue(vals ...int64) *Value {
    values := make([]*Value, 0, len(vals))
    for _, v := range vals {
        values = append(values, NewInt64Value(v))
    }
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewUint32ArrayValue(vals ...uint32) *Value {
    values := make([]*Value, 0, len(vals))
    for _, v := range vals {
        values = append(values, NewUint32Value(v))
    }
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewUint64ArrayValue(vals ...uint64) *Value {
    values := make([]*Value, 0, len(vals))
    for _, v := range vals {
        values = append(values, NewUint64Value(v))
    }
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewFloat32ArrayValue(vals ...float32) *Value {
    values := make([]*Value, 0, len(vals))
    for _, v := range vals {
        values = append(values, NewFloat32Value(v))
    }
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewFloat64ArrayValue(vals ...float64) *Value {
    values := make([]*Value, 0, len(vals))
    for _, v := range vals {
        values = append(values, NewFloat64Value(v))
    }
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewStringArrayValue(vals ...string) *Value {
    values := make([]*Value, 0, len(vals))
    for _, v := range vals {
        values = append(values, NewStringValue(v))
    }
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func NewObjectArrayValue(vals ...*Object) *Value {
    values := make([]*Value, 0, len(vals))
    for _, v := range vals {
        values = append(values, NewObjectValue(v))
    }
    return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Vals: values}}}
}

func (m *Value) GetBool() bool {
    if m == nil {
        return false
    }
    return m.GetBoolVal()
}

func (m *Value) GetInt() int {
    if m == nil {
        return 0
    }
    return int(m.GetInt64())
}

func (m *Value) GetInt32() int32 {
    if m == nil {
        return 0
    }

    return int32(m.GetInt64())
}

func (m *Value) GetInt64() int64 {
    if m == nil {
        return 0
    }

    return m.GetInt64Val()
}

func (m *Value) GetUint() uint {
    if m == nil {
        return 0
    }

    return uint(m.GetInt64())
}

func (m *Value) GetUint32() uint32 {
    if m == nil {
        return 0
    }

    return uint32(m.GetInt64())
}

func (m *Value) GetUint64() uint64 {
    if m == nil {
        return 0
    }

    return uint64(m.GetInt64())
}

func (m *Value) GetFloat32() float32 {
    if m == nil {
        return 0
    }

    return float32(m.GetFloat64())
}

func (m *Value) GetFloat64() float64 {
    if m == nil {
        return 0
    }

    return m.GetDoubleVal()
}

func (m *Value) GetDouble() float64 {
    if m == nil {
        return 0
    }

    return m.GetDoubleVal()
}

func (m *Value) GetString() string {
    if m == nil {
        return ""
    }

    return m.GetStringVal()
}

func (m *Value) GetObject() *Object {
    if m == nil {
        return nil
    }

    return m.GetObjectVal()
}

func (m *Value) GetValues() []*Value {
    if m == nil {
        return []*Value{}
    }
    values := m.GetValuesVal()
    if values == nil {
        return []*Value{}
    }

    return values.Vals
}

func (m *Value) GetObjectArray() []*Object {
    values := m.GetValues()
    objects := make([]*Object, 0, len(values))
    for _, value := range values {
        objects = append(objects, value.GetObject())
    }
    return objects
}

func (m *Value) GetStringArray() []string {
    values := m.GetValues()
    array := make([]string, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetString())
    }
    return array
}

func (m *Value) GetIntArray() []int {
    values := m.GetValues()
    array := make([]int, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetInt())
    }
    return array
}

func (m *Value) GetInt32Array() []int32 {
    values := m.GetValues()
    array := make([]int32, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetInt32())
    }
    return array
}

func (m *Value) GetInt64Array() []int64 {
    values := m.GetValues()
    array := make([]int64, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetInt64())
    }
    return array
}

func (m *Value) GetUintArray() []uint {
    values := m.GetValues()
    array := make([]uint, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetUint())
    }
    return array
}

func (m *Value) GetUint32Array() []uint32 {
    values := m.GetValues()
    array := make([]uint32, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetUint32())
    }
    return array
}

func (m *Value) GetUint64Array() []uint64 {
    values := m.GetValues()
    array := make([]uint64, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetUint64())
    }
    return array
}

func (m *Value) GetFloat32Array() []float32 {
    values := m.GetValues()
    array := make([]float32, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetFloat32())
    }
    return array
}

func (m *Value) GetFloat64Array() []float64 {
    values := m.GetValues()
    array := make([]float64, 0, len(values))
    for _, value := range values {
        array = append(array, value.GetFloat64())
    }
    return array
}
