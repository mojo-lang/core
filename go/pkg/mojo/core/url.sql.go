package core

import (
    "database/sql/driver"
    "fmt"
    "reflect"
)

// Value Implement driver.Valuer and sql.Scanner interfaces on Brand
func (x *Url) Value() (driver.Value, error) {
    if x == nil {
        return nil, nil
    }
    return x.Format(), nil
}

func (x *Url) Scan(src interface{}) error {
    v := reflect.ValueOf(src)
    if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
        return nil
    }

    switch bs := src.(type) {
    case string:
        x.Parse(bs)
        return nil
    default:
        return fmt.Errorf("failed to Decode type %T -> %T", src, x)
    }
}
