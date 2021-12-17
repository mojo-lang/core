package core

import (
	"strings"
)

func (m Ordering) Format() string {
	if len(m.Vals) > 0 {
		var values []string
		for _, v := range m.Vals {
			values = append(values, v.Format())
		}
		return strings.Join(values, ", ")
	}
	return ""
}

func ParseOrdering(value string) (*Ordering, error) {
	o := &Ordering{}
	if err := o.Parse(value); err != nil {
		return nil, err
	}
	return o, nil
}

func (m *Ordering) Parse(value string) error {
	if m != nil {
		segments := strings.Split(value, ",")
		for _, segment := range segments {
			v := &Ordering_Value{}
			if err := v.Parse(segment); err != nil {
				return err
			}
			m.Vals = append(m.Vals, v)
		}
	}
	return nil
}

func (m Ordering_Value) Format() string {
	if m.Sort != Ordering_SORT_UNSPECIFIED {
		return m.Field + " " + m.Sort.Format()
	} else {
		return m.Field
	}
}

func (m *Ordering_Value) Parse(value string) error {
	value = strings.TrimSpace(value)
	if len(value) > 0 {
		pos := strings.Index(value, " ")
		if pos < 0 {
			m.Field = value
		} else {
			m.Field = value[:pos]
			value = strings.TrimSpace(value[pos:])
			if err := m.Sort.Parse(strings.TrimSpace(value)); err != nil {
				return err
			}
		}
	}
	return nil
}
