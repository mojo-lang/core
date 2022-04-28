package core

import (
    "time"
)

const TimestampTypeName = "Timestamp"
const TimestampTypeFullName = "mojo.core.Timestamp"

func Now() *Timestamp {
    return FromTime(time.Now())
}

func FromTime(t time.Time) *Timestamp {
    seconds := t.Unix()
    return &Timestamp{
        Seconds:     seconds,
        Nanoseconds: int32(t.UnixNano() - seconds*1000000000),
    }
}

func Since(t *Timestamp) *Duration {
    if t != nil {
        return FromDuration(time.Since(t.ToTime()))
    }
    return nil
}

func Until(t *Timestamp) *Duration {
    if t != nil {
        return FromDuration(time.Until(t.ToTime()))
    }
    return nil
}

func (x *Timestamp) FromTime(t time.Time) {
    ts := FromTime(t)
    x.Seconds = ts.Seconds
    x.Nanoseconds = ts.Nanoseconds
}

func (x *Timestamp) ToTime() time.Time {
    return time.Unix(x.Seconds, int64(x.Nanoseconds)).In(loc)
}

func (x *Timestamp) After(u *Timestamp) bool {
    if x != nil && u != nil {
        if x != u {
            return x.ToTime().After(u.ToTime())
        }
    }
    return false
}

// Before reports whether the time instant t is before u.
func (x *Timestamp) Before(u *Timestamp) bool {
    if x != nil && u != nil {
        if x != u {
            return x.ToTime().Before(u.ToTime())
        }
    }
    return false
}

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
func (x *Timestamp) Equal(u *Timestamp) bool {
    if x != nil && u != nil {
        if x == u {
            return true
        }
        return x.ToTime().Equal(u.ToTime())
    }
    return false
}

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
func (x *Timestamp) Sub(u *Timestamp) *Duration {
    if x != nil && u != nil {
        return FromDuration(x.ToTime().Sub(u.ToTime()))
    }
    return nil
}

func (x *Timestamp) SinceNow() *Duration {
    return Since(x)
}

func (x *Timestamp) UntilNow() *Duration {
    return Until(x)
}
