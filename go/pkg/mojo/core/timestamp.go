package core

import (
    "time"
)

const TimestampTypeName = "Timestamp"
const TimestampTypeFullName = "mojo.core.Timestamp"

func Now() *Timestamp {
    return FromTime(time.Now())
}

func (m Timestamp) ToTime() time.Time {
    return time.Unix(m.Seconds, int64(m.Nanoseconds)).In(loc)
}

func FromTime(t time.Time) *Timestamp {
    seconds := t.Unix()
    return &Timestamp{
        Seconds:     seconds,
        Nanoseconds: int32(t.UnixNano() - seconds*1000000000),
    }
}

func (m *Timestamp) FromTime(t time.Time) {
    ts := FromTime(t)
    m.Seconds = ts.Seconds
    m.Nanoseconds = ts.Nanoseconds
}
