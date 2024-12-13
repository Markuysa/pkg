package errs

import (
	"encoding/json"

	"google.golang.org/grpc/codes"
)

type (
	Error struct {
		msg string
		// Unique error code.
		code ErrorCode
	}
	ErrorCode codes.Code
)

var _ error = (*Error)(nil)

func New(msg string, code ErrorCode) *Error {
	return &Error{
		msg:  msg,
		code: code,
	}
}

func (e *Error) Error() string {
	res, _ := json.Marshal(e)

	return string(res)
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func (e *Error) Message() string {
	return e.msg
}
