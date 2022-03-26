package core

const StringValuesTypeName = "StringValues"
const StringValuesTypeFullName = "mojo.core.StringValues"

func NewStringValues(values ...string) *StringValues {
    return &StringValues{Vals: values}
}

func (m *StringValues) Unique() *StringValues {
    if m != nil {
        keys := make(map[string]bool)
        var list []string
        for _, entry := range m.Vals {
            if _, ok := keys[entry]; !ok {
                keys[entry] = true
                list = append(list, entry)
            }
        }
        m.Vals = list
    }
    return m
}
