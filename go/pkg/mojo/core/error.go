package core

import (
    "fmt"
    "github.com/golang/protobuf/proto"
)

const ErrorTypeName = "Error"
const ErrorTypeFullName = "mojo.core.Error"

func init() {
    RegisterJSONFieldEncoder("core.Error", "Code", &ErrorCodeStringCodec{IsFieldPointer: true})
    RegisterJSONFieldDecoder("core.Error", "Code", &ErrorCodeStringCodec{IsFieldPointer: true})
}

func NewError(code *ErrorCode, message string, arguments ...interface{}) *Error {
    msg := message
    if len(arguments) > 0 {
        msg = fmt.Sprintf(message, arguments...)
    }

    return &Error{
        Code:    code,
        Message: msg,
    }
}

func NewErrorFrom(code int32, message string) *Error {
    err := &Error{}
    if ec, ok := errorCodeIndex[code]; ok {
        err.Code = ec
    } else {
        if code >= 100 && code < 600 {
            err.Code = &ErrorCode{
                Val: code,
            }
        } else {
            return nil
        }
    }
    err.Message = message
    return err
}

func NewCancelledError(format string, arguments ...interface{}) *Error {
    return NewError(Cancelled, format, arguments...)
}

func NewUnknownError(format string, arguments ...interface{}) *Error {
    return NewError(UnknownError, format, arguments...)
}

func NewInvalidArgumentError(format string, arguments ...interface{}) *Error {
    return NewError(InvalidArgument, format, arguments...)
}

func NewDeadlineExceededError(format string, arguments ...interface{}) *Error {
    return NewError(DeadlineExceeded, format, arguments...)
}

func NewNotFoundError(format string, arguments ...interface{}) *Error {
    return NewError(NotFound, format, arguments...)
}

func NewAlreadyExistsError(format string, arguments ...interface{}) *Error {
    return NewError(AlreadyExists, format, arguments...)
}

func NewPermissionDeniedError(format string, arguments ...interface{}) *Error {
    return NewError(PermissionDenied, format, arguments...)
}

func NewUnauthenticatedError(format string, arguments ...interface{}) *Error {
    return NewError(Unauthenticated, format, arguments...)
}

func NewResourceExhaustedError(format string, arguments ...interface{}) *Error {
    return NewError(ResourceExhausted, format, arguments...)
}

func NewFailedPreconditionError(format string, arguments ...interface{}) *Error {
    return NewError(FailedPrecondition, format, arguments...)
}

func NewAbortedError(format string, arguments ...interface{}) *Error {
    return NewError(Aborted, format, arguments...)
}

func NewOutOfRangeError(format string, arguments ...interface{}) *Error {
    return NewError(OutOfRange, format, arguments...)
}

func NewUnimplementedError(format string, arguments ...interface{}) *Error {
    return NewError(Unimplemented, format, arguments...)
}

func NewInternalError(format string, arguments ...interface{}) *Error {
    return NewError(InternalError, format, arguments...)
}

func NewUnavailableError(format string, arguments ...interface{}) *Error {
    return NewError(Unavailable, format, arguments...)
}

func NewDataLossError(format string, arguments ...interface{}) *Error {
    return NewError(DataLoss, format, arguments...)
}

func (m *Error) Error() string {
    if len(m.Message) == 0 {
        return m.Code.Name
    }
    return m.Message
}

func (m *Error) StatusCode() int32 {
    if m != nil && m.Code != nil {
        if m.Code.HttpStatusCode > 0 {
            return m.Code.HttpStatusCode
        }
        return m.Code.Val
    }
    return 0
}

func (m *Error) AddDetail(detail interface{}) *Error {
    if _, ok := detail.(*proto.Message); ok {
    }
    return m
}
