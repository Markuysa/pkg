package errs

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

const (
	PermissionDenied = ErrorCode(codes.PermissionDenied)
	NotFound         = ErrorCode(codes.NotFound)
	InvalidArgument  = ErrorCode(codes.InvalidArgument)
	Internal         = ErrorCode(codes.Internal)
	Unauthenticated  = ErrorCode(codes.Unauthenticated)
)

var HttpMapper = map[ErrorCode]int{
	PermissionDenied: http.StatusForbidden,
	NotFound:         http.StatusNotFound,
	InvalidArgument:  http.StatusBadRequest,
	Internal:         http.StatusInternalServerError,
	Unauthenticated:  http.StatusUnauthorized,
}
