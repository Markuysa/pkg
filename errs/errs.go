package errs

import (
	"encoding/json"

	"google.golang.org/grpc/codes"
)

type (
	Error struct {
		Msg string `json:"msg"`
		// Unique error Code.
		Code ErrorCode `json:"code"`
	}
	ErrorCode codes.Code
)

var _ error = (*Error)(nil)

func New(msg string, code ErrorCode) *Error {
	return &Error{
		Msg:  msg,
		Code: code,
	}
}

func (e *Error) Error() string {
	res, _ := json.Marshal(e)

	return string(res)
}

func (e *Error) GetCode() ErrorCode {
	return e.Code
}

func (e *Error) Message() string {
	return e.Msg
}
