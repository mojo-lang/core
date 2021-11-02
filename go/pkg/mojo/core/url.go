package core

import (
	"errors"
	"net/url"
	"strconv"
)

const UrlTypeName = "mojo.core.Url"

func NewUrl(url string) *Url {
	u, err := ParseUrl(url)
	if err != nil {
		return &Url{}
	}
	return u
}

func ParseUrl(raw string) (*Url, error) {
	url := &Url{}
	err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func NewQuery() *Url_Query {
	return &Url_Query{
		Values: make(map[string]*Strings),
	}
}

func (m *Url) Parse(raw string) error {
	if m == nil {
		return errors.New("")
	}

	url, err := url.Parse(raw)
	if err != nil {
		return err
	}

	m.Scheme = url.Scheme
	m.Authority = &Url_Authority{
		Host: url.Hostname(),
	}

	port := url.Port()
	if len(port) > 0 {
		v, err := strconv.Atoi(port)
		if err != nil {
			return err
		}
		m.Authority.Port = int64(v)
	}

	m.Path = url.Path
	m.Fragment = url.Fragment

	m.Query = &Url_Query{
		Values: make(map[string]*Strings),
	}
	query := url.Query()
	for k, v := range query {
		m.Query.Values[k] = &Strings{
			Values: v,
		}
	}

	return nil
}

func (m *Url) Format() string {
	if m == nil {
		return ""
	}

	host := ""
	if m.Authority != nil {
		host = m.Authority.Host
		if m.Authority.Port > 0 {
			host = host + ":" + strconv.FormatInt(m.Authority.Port, 10)
		}
	}

	u := url.URL{
		Scheme:   m.Scheme,
		Host:     host,
		Path:     m.Path,
		Fragment: m.Fragment,
	}

	if m.Query != nil {
		query := url.Values{}
		for k, v := range m.Query.Values {
			if v != nil {
				query[k] = v.Values
			}
		}
		u.RawQuery = query.Encode()
	}

	return u.String()
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
