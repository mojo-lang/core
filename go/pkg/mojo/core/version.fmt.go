package core

import (
    "bytes"
    "fmt"
    "strconv"
    "strings"
)

func ParseVersion(version string) (*Version, error) {
    v := &Version{}
    err := v.Parse(version)
    if err != nil {
        return nil, err
    }
    return v, nil
}

func (m *Version) Parse(version string) error {
    if m != nil && len(version) > 0 {
        segments := strings.Split(version, "+")
        if len(segments) > 1 {
            m.Builds = strings.Split(segments[1], ".")
        }

        segments = strings.Split(segments[0], "-")
        if len(segments) > 1 {
            segments[1] = strings.Join(segments[1:], "-")
            segments = segments[:2]
            m.PreReleases = strings.Split(segments[1], ".")
        }

        segments = strings.Split(segments[0], ".")
        m.Level = int32(len(segments))
        if m.Level > 0 {
            major, err := strconv.Atoi(segments[0])
            if err != nil {
                return fmt.Errorf("failed to parse version in major (%s) part, error: %w", segments[0], err)
            } else {
                m.Major = uint64(major)
            }
        }
        if m.Level > 1 {
            minor, err := strconv.Atoi(segments[1])
            if err != nil {
                return fmt.Errorf("failed to parse version in minor (%s) part, error: %w", segments[1], err)
            } else {
                m.Minor = uint64(minor)
            }
        }
        if m.Level > 2 {
            patch, err := strconv.Atoi(segments[2])
            if err != nil {
                return fmt.Errorf("failed to parse version in patch (%s) part, error: %w", segments[2], err)
            } else {
                m.Patch = uint64(patch)
            }
        }
    }

    return nil
}

func (m *Version) Format() string {
    if m != nil {
        buffer := bytes.Buffer{}
        buffer.WriteString(strconv.FormatUint(m.Major, 10))

        if m.Level == 0 || m.Level > 1 {
            buffer.WriteByte('.')
            buffer.WriteString(strconv.FormatUint(m.Minor, 10))
        }
        if m.Level == 0 || m.Level > 2 {
            buffer.WriteByte('.')
            buffer.WriteString(strconv.FormatUint(m.Patch, 10))
        }

        if len(m.PreReleases) > 0 {
            buffer.WriteByte('-')
            buffer.WriteString(strings.Join(m.PreReleases, "."))
        }

        if len(m.Builds) > 0 {
            buffer.WriteByte('+')
            buffer.WriteString(strings.Join(m.PreReleases, "."))
        }

        return buffer.String()
    }
    return ""
}
