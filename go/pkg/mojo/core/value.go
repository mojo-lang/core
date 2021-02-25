package core

func (m *Object) GetValue(key string) *Value {
	if m != nil && m.Values != nil {
		return m.Values[key]
	}
	return nil
}

func (m *Object) GetString(key string) string {
	if m.Values != nil {
		m := m.Values[key]
		if m != nil {
			return m.GetString()
		}
	}
	return ""
}

func (m *Object) GetBool(key string) bool {
	if m.Values != nil {
		m := m.Values[key]
		if m != nil {
			return m.GetBool()
		}
	}
	return false
}

func (m *Object) GetInt32Array(key string) []int32 {
	if m.Values != nil {
		return m.Values[key].GetInt32Array()
	}
	return []int32{}
}

func (m *Object) GetInt64Array(key string) []int64 {
	if m.Values != nil {
		return m.Values[key].GetInt64Array()
	}
	return []int64{}
}

func (m *Object) GetUint32Array(key string) []uint32 {
	if m.Values != nil {
		return m.Values[key].GetUint32Array()
	}
	return []uint32{}
}

func (m *Object) GetUint64Array(key string) []uint64 {
	if m.Values != nil {
		return m.Values[key].GetUint64Array()
	}
	return []uint64{}
}

func (m *Object) GetStringArray(key string) []string {
	if m.Values != nil {
		return m.Values[key].GetStringArray()
	}
	return []string{}
}

func (m *Object) GetObjectArray(key string) []*Object {
	if m.Values != nil {
		return m.Values[key].GetObjectArray()
	}
	return []*Object{}
}

func (m *Object) SetValue(key string, value *Value) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = value
	return m
}

func (m *Object) SetBool(key string, value bool) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewBoolValue(value)
	return m
}

func (m *Object) SetInt(key string, value int) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewIntValue(value)
	return m
}

func (m *Object) SetInt32(key string, value int32) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewInt32Value(value)
	return m
}

func (m *Object) SetInt64(key string, value int64) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewInt64Value(value)
	return m
}

func (m *Object) SetUint32(key string, value uint32) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewUint32Value(value)
	return m
}

func (m *Object) SetUint64(key string, value uint64) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewUint64Value(value)
	return m
}

func (m *Object) SetFloat32(key string, value float32) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewFloat32Value(value)
	return m
}

func (m *Object) SetFloat64(key string, value float64) *Object {
	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewFloat64Value(value)
	return m
}

func (m *Object) SetString(key string, value string) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewStringValue(value)
	return m
}

func (m *Object) SetObject(key string, value *Object) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewObjectValue(value)
	return m
}

func (m *Object) SetInt32Array(key string, vals ...int32) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewInt32ArrayValue(vals...)
	return m
}

func (m *Object) SetInt64Array(key string, vals ...int64) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewInt64ArrayValue(vals...)
	return m
}

func (m *Object) SetUint32Array(key string, vals ...uint32) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewUint32ArrayValue(vals...)
	return m
}

func (m *Object) SetUint64Array(key string, vals ...uint64) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewUint64ArrayValue(vals...)
	return m
}

func (m *Object) SetStringArray(key string, vals ...string) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewStringArrayValue(vals...)
	return m
}

func (m *Object) SetObjectArray(key string, vals ...*Object) *Object {
	if m == nil {
		return nil
	}

	if m.Values == nil {
		m.Values = make(map[string]*Value)
	}
	m.Values[key] = NewObjectArrayValue(vals...)
	return m
}

func NewObjectValue(obj *Object) *Value {
	return &Value{Value: &Value_ObjectVal{ObjectVal: obj}}
}

func NewArrayValue(values ...*Value) *Value {
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
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
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
}

func NewInt32ArrayValue(vals ...int32) *Value {
	values := make([]*Value, 0, len(vals))
	for _, v := range vals {
		values = append(values, NewInt32Value(v))
	}
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
}

func NewInt64ArrayValue(vals ...int64) *Value {
	values := make([]*Value, 0, len(vals))
	for _, v := range vals {
		values = append(values, NewInt64Value(v))
	}
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
}

func NewUint32ArrayValue(vals ...uint32) *Value {
	values := make([]*Value, 0, len(vals))
	for _, v := range vals {
		values = append(values, NewUint32Value(v))
	}
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
}

func NewUint64ArrayValue(vals ...uint64) *Value {
	values := make([]*Value, 0, len(vals))
	for _, v := range vals {
		values = append(values, NewUint64Value(v))
	}
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
}

func NewFloat32ArrayValue(vals ...float32) *Value {
	values := make([]*Value, 0, len(vals))
	for _, v := range vals {
		values = append(values, NewFloat32Value(v))
	}
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
}

func NewFloat64ArrayValue(vals ...float64) *Value {
	values := make([]*Value, 0, len(vals))
	for _, v := range vals {
		values = append(values, NewFloat64Value(v))
	}
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
}

func NewStringArrayValue(vals ...string) *Value {
	values := make([]*Value, 0, len(vals))
	for _, v := range vals {
		values = append(values, NewStringValue(v))
	}
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
}

func NewObjectArrayValue(vals ...*Object) *Value {
	values := make([]*Value, 0, len(vals))
	for _, v := range vals {
		values = append(values, NewObjectValue(v))
	}
	return &Value{Value: &Value_ValuesVal{ValuesVal: &Values{Values: values}}}
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

	return values.Values
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
