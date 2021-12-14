package core

import "time"

func (m Duration) Format() string {
	return m.ToDuration().String()
}

func ParseDuration(value string) (*Duration, error) {
	duration := &Duration{}
	if err := duration.Parse(value); err != nil {
		return nil, err
	}
	return duration, nil
}

func (m *Duration) Parse(value string) error {
	if m != nil {
		d, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		m.FromDuration(d)
	}
	return nil
}
