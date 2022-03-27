package core

import (
    "bytes"
    "fmt"
    "strings"
)

const maxChecksumLength = 128 + 7

func (m *Checksum) Format() string {
    if !m.IsEmpty() {
        buf := bytes.NewBuffer(make([]byte, 0, maxChecksumLength))

        buf.WriteString(m.Algorithm.Format())
        buf.WriteByte(' ')
        buf.WriteString(m.Value)

        return buf.String()
    }

    return ""
}

func (m *Checksum) Parse(value string) error {
    if m != nil && len(value) > 0 {
        segments := strings.Split(value, " ")
        if len(segments) == 2 {
            if err := m.Algorithm.Parse(segments[0]); err != nil {
                return fmt.Errorf("failed to parse the checksum alogorithm: %s, error: %w", segments[0], err)
            }

            m.Value = segments[1]
            if !m.IsValid() {
                return fmt.Errorf("invalid checksum string: %s", value)
            }
        } else {
            return fmt.Errorf("failed to parse the checksum string: %s", value)
        }
    }
    return nil
}
