package core

import (
    "fmt"
    "google.golang.org/protobuf/proto"
)

const ErrorTypeName = "Error"
const ErrorTypeFullName = "mojo.core.Error"

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

func (x *Error) Error() string {
    if len(x.Message) == 0 {
        return x.Code.Name
    }
    return x.Message
}

func (x *Error) StatusCode() int32 {
    if x != nil && x.Code != nil {
        if x.Code.HttpStatusCode > 0 {
            return x.Code.HttpStatusCode
        }
        return x.Code.Val
    }
    return 0
}

func (x *Error) AddDetail(detail interface{}) *Error {
    if _, ok := detail.(*proto.Message); ok {
    }
    return x
}
