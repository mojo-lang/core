package core

type NotFoundError struct{}
func (e *NotFoundError) Error() string {
	return "Not Found"
}

type OutOfRangeError struct{}
func (e *OutOfRangeError) Error() string {
	return "Out of Range"
}

type InvalidArgumentError struct{}
func (e *InvalidArgumentError) Error() string {
	return "Invalid Argument"
}

type NotImplementedError struct{}
func (e *NotImplementedError) Error() string {
	return "Not Implemented"
}
