
type TemplateString {
    type Segment {
        content: String @1
        templated: Bool @2
    }

    segments: [Segment] @1
}
