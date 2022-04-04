package core

import (
    jsoniter "github.com/json-iterator/go"
    "unsafe"
)

func init() {
    RegisterJSONTypeDecoder("core.Error", &ErrorCodec{})
    RegisterJSONTypeEncoder("core.Error", &ErrorCodec{})
}

type ErrorCodec struct {
}

func (codec *ErrorCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    any := iter.ReadAny()
    err := (*Error)(ptr)
    if any.ValueType() == jsoniter.ObjectValue {
        code := any.Get("code").ToString()
        if len(code) > 0 {
            if ec, e := ParseErrorCode(code); e != nil {
                iter.ReportError("Decode ErrorCode", e.Error())
            } else {
                err.Code = ec
            }
        }

        err.Message = any.Get("message").ToString()
    }
}

func (codec *ErrorCodec) IsEmpty(ptr unsafe.Pointer) bool {
    err := (*Error)(ptr)
    return err == nil || err.Code == nil
}

func (codec *ErrorCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    err := (*Error)(ptr)

    stream.WriteObjectStart()

    stream.WriteObjectField("code")
    stream.WriteString(err.GetCode().Format())

    if len(err.Message) > 0 {
        stream.WriteMore()
        stream.WriteObjectField("message")
        stream.WriteString(err.Message)
    }

    if len(err.Details) > 0 {
    }
    stream.WriteObjectEnd()
}

func (x *Error) MarshalJSON() ([]byte, error) {
    return jsoniter.ConfigFastest.Marshal(x)
}

func (x *Error) UnmarshalJSON(err []byte) error {
    return jsoniter.ConfigFastest.Unmarshal(err, x)
}
