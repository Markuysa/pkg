package errs

import (
	"encoding/json"

	"google.golang.org/grpc/codes"
)

type (
	Error struct {
		Msg string `json:"msg"`
		// Code of error (internal, permission denied, etc).
		Code ErrorCode `json:"code"`
		// Index - must be a unique error identifier.
		Index int `json:"index"`
	}
	ErrorCode codes.Code
)

var _ error = (*Error)(nil)

func New(msg string, code ErrorCode, index int) *Error {
	return &Error{
		Msg:   msg,
		Code:  code,
		Index: index,
	}
}

func (e *Error) Error() string {
	res, _ := json.Marshal(e)

	return string(res)
}
