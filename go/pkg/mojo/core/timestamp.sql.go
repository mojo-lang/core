package core

import (
    "database/sql/driver"
    "fmt"
    "reflect"
    "time"
)

// Implement driver.Valuer and sql.Scanner interfaces on Brand
func (x *Timestamp) Value() (driver.Value, error) {
    if x == nil {
        return nil, nil
    }
    return x.ToTime(), nil
}

func (x *Timestamp) Scan(src interface{}) error {
    v := reflect.ValueOf(src)
    if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
        return nil
    }

    switch bs := src.(type) {
    case time.Time:
        x.FromTime(bs)
        return nil
    default:
        return fmt.Errorf("Could not not Decode type %T -> %T", src, x)
    }
}
