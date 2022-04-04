package core

import (
    "math"
    "time"
)

const DurationTypeName = "Duration"
const DurationTypeFullName = "mojo.core.Duration"

func FromDuration(d time.Duration) *Duration {
    duration := &Duration{}
    return duration.FromDuration(d)
}

func NewDurationFromSeconds(seconds float64) *Duration {
    duration := &Duration{}
    return duration.FromSeconds(seconds)
}

func (x Duration) ToDuration() time.Duration {
    return time.Duration(x.Seconds)*time.Second + time.Duration(x.Nanoseconds)
}

func (x *Duration) FromDuration(d time.Duration) *Duration {
    if x != nil {
        x.Seconds = int64(d) / int64(time.Second)
        x.Nanoseconds = int32(int64(d) - x.Seconds*int64(time.Second))
    }
    return x
}

func (x *Duration) FromSeconds(seconds float64) *Duration {
    if x != nil {
        x.Seconds = int64(seconds)
        delta := seconds - float64(x.Seconds)
        x.Nanoseconds = int32(math.Round(delta * float64(time.Second)))
    }
    return x
}

func (x Duration) ToHours() float64 {
    return x.ToDuration().Hours()
}

func (x Duration) ToMinutes() float64 {
    return x.ToDuration().Minutes()
}

func (x Duration) ToSeconds() float64 {
    return x.ToDuration().Seconds()
}

func (x Duration) ToNanoseconds() int64 {
    return x.ToDuration().Nanoseconds()
}
