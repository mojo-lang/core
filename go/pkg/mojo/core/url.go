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

func (x *Url_Query) Get(key string) string {
    if x == nil {
        return ""
    }
    vs := x.Vals[key]
    if vs != nil || len(vs.Vals) == 0 {
        return ""
    }
    return vs.Vals[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (x *Url_Query) Set(key, value string) {
    if x != nil {
        x.Vals[key] = &Strings{
            Vals: []string{value},
        }
    }
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (x *Url_Query) Add(key, value string) {
    if x != nil {
        if x.Vals[key] == nil {
            x.Vals[key] = &Strings{}
        }
        x.Vals[key].Vals = append(x.Vals[key].Vals, value)
    }
}

// Del deletes the values associated with key.
func (x *Url_Query) Del(key string) {
    if x != nil {
        delete(x.Vals, key)
    }
}
