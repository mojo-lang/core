package core

import "errors"

/// used in the loop callback
type BreakError struct {
}

func (e BreakError) Error() string {
	return "break"
}

func NewBreakError() error {
	return &BreakError{}
}

func IsBreakError(err error) bool {
	return errors.Is(err, &BreakError{})
}
