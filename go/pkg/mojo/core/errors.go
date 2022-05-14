package core

import (
	"database/sql/driver"
	"errors"
	"google.golang.org/protobuf/reflect/protoreflect"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type BadRequestError struct {
	*err
}

func NewBadRequestError(format string, arguments ...interface{}) *BadRequestError {
	return &BadRequestError{newErr(BadRequest, format, arguments...)}
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
	*err
}

func NewInvalidArgumentError(format string, arguments ...interface{}) *InvalidArgumentError {
	return &InvalidArgumentError{newErr(InvalidArgument, format, arguments...)}
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
	*err
}

func NewMalformedSyntaxError(format string, arguments ...interface{}) *MalformedSyntaxError {
	return &MalformedSyntaxError{newErr(MalformedSyntax, format, arguments...)}
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
	*err
}

func NewFailedPreconditionError(format string, arguments ...interface{}) *FailedPreconditionError {
	return &FailedPreconditionError{newErr(FailedPrecondition, format, arguments...)}
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
	*err
}

func NewOutOfRangeError(format string, arguments ...interface{}) *OutOfRangeError {
	return &OutOfRangeError{newErr(OutOfRange, format, arguments...)}
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
	*err
}

func NewUnauthenticatedError(format string, arguments ...interface{}) *UnauthenticatedError {
	return &UnauthenticatedError{newErr(Unauthenticated, format, arguments...)}
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
	*err
}

func NewPermissionDeniedError(format string, arguments ...interface{}) *PermissionDeniedError {
	return &PermissionDeniedError{newErr(PermissionDenied, format, arguments...)}
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
	*err
}

func NewNotFoundError(format string, arguments ...interface{}) *NotFoundError {
	return &NotFoundError{newErr(NotFound, format, arguments...)}
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
	*err
}

func NewAlreadyExistsError(format string, arguments ...interface{}) *AlreadyExistsError {
	return &AlreadyExistsError{newErr(AlreadyExists, format, arguments...)}
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
	*err
}

func NewAbortedError(format string, arguments ...interface{}) *AbortedError {
	return &AbortedError{newErr(Aborted, format, arguments...)}
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
	*err
}

func NewResourceExhaustedError(format string, arguments ...interface{}) *ResourceExhaustedError {
	return &ResourceExhaustedError{newErr(ResourceExhausted, format, arguments...)}
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
	*err
}

func NewCancelledError(format string, arguments ...interface{}) *CancelledError {
	return &CancelledError{newErr(Cancelled, format, arguments...)}
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
	*err
}

func NewUnknownError(format string, arguments ...interface{}) *UnknownErrorError {
	return &UnknownErrorError{newErr(UnknownError, format, arguments...)}
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
	*err
}

func NewInternalError(format string, arguments ...interface{}) *InternalErrorError {
	return &InternalErrorError{newErr(InternalError, format, arguments...)}
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
	*err
}

func NewDataLossError(format string, arguments ...interface{}) *DataLossError {
	return &DataLossError{newErr(DataLoss, format, arguments...)}
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
	*err
}

func NewUnimplementedError(format string, arguments ...interface{}) *UnimplementedError {
	return &UnimplementedError{newErr(Unimplemented, format, arguments...)}
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
	*err
}

func NewUnavailableError(format string, arguments ...interface{}) *UnavailableError {
	return &UnavailableError{newErr(Unavailable, format, arguments...)}
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
	*err
}

func NewDeadlineExceededError(format string, arguments ...interface{}) *DeadlineExceededError {
	return &DeadlineExceededError{newErr(DeadlineExceeded, format, arguments...)}
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

func init() {
	RegisterJSONFieldEncoder("core.err", "Code", &ErrorCodeStringCodec{IsFieldPointer: true})
	RegisterJSONFieldDecoder("core.err", "Code", &ErrorCodeStringCodec{IsFieldPointer: true})
}

type err Error

func newErr(code *ErrorCode, message string, arguments ...interface{}) *err {
	return (*err)(NewError(code, message, arguments...))
}

func (e *err) Error() string {
	return (*Error)(e).Error()
}

func (e *err) ToError() *Error {
	return (*Error)(e)
}

func (e *err) StatusCode() int32 {
	return (*Error)(e).StatusCode()
}

func (e *err) AddDetail(detail interface{}) *err {
	return (*err)((*Error)(e).AddDetail(detail))
}

func (e *err) Value() (driver.Value, error) {
	if e != nil {
		return (*Error)(e).Value()
	}
	return nil, nil
}

func (e *err) Scan(src interface{}) error {
	if e != nil && src != nil {
		return (*Error)(e).Scan(src)
	}
	return nil
}

func (e *err) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return e.ToError().GormDBDataType(db, field)
}

func (e *err) GormDataType() string {
	return e.ToError().GormDataType()
}

func (e *err) Reset() {
	(*Error)(e).Reset()
}

func (e *err) String() string {
	return (*Error)(e).String()
}

func (*err) ProtoMessage() {}

func (e *err) ProtoReflect() protoreflect.Message {
	return (*Error)(e).ProtoReflect()
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (e *err) Descriptor() ([]byte, []int) {
	return (*Error)(e).Descriptor()
}

func (e *err) GetCode() *ErrorCode {
	return (*Error)(e).GetCode()
}

func (e *err) GetMessage() string {
	return (*Error)(e).GetMessage()
}

func (e *err) GetDetails() []*Any {
	return (*Error)(e).GetDetails()
}
