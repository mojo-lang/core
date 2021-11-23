package core

var errorCodeIndex map[int32]*ErrorCode
var errorCodeNameIndex map[string]*ErrorCode

func init() {
	errorCodeIndex = make(map[int32]*ErrorCode)
	errorCodeNameIndex = make(map[string]*ErrorCode)

	addIndex := func(codes ...*ErrorCode) {
		for _, code := range codes {
			errorCodeIndex[code.Value] = code
			errorCodeNameIndex[code.Name] = code
		}
	}

	addIndex(Cancelled, UnknownError, InvalidArgument, DeadlineExceeded, NotFound,
		AlreadyExists, PermissionDenied, Unauthenticated, ResourceExhausted,
		FailedPrecondition, Aborted, OutOfRange, Unimplemented, InternalError,
		Unavailable, DataLoss)
}

var Cancelled = &ErrorCode{
	Value:          499,
	Name:           "CANCELLED",
	Description:    "The operation was cancelled, typically by the caller.",
	Document:       nil,
	HttpStatusCode: 499,
}

var UnknownError = &ErrorCode{
	Value:          2,
	Name:           "UNKNOWN_ERROR",
	Description:    "Unknown error.  For example, this error may be returned when a `Status` Value received from another address space belongs to an error space that is not known in this address space.",
	HttpStatusCode: 500,
}

var InvalidArgument = &ErrorCode{
	Value:          3,
	Name:           "INVALID_ARGUMENT",
	Description:    "The client specified an invalid argument.",
	HttpStatusCode: 400,
}

var DeadlineExceeded = &ErrorCode{
	Value:          504,
	Name:           "DEADLINE_EXCEEDED",
	Description:    "The deadline expired before the operation could complete.",
	HttpStatusCode: 504,
}

var NotFound = &ErrorCode{
	Value:          404,
	Name:           "NOT_FOUND",
	Description:    "Some requested entity (e.g., file or directory) was not found.",
	Document:       nil,
	HttpStatusCode: 404,
}

var AlreadyExists = &ErrorCode{
	Value:          6,
	Name:           "ALREADY_EXISTS",
	Description:    "The entity that a client attempted to create (e.g., file or directory) already exists.",
	HttpStatusCode: 409,
}

var PermissionDenied = &ErrorCode{
	Value:          403,
	Name:           "PERMISSION_DENIED",
	Description:    "The caller does not have permission to execute the specified operation.",
	HttpStatusCode: 403,
}

var Unauthenticated = &ErrorCode{
	Value:          401,
	Name:           "UNAUTHENTICATED",
	Description:    "The request does not have valid authentication credentials for the operation.",
	HttpStatusCode: 401,
}

var ResourceExhausted = &ErrorCode{
	Value:          429,
	Name:           "RESOURCE_EXHAUSTED",
	Description:    "Some resource has been exhausted, perhaps a per-user quota, or perhaps the entire file system is out of space.",
	HttpStatusCode: 429,
}

var FailedPrecondition = &ErrorCode{
	Value:          9,
	Name:           "FAILED_PRECONDITION",
	Description:    "The operation was rejected because the system is not in a state required for the operation's execution.",
	HttpStatusCode: 400,
}

var Aborted = &ErrorCode{
	Value:          10,
	Name:           "ABORTED",
	Description:    "The operation was aborted, typically due to a concurrency issue such as a sequencer check failure or transaction abort.",
	HttpStatusCode: 409,
}

var OutOfRange = &ErrorCode{
	Value:          11,
	Name:           "OUT_OF_RANGE",
	Description:    "The operation was attempted past the valid range.",
	HttpStatusCode: 400,
}

var Unimplemented = &ErrorCode{
	Value:          501,
	Name:           "UNIMPLEMENTED",
	Description:    "The operation is not implemented or is not supported/enabled in this service.",
	HttpStatusCode: 501,
}

var InternalError = &ErrorCode{
	Value:          500,
	Name:           "INTERNAL_ERROR",
	Description:    "Internal errors.  This means that some invariants expected by the underlying system have been broken.",
	HttpStatusCode: 500,
}

var Unavailable = &ErrorCode{
	Value:          503,
	Name:           "UNAVAILABLE",
	Description:    "The service is currently unavailable.",
	HttpStatusCode: 503,
}

var DataLoss = &ErrorCode{
	Value:          15,
	Name:           "DATA_LOSS",
	Description:    "Unrecoverable data loss or corruption.",
	HttpStatusCode: 500,
}
