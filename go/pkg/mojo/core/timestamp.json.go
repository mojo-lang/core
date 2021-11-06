package core

import (
	"github.com/araddon/dateparse"
	jsoniter "github.com/json-iterator/go"
	"time"
	"unsafe"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func init() {
	jsoniter.RegisterTypeDecoder("core.Timestamp", &TimestampCodec{})
	jsoniter.RegisterTypeEncoder("core.Timestamp", &TimestampCodec{})
}

type TimestampCodec struct {
}

func (codec *TimestampCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	ts := (*Timestamp)(ptr)
	if any.ValueType() == jsoniter.NumberValue {
		number := any.ToInt64()

		if number < int64(MaxInt) {
			ts.Seconds = number
		} else {
			ts.Seconds = number / 1000
			ts.Nanoseconds = int32((number - ts.Seconds*1000) * 1000000)
		}
	} else if any.ValueType() == jsoniter.StringValue {
		str := any.ToString()

		// ts has timezone info, like "2006-01-02 15:04:05+0800"
		// since '+' will be replaced by space in url, we restore it to '+' if possible
		if len(str) > normalFormatLen && str[normalFormatLen] == ' ' {
			str = str[:normalFormatLen] + "+" + str[normalFormatLen+1:]
		}

		t, err := dateparse.ParseIn(str, loc)
		if err != nil {
		}

		ts.FromTime(t)
	}
}

func (codec *TimestampCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return ((*Timestamp)(ptr)).Seconds == 0
}

func (codec *TimestampCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ts := (*Timestamp)(ptr)
	stream.WriteString(ts.ToTime().Format(time.RFC3339))
}
