package core

import (
	"errors"
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

func NewErrorFrom(code int32, message string, arguments ...interface{}) *Error {
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
	if len(arguments) > 0 {
		err.Message = fmt.Sprintf(message, arguments...)
	}
	return err
}

func IsError(err error) bool {
	return errors.Is(err, &Error{})
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
	if x != nil {
		if _, ok := detail.(proto.Message); ok {
			x.Details = append(x.Details, NewAny(detail))
		}
	}
	return x
}

func (*Error) Is(e error) bool {
	if _, ok := e.(*Error); ok {
		return true
	}
	if _, ok := e.(*err); ok {
		return true
	}
	return false
}
