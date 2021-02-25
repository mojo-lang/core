package core

func NewError(code int, message string) *Error {
	return &Error{
		Code: &ErrorCode{
			Value:       int64(code),
			Domain:      "",
			Brief:       "",
			Description: "",
		},
		Message: "",
	}
}

func (m *Error) Error() string {
	return ""
}

func (m *Error) StatusCode() int64 {
	if m != nil && m.Code != nil {
		return m.Code.Value
	}
	return 0
}
