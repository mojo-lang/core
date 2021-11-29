package core

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func ParseErrorCode(code string) (*ErrorCode, error) {
	ec := &ErrorCode{}
	err := ec.Parse(code)
	if err != nil {
		return nil, err
	}
	return ec, nil
}

func (m *ErrorCode) Parse(code string) error {
	if m != nil && len(code) > 0 {
		segments := strings.Split(code, ".")
		value := ""
		if len(segments) > 1 {
			m.Domain = segments[0]
			value = segments[1]
		} else {
			value = segments[0]
		}

		v, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("failed to parse error code %w", err)
		}

		m.Val = int32(v)
	}
	return nil
}

func (m *ErrorCode) Format() string {
	if m != nil {
		buffer := bytes.Buffer{}

		if len(m.Domain) > 0 {
			buffer.WriteString(m.Domain)
			buffer.WriteByte('.')
		}
		buffer.WriteString(strconv.FormatInt(int64(m.Val), 10))
		return buffer.String()
	}
	return ""
}
