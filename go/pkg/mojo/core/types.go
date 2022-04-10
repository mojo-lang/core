package core

import (
    "encoding/hex"
    "strconv"
)

type StringType interface {
    ToString() string
}

type ScalarType interface {
    ToScalar() interface{}
}

type ArrayType interface {
    ToArray() interface{}
}

type MapType interface {
    ToMap() interface{}
}

func ToString(value interface{}) string {
    switch v := value.(type) {
    case bool:
        return strconv.FormatBool(v)
    case int8, int16, int32, int64, int:
        return strconv.FormatInt(value.(int64), 10)
    case uint8, uint16, uint32, uint64, uint:
        return strconv.FormatUint(value.(uint64), 10)
    case float32, float64:
        return strconv.FormatFloat(value.(float64), 'e', 8, 10)
    case []byte:
        return hex.EncodeToString(v)
    case string:
        return v
    default:
        if s, ok := value.(StringType); ok {
            return s.ToString()
        } else if f, ok := value.(Formatter); ok {
            return f.Format()
        }
    }
    return ""
}
