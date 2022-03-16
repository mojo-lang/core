package core

import "strings"

const TemplateStringTypeName = "TemplateString"
const TemplateStringTypeFullName = "mojo.core.TemplateString"

func NewTemplateString(str string) *TemplateString {
    ts := &TemplateString{}
    for len(str) > 0 {
        leftPos := strings.Index(str, "{")
        if leftPos >= 0 {
            rightPos := strings.Index(str[leftPos:], "}")
            if rightPos > 0 {
                rightPos = rightPos + leftPos
                if leftPos > 0 {
                    ts.AppendSegment(str[0:leftPos])
                }
                ts.AppendTemplatedSegment(str[leftPos+1 : rightPos])
                str = str[rightPos+1:]
            } else {
                ts.AppendSegment(str)
                break
            }
        } else {
            ts.AppendSegment(str)
            break
        }
    }
    return ts
}

func (m *TemplateString) AppendSegment(segment string) {
    if m != nil {
        m.Segments = append(m.Segments, &TemplateString_Segment{Content: segment})
    }
}

func (m *TemplateString) AppendTemplatedSegment(segment string) {
    if m != nil {
        m.Segments = append(m.Segments,
            &TemplateString_Segment{Content: segment, Templated: true},
        )
    }
}
