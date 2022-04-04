package core

import jsoniter "github.com/json-iterator/go"

const ObjectTypeName = "Object"
const ObjectTypeFullName = "mojo.core.Object"

func NewObject() *Object {
    return &Object{}
}

func (x *Object) To(value interface{}) error {
    if x != nil {
        if json, err := jsoniter.ConfigFastest.Marshal(x); err != nil {
            return err
        } else {
            return jsoniter.ConfigFastest.Unmarshal(json, value)
        }
    }
    return nil
}

func (x *Object) From(value interface{}) error {
    if x != nil {
        if json, err := jsoniter.ConfigFastest.Marshal(value); err != nil {
            return err
        } else {
            return jsoniter.ConfigFastest.Unmarshal(json, x)
        }
    }
    return nil
}

func (x *Object) GetValue(key string) *Value {
    if x != nil && x.Vals != nil {
        return x.Vals[key]
    }
    return nil
}

func (x *Object) GetString(key string) string {
    if x.Vals != nil {
        if v, ok := x.Vals[key]; ok {
            return v.GetString()
        }
    }
    return ""
}

func (x *Object) GetBool(key string) bool {
    if x.Vals != nil {
        if v, ok := x.Vals[key]; ok {
            return v.GetBool()
        }
    }
    return false
}

func (x *Object) GetInt32(key string) int32 {
    if x.Vals != nil {
        if v, ok := x.Vals[key]; ok {
            return v.GetInt32()
        }
    }
    return 0
}

func (x *Object) GetInt64(key string) int64 {
    if x.Vals != nil {
        if v, ok := x.Vals[key]; ok {
            return v.GetInt64()
        }
    }
    return 0
}

func (x *Object) GetUint64(key string) uint64 {
    if x.Vals != nil {
        if v, ok := x.Vals[key]; ok {
            return v.GetUint64()
        }
    }
    return 0
}

func (x *Object) GetInt32Array(key string) []int32 {
    if x.Vals != nil {
        return x.Vals[key].GetInt32Array()
    }
    return []int32{}
}

func (x *Object) GetInt64Array(key string) []int64 {
    if x.Vals != nil {
        return x.Vals[key].GetInt64Array()
    }
    return []int64{}
}

func (x *Object) GetUint32Array(key string) []uint32 {
    if x.Vals != nil {
        return x.Vals[key].GetUint32Array()
    }
    return []uint32{}
}

func (x *Object) GetUint64Array(key string) []uint64 {
    if x.Vals != nil {
        return x.Vals[key].GetUint64Array()
    }
    return []uint64{}
}

func (x *Object) GetStringArray(key string) []string {
    if x.Vals != nil {
        return x.Vals[key].GetStringArray()
    }
    return []string{}
}

func (x *Object) GetObjectArray(key string) []*Object {
    if x.Vals != nil {
        return x.Vals[key].GetObjectArray()
    }
    return []*Object{}
}

func (x *Object) SetValue(key string, value *Value) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = value
    return x
}

func (x *Object) SetBool(key string, value bool) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewBoolValue(value)
    return x
}

func (x *Object) SetInt(key string, value int) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewIntValue(value)
    return x
}

func (x *Object) SetInt32(key string, value int32) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewInt32Value(value)
    return x
}

func (x *Object) SetInt64(key string, value int64) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewInt64Value(value)
    return x
}

func (x *Object) SetUint32(key string, value uint32) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewUint32Value(value)
    return x
}

func (x *Object) SetUint64(key string, value uint64) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewUint64Value(value)
    return x
}

func (x *Object) SetFloat32(key string, value float32) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewFloat32Value(value)
    return x
}

func (x *Object) SetFloat64(key string, value float64) *Object {
    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewFloat64Value(value)
    return x
}

func (x *Object) SetString(key string, value string) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewStringValue(value)
    return x
}

func (x *Object) SetObject(key string, value *Object) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewObjectValue(value)
    return x
}

func (x *Object) SetInt32Array(key string, vals ...int32) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewInt32ArrayValue(vals...)
    return x
}

func (x *Object) SetInt64Array(key string, vals ...int64) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewInt64ArrayValue(vals...)
    return x
}

func (x *Object) SetUint32Array(key string, vals ...uint32) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewUint32ArrayValue(vals...)
    return x
}

func (x *Object) SetUint64Array(key string, vals ...uint64) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewUint64ArrayValue(vals...)
    return x
}

func (x *Object) SetStringArray(key string, vals ...string) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewStringArrayValue(vals...)
    return x
}

func (x *Object) SetObjectArray(key string, vals ...*Object) *Object {
    if x == nil {
        return nil
    }

    if x.Vals == nil {
        x.Vals = make(map[string]*Value)
    }
    x.Vals[key] = NewObjectArrayValue(vals...)
    return x
}
