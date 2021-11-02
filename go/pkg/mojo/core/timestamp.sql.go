package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"
)

// Implement driver.Valuer and sql.Scanner interfaces on Brand
func (m *Timestamp) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return m.ToTime(), nil
}

func (m *Timestamp) Scan(src interface{}) error {
	v := reflect.ValueOf(src)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch bs := src.(type) {
	case time.Time:
		m.FromTime(bs)
		return nil
	default:
		return fmt.Errorf("Could not not Decode type %T -> %T", src, m)
	}
}
