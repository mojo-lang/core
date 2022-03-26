package core

import (
    "fmt"
    "math"
    "strconv"
)

func (m *Color) Format() string {
    if m != nil {
        buf := make([]byte, 0, 10)
        buf = append(buf, '#')

        buf = strconv.AppendUint(buf, uint64(m.Red), 16)
        buf = strconv.AppendUint(buf, uint64(m.Green), 16)
        buf = strconv.AppendUint(buf, uint64(m.Blue), 16)

        if m.Alpha != nil {
            alpha := uint64(math.Round(float64(m.GetAlphaValue() * math.MaxUint8)))
            buf = strconv.AppendUint(buf, alpha, 16)
        }
        return string(buf)
    }
    return ""
}

func (m *Color) Parse(value string) error {
    if m != nil && len(value) > 0 {
        if value[0] == '#' && len(value) >= 7 {
            if red, err := strconv.ParseInt(value[1:3], 16, 32); err != nil {
                return fmt.Errorf("failed to parse the color string: %s, error: %w", value, err)
            } else {
                m.Red = uint32(red)
            }

            if green, err := strconv.ParseInt(value[3:5], 16, 32); err != nil {
                return fmt.Errorf("failed to parse the color string: %s, error: %w", value, err)
            } else {
                m.Green = uint32(green)
            }

            if blue, err := strconv.ParseInt(value[5:7], 16, 32); err != nil {
                return fmt.Errorf("failed to parse the color string: %s, error: %w", value, err)
            } else {
                m.Blue = uint32(blue)
            }

            if len(value) == 9 {
                if alpha, err := strconv.ParseInt(value[7:9], 16, 32); err != nil {
                    return fmt.Errorf("failed to parse the color string: %s, error: %w", value, err)
                } else {
                    m.SetAlphaValue(float32(RoundNaive(float64(alpha)/255, 0.01)))
                }
            }
        } else {
            return fmt.Errorf("failed to parse the color string: %s", value)
        }
    }
    return nil
}
