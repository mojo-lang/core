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

func (x *Url) Parse(raw string) error {
    if x == nil {
        return errors.New("")
    }

    url, err := url.Parse(raw)
    if err != nil {
        return err
    }

    x.Scheme = url.Scheme
    x.Authority = &Url_Authority{
        Host: url.Hostname(),
    }

    if len(x.Authority.Host) == 0 {
        url, err = url.Parse("http://" + raw)
        if err != nil {
            return err
        }
        x.Scheme = ""
        x.Authority.Host = url.Hostname()
        if len(x.Authority.Host) == 0 {
            return errors.New("failed to parse the url")
        }
    }

    port := url.Port()
    if len(port) > 0 {
        v, err := strconv.Atoi(port)
        if err != nil {
            return err
        }
        x.Authority.Port = int64(v)
    }

    x.Path = url.Path
    x.Fragment = url.Fragment

    x.Query = &Url_Query{
        Vals: make(map[string]*StringValues),
    }
    query := url.Query()
    for k, v := range query {
        x.Query.Vals[k] = &StringValues{
            Vals: v,
        }
    }

    return nil
}

func (x *Url) Format() string {
    if x == nil {
        return ""
    }

    host := ""
    if x.Authority != nil {
        host = x.Authority.Host
        if x.Authority.Port > 0 {
            host = host + ":" + strconv.FormatInt(x.Authority.Port, 10)
        }
    }

    u := url.URL{
        Scheme:   x.Scheme,
        Host:     host,
        Path:     x.Path,
        Fragment: x.Fragment,
    }

    if x.Query != nil {
        query := url.Values{}
        for k, v := range x.Query.Vals {
            if v != nil {
                query[k] = v.Vals
            }
        }
        u.RawQuery = query.Encode()
    }

    return u.String()
}

// get url like: "apis.company.com/path/to/resource"
func (x *Url) FormatWithoutSchema() string {
    if x == nil {
        return ""
    }

    url := &Url{
        Authority: &Url_Authority{
            Host: x.GetAuthority().GetHost(),
            Port: x.GetAuthority().GetPort(),
            //UserInfo: x.GetAuthority().GetUserInfo(),
        },
        Path:     x.Path,
        Query:    x.Query,
        Fragment: x.Fragment,
    }
    return strings.TrimPrefix(url.Format(), "//")
}
