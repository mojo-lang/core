package core

import "errors"

type BadRequestError struct {
    *WrapError
}

func NewBadRequestError(format string, arguments ...interface{}) *BadRequestError {
    return &BadRequestError{NewWrapError(BadRequest, format, arguments...)}
}

func IsBadRequestError(err error) bool {
    return errors.Is(err, &BadRequestError{})
}

func (*BadRequestError) Is(err error) bool {
    if _, ok := err.(*BadRequestError); ok {
        return true
    }
    return false
}

type InvalidArgumentError struct {
    *WrapError
}

func NewInvalidArgumentError(format string, arguments ...interface{}) *InvalidArgumentError {
    return &InvalidArgumentError{NewWrapError(InvalidArgument, format, arguments...)}
}

func IsInvalidArgumentError(err error) bool {
    return errors.Is(err, &InvalidArgumentError{})
}

func (*InvalidArgumentError) Is(err error) bool {
    if _, ok := err.(*InvalidArgumentError); ok {
        return true
    }
    return false
}

type MalformedSyntaxError struct {
    *WrapError
}

func NewMalformedSyntaxError(format string, arguments ...interface{}) *MalformedSyntaxError {
    return &MalformedSyntaxError{NewWrapError(MalformedSyntax, format, arguments...)}
}

func IsMalformedSyntaxError(err error) bool {
    return errors.Is(err, &MalformedSyntaxError{})
}

func (*MalformedSyntaxError) Is(err error) bool {
    if _, ok := err.(*MalformedSyntaxError); ok {
        return true
    }
    return false
}

type FailedPreconditionError struct {
    *WrapError
}

func NewFailedPreconditionError(format string, arguments ...interface{}) *FailedPreconditionError {
    return &FailedPreconditionError{NewWrapError(FailedPrecondition, format, arguments...)}
}

func IsFailedPreconditionError(err error) bool {
    return errors.Is(err, &FailedPreconditionError{})
}

func (*FailedPreconditionError) Is(err error) bool {
    if _, ok := err.(*FailedPreconditionError); ok {
        return true
    }
    return false
}

type OutOfRangeError struct {
    *WrapError
}

func NewOutOfRangeError(format string, arguments ...interface{}) *OutOfRangeError {
    return &OutOfRangeError{NewWrapError(OutOfRange, format, arguments...)}
}

func IsOutOfRangeError(err error) bool {
    return errors.Is(err, &OutOfRangeError{})
}

func (*OutOfRangeError) Is(err error) bool {
    if _, ok := err.(*OutOfRangeError); ok {
        return true
    }
    return false
}

type UnauthenticatedError struct {
    *WrapError
}

func NewUnauthenticatedError(format string, arguments ...interface{}) *UnauthenticatedError {
    return &UnauthenticatedError{NewWrapError(Unauthenticated, format, arguments...)}
}

func IsUnauthenticatedError(err error) bool {
    return errors.Is(err, &UnauthenticatedError{})
}

func (*UnauthenticatedError) Is(err error) bool {
    if _, ok := err.(*UnauthenticatedError); ok {
        return true
    }
    return false
}

type PermissionDeniedError struct {
    *WrapError
}

func NewPermissionDeniedError(format string, arguments ...interface{}) *PermissionDeniedError {
    return &PermissionDeniedError{NewWrapError(PermissionDenied, format, arguments...)}
}

func IsPermissionDeniedError(err error) bool {
    return errors.Is(err, &PermissionDeniedError{})
}

func (*PermissionDeniedError) Is(err error) bool {
    if _, ok := err.(*PermissionDeniedError); ok {
        return true
    }
    return false
}

type NotFoundError struct {
    *WrapError
}

func NewNotFoundError(format string, arguments ...interface{}) *NotFoundError {
    return &NotFoundError{NewWrapError(NotFound, format, arguments...)}
}

func IsNotFoundError(err error) bool {
    return errors.Is(err, &NotFoundError{})
}

func (*NotFoundError) Is(err error) bool {
    if _, ok := err.(*NotFoundError); ok {
        return true
    }
    return false
}

type AlreadyExistsError struct {
    *WrapError
}

func NewAlreadyExistsError(format string, arguments ...interface{}) *AlreadyExistsError {
    return &AlreadyExistsError{NewWrapError(AlreadyExists, format, arguments...)}
}

func IsAlreadyExistsError(err error) bool {
    return errors.Is(err, &AlreadyExistsError{})
}

func (*AlreadyExistsError) Is(err error) bool {
    if _, ok := err.(*AlreadyExistsError); ok {
        return true
    }
    return false
}

type AbortedError struct {
    *WrapError
}

func NewAbortedError(format string, arguments ...interface{}) *AbortedError {
    return &AbortedError{NewWrapError(Aborted, format, arguments...)}
}

func IsAbortedError(err error) bool {
    return errors.Is(err, &AbortedError{})
}

func (*AbortedError) Is(err error) bool {
    if _, ok := err.(*AbortedError); ok {
        return true
    }
    return false
}

type ResourceExhaustedError struct {
    *WrapError
}

func NewResourceExhaustedError(format string, arguments ...interface{}) *ResourceExhaustedError {
    return &ResourceExhaustedError{NewWrapError(ResourceExhausted, format, arguments...)}
}

func IsResourceExhaustedError(err error) bool {
    return errors.Is(err, &ResourceExhaustedError{})
}

func (*ResourceExhaustedError) Is(err error) bool {
    if _, ok := err.(*ResourceExhaustedError); ok {
        return true
    }
    return false
}

type CancelledError struct {
    *WrapError
}

func NewCancelledError(format string, arguments ...interface{}) *CancelledError {
    return &CancelledError{NewWrapError(Cancelled, format, arguments...)}
}

func IsCancelledError(err error) bool {
    return errors.Is(err, &CancelledError{})
}

func (*CancelledError) Is(err error) bool {
    if _, ok := err.(*CancelledError); ok {
        return true
    }
    return false
}

type UnknownErrorError struct {
    *WrapError
}

func NewUnknownError(format string, arguments ...interface{}) *UnknownErrorError {
    return &UnknownErrorError{NewWrapError(UnknownError, format, arguments...)}
}

func IsUnknownError(err error) bool {
    return errors.Is(err, &UnknownErrorError{})
}

func (*UnknownErrorError) Is(err error) bool {
    if _, ok := err.(*UnknownErrorError); ok {
        return true
    }
    return false
}

type InternalErrorError struct {
    *WrapError
}

func NewInternalError(format string, arguments ...interface{}) *InternalErrorError {
    return &InternalErrorError{NewWrapError(InternalError, format, arguments...)}
}

func IsInternalError(err error) bool {
    return errors.Is(err, &InternalErrorError{})
}

func (*InternalErrorError) Is(err error) bool {
    if _, ok := err.(*InternalErrorError); ok {
        return true
    }
    return false
}

type DataLossError struct {
    *WrapError
}

func NewDataLossError(format string, arguments ...interface{}) *DataLossError {
    return &DataLossError{NewWrapError(DataLoss, format, arguments...)}
}

func IsDataLossError(err error) bool {
    return errors.Is(err, &DataLossError{})
}

func (*DataLossError) Is(err error) bool {
    if _, ok := err.(*DataLossError); ok {
        return true
    }
    return false
}

type UnimplementedError struct {
    *WrapError
}

func NewUnimplementedError(format string, arguments ...interface{}) *UnimplementedError {
    return &UnimplementedError{NewWrapError(Unimplemented, format, arguments...)}
}

func IsUnimplementedError(err error) bool {
    return errors.Is(err, &UnimplementedError{})
}

func (*UnimplementedError) Is(err error) bool {
    if _, ok := err.(*UnimplementedError); ok {
        return true
    }
    return false
}

type UnavailableError struct {
    *WrapError
}

func NewUnavailableError(format string, arguments ...interface{}) *UnavailableError {
    return &UnavailableError{NewWrapError(Unavailable, format, arguments...)}
}

func IsUnavailableError(err error) bool {
    return errors.Is(err, &UnavailableError{})
}

func (*UnavailableError) Is(err error) bool {
    if _, ok := err.(*UnavailableError); ok {
        return true
    }
    return false
}

type DeadlineExceededError struct {
    *WrapError
}

func NewDeadlineExceededError(format string, arguments ...interface{}) *DeadlineExceededError {
    return &DeadlineExceededError{NewWrapError(DeadlineExceeded, format, arguments...)}
}

func IsDeadlineExceededError(err error) bool {
    return errors.Is(err, &DeadlineExceededError{})
}

func (*DeadlineExceededError) Is(err error) bool {
    if _, ok := err.(*DeadlineExceededError); ok {
        return true
    }
    return false
}

type WrapError Error

func NewWrapError(code *ErrorCode, message string, arguments ...interface{}) *WrapError {
    return (*WrapError)(NewError(code, message, arguments...))
}

func (e *WrapError) Error() string {
    return (*Error)(e).Error()
}

func (e *WrapError) ToError() *Error {
    return (*Error)(e)
}

func (e *WrapError) StatusCode() int32 {
    return (*Error)(e).StatusCode()
}

func (e *WrapError) AddDetail(detail interface{}) *WrapError {
    (*Error)(e).AddDetail(detail)
    return e
}
