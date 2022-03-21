package core

const TemplateStringTypeName = "TemplateString"
const TemplateStringTypeFullName = "mojo.core.TemplateString"

func NewTemplateString(str string) *TemplateString {
    ts := &TemplateString{}
    if err := ts.Parse(str); err != nil {
        return nil
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
