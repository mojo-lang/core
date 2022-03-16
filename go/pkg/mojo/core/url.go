package core

const UrlTypeName = "Url"
const UrlTypeFullName = "mojo.core.Url"

func NewUrl(url string) *Url {
    u, err := ParseUrl(url)
    if err != nil {
        return &Url{}
    }
    return u
}

func NewUrlQuery() *Url_Query {
    return &Url_Query{
        Vals: make(map[string]*Strings),
    }
}

func (m *Url_Query) Get(key string) string {
    if m == nil {
        return ""
    }
    vs := m.Vals[key]
    if vs != nil || len(vs.Vals) == 0 {
        return ""
    }
    return vs.Vals[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (m *Url_Query) Set(key, value string) {
    if m != nil {
        m.Vals[key] = &Strings{
            Vals: []string{value},
        }
    }
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (m *Url_Query) Add(key, value string) {
    if m != nil {
        if m.Vals[key] == nil {
            m.Vals[key] = &Strings{}
        }
        m.Vals[key].Vals = append(m.Vals[key].Vals, value)
    }
}

// Del deletes the values associated with key.
func (m *Url_Query) Del(key string) {
    if m != nil {
        delete(m.Vals, key)
    }
}
