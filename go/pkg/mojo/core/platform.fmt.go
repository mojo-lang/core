package core

import (
    "fmt"
    "path"
    "strings"
)

func (m *Platform) Format() string {
    if m != nil {
        firstPart := ""
        if len(m.Variant) > 0 {
            firstPart = path.Join(m.Os.Format(), m.Architecture.Format(), m.Variant)
        } else {
            firstPart = path.Join(m.Os.Format(), m.Architecture.Format())
        }
        if len(m.OsName) > 0 {
            if len(m.OsVersion) > 0 {
                return strings.Join([]string{firstPart, path.Join(m.OsName, m.OsVersion)}, "-")
            }
            return strings.Join([]string{firstPart, m.OsName}, "-")
        }
        return firstPart
    }
    return ""
}

func (m *Platform) Parse(value string) error {
    if m != nil && len(value) > 0 {
        parts := strings.Split(value, "-")
        if len(parts) > 0 {
            segments := strings.Split(parts[0], "/")
            if len(segments) >= 2 {
                if err := m.Os.Parse(segments[0]); err != nil {
                    return fmt.Errorf("failed to parse platfrom in os (%s) part, error: %w", segments[0], err)
                }
                if err := m.Architecture.Parse(segments[1]); err != nil {
                    return fmt.Errorf("failed to parse platfrom in cpu (%s) part, error: %w", segments[1], err)
                }
            }
            if len(segments) == 3 {
                m.Variant = segments[2]
            }
        }
        if len(parts) > 1 {
            segments := strings.Split(parts[1], "/")
            if len(segments) > 0 {
                m.OsName = segments[0]
            }
            if len(segments) > 1 {
                m.OsVersion = segments[1]
            }
        }
    }
    return nil
}
