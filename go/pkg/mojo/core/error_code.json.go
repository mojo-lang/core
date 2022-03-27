package core

import (
    jsoniter "github.com/json-iterator/go"
    "unsafe"
)

// BareErrorCode will be jsonify to raw, without any codec
type BareErrorCode ErrorCode

type ErrorCodeStringCodec struct {
    IsFieldPointer bool
}

func (codec *ErrorCodeStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    value := iter.ReadString()
    errorCode := codec.errorCode(ptr)
    if errorCode == nil {
        errorCode = &ErrorCode{}
        *(**ErrorCode)(ptr) = errorCode
    }

    if err := errorCode.Parse(value); err != nil {
        iter.ReportError("ErrorCodeStringCodec", err.Error())
    }
}

func (codec *ErrorCodeStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
    return codec.errorCode(ptr) != nil
}

func (codec *ErrorCodeStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    errorCode := codec.errorCode(ptr)
    stream.WriteString(errorCode.Format())
}

func (codec *ErrorCodeStringCodec) errorCode(ptr unsafe.Pointer) *ErrorCode {
    if codec.IsFieldPointer {
        return *(**ErrorCode)(ptr)
    }
    return (*ErrorCode)(ptr)
}

type ErrorCodeStructCodec struct {
    IsFieldPointer bool
}

func (codec *ErrorCodeStructCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    errorCode := codec.bareErrorCode(ptr)
    if value := iter.ReadAny(); value.ValueType() == jsoniter.ObjectValue {
        if errorCode == nil {
            errorCode = &BareErrorCode{}
            *(**BareErrorCode)(ptr) = errorCode
        }
        value.ToVal(errorCode)
    }
}

func (codec *ErrorCodeStructCodec) IsEmpty(ptr unsafe.Pointer) bool {
    errorCode := codec.bareErrorCode(ptr)
    return errorCode == nil
}

func (codec *ErrorCodeStructCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    c := codec.bareErrorCode(ptr)
    stream.WriteVal(c)
}

func (codec *ErrorCodeStructCodec) bareErrorCode(ptr unsafe.Pointer) *BareErrorCode {
    if codec.IsFieldPointer {
        return *(**BareErrorCode)(ptr)
    }
    return (*BareErrorCode)(ptr)
}
