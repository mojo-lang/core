package core

import (
	"fmt"
	"unicode/utf8"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const ObjectTypeName = "Object"
const ObjectTypeFullName = "mojo.core.Object"

func NewObject() *Object {
	return &Object{}
}

// NewObjectFromMap constructs a Struct from a general-purpose Go map.
// The map keys must be valid UTF-8.
// The map values are converted using NewValue.
func NewObjectFromMap(v map[string]interface{}) (*Object, error) {
	x := &Object{Vals: make(map[string]*Value, len(v))}
	for k, v := range v {
		if !utf8.ValidString(k) {
			return nil, protoimpl.X.NewError("invalid UTF-8 in string: %q", k)
		}
		var err error
		x.Vals[k], err = NewValue(v)
		if err != nil {
			return nil, err
		}
	}
	return x, nil
}

func NewObjectFromKeyValues(kvs ...interface{}) (*Object, error) {
	if len(kvs)%2 != 0 {
		return nil, fmt.Errorf("invalid key value pairs")
	}

	m := make(map[string]interface{})
	for i := 0; i < len(kvs); i += 2 {
		key, ok := kvs[i].(string)
		if !ok {
			return nil, fmt.Errorf("invalid key type, must be string. key: %v", kvs[i])
		}
		m[key] = kvs[i+1]
	}
	return NewObjectFromMap(m)
}

func NewObjectFrom(value interface{}) (*Object, error) {
	obj := &Object{}
	return obj, obj.From(value)
}

func MergeObjects(objs ...*Object) *Object {
	obj := &Object{
		Vals: make(map[string]*Value),
	}
	for _, o := range objs {
		obj.Merge(o)
	}
	return obj
}

func (x *Object) ToMap() interface{} {
	if x != nil && x.Vals != nil {
		return x.Vals
	}
	return nil
}

func (x *Object) ToMapInterface() map[string]interface{} {
	if x != nil && x.Vals != nil {
		vs := make(map[string]interface{})
		for k, v := range x.Vals {
			vs[k] = v.ToInterface()
		}
		return vs
	}
	return nil
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

func (x *Object) IsEmpty() bool {
	if x != nil {
		return len(x.Vals) == 0
	}
	return true
}

func (x *Object) GetValue(key string) *Value {
	if x != nil && x.Vals != nil {
		return x.Vals[key]
	}
	return nil
}

func (x *Object) GetString(key string) string {
	return x.GetValue(key).GetString()
}

func (x *Object) GetBool(key string) bool {
	return x.GetValue(key).GetBool()
}

func (x *Object) GetInt32(key string) int32 {
	return x.GetValue(key).GetInt32()
}

func (x *Object) GetInt64(key string) int64 {
	return x.GetValue(key).GetInt64()
}

func (x *Object) GetUint64(key string) uint64 {
	return x.GetValue(key).GetUint64()
}

func (x *Object) GetInt32Array(key string) []int32 {
	return x.GetValue(key).GetInt32Array()
}

func (x *Object) GetInt64Array(key string) []int64 {
	return x.GetValue(key).GetInt64Array()
}

func (x *Object) GetUint32Array(key string) []uint32 {
	return x.GetValue(key).GetUint32Array()
}

func (x *Object) GetUint64Array(key string) []uint64 {
	return x.GetValue(key).GetUint64Array()
}

func (x *Object) GetStringArray(key string) []string {
	return x.GetValue(key).GetStringArray()
}

func (x *Object) GetObjectArray(key string) []*Object {
	return x.GetValue(key).GetObjectArray()
}

func (x *Object) init() {
	if x != nil && x.Vals == nil {
		x.Vals = make(map[string]*Value)
	}
}

func (x *Object) SetValue(key string, value *Value) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = value
	}

	return x
}

func (x *Object) SetBool(key string, value bool) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewBoolValue(value)
	}
	return x
}

func (x *Object) SetInt(key string, value int) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewIntValue(value)
	}

	return x
}

func (x *Object) SetInt32(key string, value int32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewInt32Value(value)
	}
	return x
}

func (x *Object) SetInt64(key string, value int64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewInt64Value(value)
	}
	return x
}

func (x *Object) SetUint32(key string, value uint32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewUInt32Value(value)
	}
	return x
}

func (x *Object) SetUint64(key string, value uint64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewUInt64Value(value)
	}
	return x
}

func (x *Object) SetFloat32(key string, value float32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewFloat32Value(value)
	}
	return x
}

func (x *Object) SetFloat64(key string, value float64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewFloat64Value(value)
	}
	return x
}

func (x *Object) SetString(key string, value string) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewStringValue(value)
	}
	return x
}

func (x *Object) SetObject(key string, value *Object) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewObjectValue(value)
	}
	return x
}

func (x *Object) SetInt32Array(key string, vals ...int32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewInt32ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetInt64Array(key string, vals ...int64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewInt64ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetUint32Array(key string, vals ...uint32) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewUInt32ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetUint64Array(key string, vals ...uint64) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewUInt64ArrayValue(vals...)
	}
	return x
}

func (x *Object) SetStringArray(key string, vals ...string) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewStringArrayValue(vals...)
	}
	return x
}

func (x *Object) SetObjectArray(key string, vals ...*Object) *Object {
	if x != nil {
		x.init()
		x.Vals[key] = NewObjectArrayValue(vals...)
	}
	return x
}

func (x *Object) Delete(key string) *Object {
	if x != nil && len(x.Vals) > 0 {
		delete(x.Vals, key)
	}
	return x
}

func (x *Object) Merge(obj *Object) *Object {
	if x != nil {
		for k, v := range obj.GetVals() {
			x.Vals[k] = v
		}
	}
	return x
}
