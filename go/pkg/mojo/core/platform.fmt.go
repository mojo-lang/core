package core

import (
    "fmt"
    "path"
    "strings"
)

func ParsePlatform(value string) (*Platform, error) {
    p := &Platform{}
    if err := p.Parse(value); err != nil {
        return nil, err
    }
    return p, nil
}

func (x *Platform) Format() string {
    if x != nil {
        firstPart := ""
        if len(x.Variant) > 0 {
            firstPart = path.Join(x.Os.Format(), x.Architecture.Format(), x.Variant)
        } else if x.Architecture > 0 {
            firstPart = path.Join(x.Os.Format(), x.Architecture.Format())
        } else {
            firstPart = x.Os.Format()
        }
        if len(x.OsName) > 0 {
            if len(x.OsVersion) > 0 {
                return strings.Join([]string{firstPart, path.Join(x.OsName, x.OsVersion)}, "-")
            }
            return strings.Join([]string{firstPart, x.OsName}, "-")
        }
        return firstPart
    }
    return ""
}

func (x *Platform) Parse(value string) error {
    if x != nil && len(value) > 0 {
        parts := strings.Split(value, "-")
        if len(parts) > 0 {
            segments := strings.Split(parts[0], "/")
            if len(segments) > 0 {
                if err := x.Os.Parse(segments[0]); err != nil {
                    return fmt.Errorf("failed to parse platfrom in os (%s) part, error: %w", segments[0], err)
                }
            }
            if len(segments) > 1 {
                if err := x.Architecture.Parse(segments[1]); err != nil {
                    return fmt.Errorf("failed to parse platfrom in cpu (%s) part, error: %w", segments[1], err)
                }
            }
            if len(segments) == 3 {
                x.Variant = segments[2]
            }
        }
        if len(parts) > 1 {
            segments := strings.Split(parts[1], "/")
            if len(segments) > 0 {
                x.OsName = segments[0]
            }
            if len(segments) > 1 {
                x.OsVersion = segments[1]
            }
        }
    }
    return nil
}
