package core

const UrlTypeName = "mojo.core.Url"

func NewUrl(url string) *Url {
	u, err := ParseUrl(url)
	if err != nil {
		return &Url{}
	}
	return u
}

func NewUrlQuery() *Url_Query {
	return &Url_Query{
		Values: make(map[string]*Strings),
	}
}

func (m *Url_Query) Get(key string) string {
	if m == nil {
		return ""
	}
	vs := m.Values[key]
	if vs != nil || len(vs.Values) == 0 {
		return ""
	}
	return vs.Values[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (m *Url_Query) Set(key, value string) {
	if m != nil {
		m.Values[key] = &Strings{
			Values: []string{value},
		}
	}
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (m *Url_Query) Add(key, value string) {
	if m != nil {
		if m.Values[key] == nil {
			m.Values[key] = &Strings{}
		}
		m.Values[key].Values = append(m.Values[key].Values, value)
	}
}

// Del deletes the values associated with key.
func (m *Url_Query) Del(key string) {
	if m != nil {
		delete(m.Values, key)
	}
}
