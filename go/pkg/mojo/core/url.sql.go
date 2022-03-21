package core

import (
    "database/sql/driver"
    "fmt"
    "reflect"
)

// Implement driver.Valuer and sql.Scanner interfaces on Brand
func (m *Url) Value() (driver.Value, error) {
    if m == nil {
        return nil, nil
    }
    return m.Format(), nil
}

func (m *Url) Scan(src interface{}) error {
    v := reflect.ValueOf(src)
    if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
        return nil
    }

    switch bs := src.(type) {
    case string:
        m.Parse(bs)
        return nil
    default:
        return fmt.Errorf("failed to Decode type %T -> %T", src, m)
    }
}
