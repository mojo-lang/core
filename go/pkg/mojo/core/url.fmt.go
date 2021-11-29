package core

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

func ParseUrl(raw string) (*Url, error) {
	url := &Url{}
	err := url.Parse(raw)
	if err != nil {
		return nil, err
	}
	return url, nil
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

	if len(m.Authority.Host) == 0 {
		url, err = url.Parse("http://" + raw)
		if err != nil {
			return err
		}
		m.Scheme = ""
		m.Authority.Host = url.Hostname()
		if len(m.Authority.Host) == 0 {
			return errors.New("failed to parse the url")
		}
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
		Vals: make(map[string]*Strings),
	}
	query := url.Query()
	for k, v := range query {
		m.Query.Vals[k] = &Strings{
			Vals: v,
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
		for k, v := range m.Query.Vals {
			if v != nil {
				query[k] = v.Vals
			}
		}
		u.RawQuery = query.Encode()
	}

	return u.String()
}

// get url like: "apis.company.com/path/to/resource"
func (m *Url) FormatWithoutSchema() string {
	if m == nil {
		return ""
	}

	url := &Url{
		Authority: &Url_Authority{
			Host: m.GetAuthority().GetHost(),
			Port: m.GetAuthority().GetPort(),
			//UserInfo: m.GetAuthority().GetUserInfo(),
		},
		Path:     m.Path,
		Query:    m.Query,
		Fragment: m.Fragment,
	}
	return strings.TrimPrefix(url.Format(), "//")
}
