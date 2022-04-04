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

func (x *TemplateString) AppendSegment(segment string) {
    if x != nil {
        x.Segments = append(x.Segments, &TemplateString_Segment{Content: segment})
    }
}

func (x *TemplateString) AppendTemplatedSegment(segment string) {
    if x != nil {
        x.Segments = append(x.Segments,
            &TemplateString_Segment{Content: segment, Templated: true},
        )
    }
}
