package errcode

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	Code    Code
	origin  error
	stack   string
	callers []uintptr
}

func (e *Error) Error() string {
	if e.origin == nil {
		return fmt.Sprintf("%s: StackTrace:\n%s", e.Code.String(), e.stack)
	}
	return fmt.Sprintf("%s: %s\nStackTrace:\n%s", e.Code.String(), e.origin.Error(), e.stack)
}

func (e *Error) Stack() string {
	return e.stack
}

func (e *Error) Unwrap() error {
	return e.origin
}

func (e *Error) Callers() []uintptr {
	return e.callers
}

func New(err error) error {
	if err == nil {
		return nil
	}
	// if err is already Error type, nothing.
	var e *Error
	if errors.As(err, &e) {
		return err
	}

	const maxStackDepth, skipCallers = 30, 2
	stack := make([]uintptr, maxStackDepth)
	runtime.Callers(skipCallers, stack)
	newErr := &Error{
		Code:    CodeInternal,
		origin:  err,
		stack:   string(debug.Stack()),
		callers: stack,
	}

	// check context cancelled
	if errors.Is(err, context.Canceled) {
		newErr.Code = CodeCancelled
		return newErr
	}

	// check validation error
	var vErr validator.ValidationErrors
	if errors.As(err, &vErr) {
		newErr.Code = CodeInvalidArgument
		return newErr
	}

	// check grpc error
	// nolint:exhaustive
	switch status.Code(err) {
	case codes.InvalidArgument:
		newErr.Code = CodeInvalidArgument
	case codes.NotFound:
		newErr.Code = CodeNotFound
	case codes.Internal:
		newErr.Code = CodeInternal
	case codes.AlreadyExists:
		newErr.Code = CodeAlreadyExists
	case codes.Unavailable:
		newErr.Code = CodeUnavailable
	case codes.Aborted:
		newErr.Code = CodeAborted
	case codes.ResourceExhausted:
		newErr.Code = CodeResourceExhausted
	case codes.Canceled:
		newErr.Code = CodeCancelled
	}
	return newErr
}

func newWithCode(code Code, format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   code,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func NewGrpcError(err error) error {
	if err == nil {
		return nil
	}

	var e *Error
	if !errors.As(err, &e) {
		return status.Error(codes.Unknown, err.Error())
	}
	return status.Error(e.Code.grpcCode(), e.Error())
}

func NewNotFound(format string, a ...interface{}) error {
	return newWithCode(CodeNotFound, format, a...)
}

func NewInvalidArgument(format string, a ...interface{}) error {
	return newWithCode(CodeInvalidArgument, format, a...)
}

func NewAborted(format string, a ...interface{}) error {
	return newWithCode(CodeAborted, format, a...)
}

func NewInternal(format string, a ...interface{}) error {
	return newWithCode(CodeInternal, format, a...)
}

func NewUnimplemented(format string, a ...interface{}) error {
	return newWithCode(CodeUnimplemented, format, a...)
}

func NewResourceExhausted(format string, a ...interface{}) error {
	return newWithCode(CodeResourceExhausted, format, a...)
}

func NewUnknown(format string, a ...interface{}) error {
	return newWithCode(CodeUnknown, format, a...)
}

func NewFailedPrecondition(format string, a ...interface{}) error {
	return newWithCode(CodeFailedPrecondition, format, a...)
}

func isCode(err error, code Code) bool {
	if err == nil {
		return false
	}
	var e *Error
	if errors.As(err, &e) {
		return e.Code == code
	}
	return false
}

func IsAlreadyExists(err error) bool {
	return isCode(err, CodeAlreadyExists)
}

func IsNotfound(err error) bool {
	return isCode(err, CodeNotFound)
}

func IsInvalidArgument(err error) bool {
	return isCode(err, CodeInvalidArgument)
}

func IsInternal(err error) bool {
	return isCode(err, CodeInternal)
}

func IsUnknown(err error) bool {
	return isCode(err, CodeUnknown)
}

func IsAborted(err error) bool {
	return isCode(err, CodeAborted)
}

func IsCancelled(err error) bool {
	return isCode(err, CodeCancelled)
}

func IsServerError(err error) bool {
	return IsInternal(err) || IsUnknown(err) || IsAborted(err)
}
