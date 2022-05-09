package core

import (
    "fmt"
    jsoniter "github.com/json-iterator/go"
    "net/url"
    "reflect"
    "regexp"
    "strings"
)

func NewUrlQuery() *Url_Query {
    return &Url_Query{
        Vals: make(map[string]*StringValues),
    }
}

func NewUrlQueryFrom(values url.Values) *Url_Query {
    query := &Url_Query{
        Vals: make(map[string]*StringValues),
    }
    return query.FromUrlValues(values)
}

func (x *Url_Query) FromUrlValues(values url.Values) *Url_Query {
    if x != nil {
        if x.Vals == nil {
            x.Vals = make(map[string]*StringValues)
        }

        for k, v := range values {
            x.Vals[k] = &StringValues{Vals: v}
        }
    }
    return x
}

func (x *Url_Query) ToUrlValues() url.Values {
    if x != nil && x.Vals != nil {
        values := make(url.Values)
        for k, v := range x.Vals {
            values[k] = v.Vals
        }
    }

    return nil
}

func (x *Url_Query) Has(key string) bool {
    if x != nil {
        _, ok := x.Vals[key]
        return ok
    }
    return false
}

func (x *Url_Query) Get(key string) string {
    if x == nil {
        return ""
    }
    vs := x.Vals[key]
    if vs != nil || len(vs.Vals) == 0 {
        return ""
    }
    return vs.Vals[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (x *Url_Query) Set(key, value string) {
    if x != nil {
        x.Vals[key] = &StringValues{
            Vals: []string{value},
        }
    }
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (x *Url_Query) Add(key, value string) {
    if x != nil {
        if x.Vals[key] == nil {
            x.Vals[key] = &StringValues{}
        }
        x.Vals[key].Vals = append(x.Vals[key].Vals, value)
    }
}

// Del deletes the values associated with key.
func (x *Url_Query) Del(key string) {
    if x != nil {
        delete(x.Vals, key)
    }
}

func (x *Url_Query) Format() string {
    if x != nil {
        sb := &strings.Builder{}
        first := true
        for k, vals := range x.Vals {
            if vals != nil {
                for _, v := range vals.Vals {
                    if !first {
                        sb.WriteString("&")
                    } else {
                        first = false
                    }

                    sb.WriteString(k)
                    sb.WriteString("=")
                    sb.WriteString(v)
                }
            }
        }
        return sb.String()
    }
    return ""
}

// Unmarshal
// list type
//      foo=bar&foo=baz
//      foo=bar,baz
//      foo="bar","baz"
//      foo=["bar","baz"]
// map
//      foo=key1,bar,key2,baz <not support>
// object
//      foo={"key1":"bar","key2","baz"}
func (x *Url_Query) Unmarshal(name string, value interface{}) error {
    if x == nil || x.Vals == nil {
        return nil
    }

    param, ok := x.Vals[name]
    if !ok || param == nil || len(param.Vals) == 0 {
        return nil
    }

    v := reflect.Indirect(reflect.ValueOf(value))
    switch v.Kind() {
    case reflect.Slice, reflect.Array:
        elem := v.Type().Elem()
        isStringType := elem.Kind() == reflect.String
        if elem.Kind() == reflect.Ptr {
            if _, ok := reflect.New(elem.Elem()).Interface().(StringLike); ok {
                isStringType = true
            }
        }

        var sliceValue string
        if len(param.Vals) == 1 {
            sliceValue = param.Vals[0]
        } else {
            if isStringType {
                for i := 0; i < len(param.Vals); i++ {
                    param.Vals[i] = QuoteString(param.Vals[i])
                }
            }
            sliceValue = "[" + strings.Join(param.Vals, ",") + "]"
        }
        return UnmarshalParam(sliceValue, value)
    default:
        if len(param.Vals) > 0 {
            return UnmarshalParam(param.Vals[0], value)
        }
    }
    return nil
}

func UnmarshalParam(str string, value interface{}) error {
    v := reflect.Indirect(reflect.ValueOf(value))
    switch v.Kind() {
    case reflect.String:
        v.SetString(str)
    case reflect.Slice, reflect.Array:
        elem := v.Type().Elem()
        isStringType := elem.Kind() == reflect.String
        if elem.Kind() == reflect.Ptr {
            if _, ok := reflect.New(elem.Elem()).Interface().(StringLike); ok {
                isStringType = true
            }
        }

        str = strings.TrimSpace(str)
        if str[0] != '[' {
            if isStringType {
                vals := splitQuotedString(str)
                for i := 0; i < len(vals); i++ {
                    vals[i] = strings.TrimSpace(vals[i])
                    vals[i] = QuoteString(vals[i])
                }
                str = "[" + strings.Join(vals, ",") + "]"
            } else {
                str = "[" + str + "]"
            }
        }
        return jsoniter.Unmarshal([]byte(str), value)
    default:
        if _, ok := reflect.New(v.Type()).Interface().(StringLike); ok {
            str = QuoteString(str)
        }

        err := jsoniter.ConfigFastest.Unmarshal([]byte(str), value)
        if err != nil {
            return fmt.Errorf("couldn't decode value from %v, error: %w", str, err)
        }
    }
    return nil
}

var separator = regexp.MustCompile(`" *, *"`)

func splitQuotedString(str string) []string {
    if !IsQuotedString(str, `"`) {
        return strings.Split(str, ",")
    }

    vals := separator.Split(str, -1)
    if len(vals) > 1 {
        vals[0] = vals[0] + `"`
        vals[len(vals)-1] = `"` + vals[len(vals)-1]

        for i := 1; i < len(vals)-1; i++ {
            vals[i] = QuoteString(vals[i])
        }
    }
    return vals
}
