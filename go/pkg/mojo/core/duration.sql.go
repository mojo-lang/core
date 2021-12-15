package core

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strconv"
)

// Implement driver.Valuer and sql.Scanner interfaces on Duration
func (m *Duration) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return m.ToSeconds(), nil
}

func (m *Duration) Scan(src interface{}) error {
	if v := reflect.ValueOf(src); !v.IsValid() || (v.CanAddr() && v.IsNil()) {
		return nil
	}

	switch duration := src.(type) {
	case float64:
		m.FromSeconds(duration)
	case string:
		seconds, err := strconv.ParseFloat(duration, 64)
		if err != nil {
			return err
		}
		m.FromSeconds(seconds)
	default:
		return fmt.Errorf("could not not Decode type %T -> %T", src, m)
	}
	return nil
}
