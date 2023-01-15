package errcode

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
)

type Code int

const (
	CodeUnknown Code = iota
	CodeInvalidArgument
	CodeNotFound
	CodeForbidden
	CodeAlreadyExists
	CodeAborted
	CodeInternal
	CodeUnavailable
	CodeUnimplemented
	CodeCancelled
	CodeResourceExhausted
	CodeFailedPrecondition
)

func (c Code) String() string {
	switch c {
	case CodeUnknown:
		return "Unknown"
	case CodeInvalidArgument:
		return "Invalid argument"
	case CodeNotFound:
		return "Not found"
	case CodeForbidden:
		return "Forbidden"
	case CodeAlreadyExists:
		return "Already exists"
	case CodeAborted:
		return "Aborted"
	case CodeInternal:
		return "Internal"
	case CodeUnavailable:
		return "Unavailable"
	case CodeUnimplemented:
		return "Unimplemented"
	case CodeCancelled:
		return "Cancelled"
	case CodeFailedPrecondition:
		return "Failed precondition"
	}
	return fmt.Sprintf("Unknown: %d", c)
}

func (c Code) grpcCode() codes.Code {
	switch c {
	case CodeUnknown:
		return codes.Unknown
	case CodeInvalidArgument:
		return codes.InvalidArgument
	case CodeNotFound:
		return codes.NotFound
	case CodeForbidden:
		return codes.PermissionDenied
	case CodeAlreadyExists:
		return codes.AlreadyExists
	case CodeAborted:
		return codes.Aborted
	case CodeInternal:
		return codes.Internal
	case CodeUnavailable:
		return codes.Unavailable
	case CodeUnimplemented:
		return codes.Unimplemented
	case CodeCancelled:
		return codes.Canceled
	case CodeFailedPrecondition:
		return codes.FailedPrecondition
	}
	return codes.Unknown
}

func NewCode(err error) Code {
	if err == nil {
		return CodeUnknown
	}
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return CodeUnknown
}
