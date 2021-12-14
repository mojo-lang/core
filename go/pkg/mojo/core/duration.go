package core

import (
	"math"
	"time"
)

const DurationTypeName = "mojo.core.Duration"

func (m Duration) ToDuration() time.Duration {
	return time.Duration(m.Seconds)*time.Second + time.Duration(m.Nanoseconds)
}

func FromDuration(d time.Duration) *Duration {
	duration := &Duration{}
	return duration.FromDuration(d)
}

func (m *Duration) FromDuration(d time.Duration) *Duration {
	if m != nil {
		m.Seconds = int64(d) / int64(time.Second)
		m.Nanoseconds = int32(int64(d) - m.Seconds*int64(time.Second))
	}
	return m
}

func (m *Duration) FromSeconds(seconds float64) *Duration {
	if m != nil {
		m.Seconds = int64(seconds)
		delta := seconds - float64(m.Seconds)
		m.Nanoseconds = int32(math.Round(delta * float64(time.Second)))
	}
	return m
}

func (m Duration) ToHours() float64 {
	return m.ToDuration().Hours()
}

func (m Duration) ToMinutes() float64 {
	return m.ToDuration().Minutes()
}

func (m Duration) ToSeconds() float64 {
	return m.ToDuration().Seconds()
}

func (m Duration) ToNanoseconds() int64 {
	return m.ToDuration().Nanoseconds()
}
