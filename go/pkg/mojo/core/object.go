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