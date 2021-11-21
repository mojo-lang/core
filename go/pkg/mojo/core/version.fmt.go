package core

import (
	"bytes"
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
	if m != nil {
		if len(version) > 0 {
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
			if len(segments) == 3 {
				major, err := strconv.Atoi(segments[0])
				if err != nil {
					return err
				} else {
					m.Minor = uint64(major)
				}

				minor, err := strconv.Atoi(segments[1])
				if err != nil {
					return err
				} else {
					m.Minor = uint64(minor)
				}

				patch, err := strconv.Atoi(segments[2])
				if err != nil {
					return err
				} else {
					m.Patch = uint64(patch)
				}
			}
		}
	}

	return nil
}

func (m *Version) Format() string {
	if m != nil {
		buffer := bytes.Buffer{}
		buffer.WriteString(strconv.FormatUint(m.Major, 10))
		buffer.WriteByte('.')
		buffer.WriteString(strconv.FormatUint(m.Minor, 10))
		buffer.WriteByte('.')
		buffer.WriteString(strconv.FormatUint(m.Patch, 10))

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