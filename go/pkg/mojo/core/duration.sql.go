package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

// Implement driver.Valuer and sql.Scanner interfaces on Duration
func (m *Duration) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return m.ToSeconds(), nil
}

func (m *Duration) Scan(src interface{}) error {
	v := reflect.ValueOf(src)
	if !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch duration := src.(type) {
	case float64:
		m.FromSeconds(duration)
		return nil
	default:
		return fmt.Errorf("could not not Decode type %T -> %T", src, m)
	}
}
