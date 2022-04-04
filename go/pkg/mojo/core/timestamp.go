package core

import (
    "time"
)

const TimestampTypeName = "Timestamp"
const TimestampTypeFullName = "mojo.core.Timestamp"

func Now() *Timestamp {
    return FromTime(time.Now())
}

func (x Timestamp) ToTime() time.Time {
    return time.Unix(x.Seconds, int64(x.Nanoseconds)).In(loc)
}

func FromTime(t time.Time) *Timestamp {
    seconds := t.Unix()
    return &Timestamp{
        Seconds:     seconds,
        Nanoseconds: int32(t.UnixNano() - seconds*1000000000),
    }
}

func (x *Timestamp) FromTime(t time.Time) {
    ts := FromTime(t)
    x.Seconds = ts.Seconds
    x.Nanoseconds = ts.Nanoseconds
}
