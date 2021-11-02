package core

import "strings"

// 支持如下的格式
// field1,field2{filed2-1, field2-2}, field3
func NewFieldMask(fields string) *FieldMask {
	mask := &FieldMask{}
	mask.Parse(fields)
	return mask
}

func ParseFieldMask(fieldMask string) (*FieldMask, error) {
	mask := NewFieldMask(fieldMask)
	return mask, nil
}

func (m *FieldMask) Parse(field string) error {
	var paths []string
	paths = parseFields(field, paths)

	m.Paths = paths
	return nil
}

func (m *FieldMask) HasField(field string) bool {
	for _, f := range m.Paths {
		if f == field {
			return true
		}
	}
	return false
}

func parseFields(fields string, segments []string) []string {
	if len(fields) == 0 {
		return segments
	}

	comma := strings.Index(fields, ",")
	leftBrace := strings.Index(fields, "{")

	if comma < 0 {
		if leftBrace < 0 {
			segments = append(segments, strings.TrimSpace(fields))
		} else {
			rightBrace := strings.Index(fields, "}")
			if rightBrace > 0 {
				segments = append(segments, strings.TrimSpace(fields[:leftBrace])+"."+strings.TrimSpace(fields[leftBrace+1:rightBrace]))
			} else {
				segments = nil
			}
		}
	} else if leftBrace < 0 {
		list := strings.Split(fields, ",")
		for i := 0; i < len(list); i++ {
			if len(list[i]) > 0 {
				segments = append(segments, strings.TrimSpace(list[i]))
			}
		}
	} else {
		if leftBrace < comma {
			rightBrace := strings.Index(fields, "}")
			if rightBrace <= 0 {
				segments = nil
			} else {
				parent := strings.TrimSpace(fields[:leftBrace])
				brace := fields[leftBrace+1 : rightBrace]

				var subs []string
				subs = parseFields(brace, subs)
				if len(subs) == 0 {
					segments = append(segments, parent)
				} else {
					for i := 0; i < len(subs); i++ {
						segments = append(segments, parent+"."+strings.TrimSpace(subs[i]))
					}
				}
				segments = parseFields(fields[rightBrace+1:], segments)
			}
		} else {
			segments = append(segments, strings.TrimSpace(fields[:comma]))
			segments = parseFields(fields[comma+1:], segments)
		}
	}

	return segments
}
