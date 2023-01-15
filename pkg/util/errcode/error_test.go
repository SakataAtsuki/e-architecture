package errcode

import (
	"errors"
	"runtime"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestError_Error(t *testing.T) {
	// no origin
	err := &Error{
		Code:   1,
		origin: nil,
		stack:  "stack",
	}
	got := err.Error()
	require.Equal(t, "Invalid argument: StackTrace:\nstack", got)
	require.Equal(t, "stack", err.Stack())

	// has origin
	err = &Error{
		Code:   1,
		origin: errors.New("error"),
		stack:  "stack",
	}
	got = err.Error()
	require.Equal(t, "Invalid argument: error\nStackTrace:\nstack", got)
	require.Equal(t, "error", err.Unwrap().Error())
}

func TestError_Callers(t *testing.T) {
	stack := make([]uintptr, 10)
	runtime.Callers(0, stack)
	err := &Error{
		callers: stack,
	}
	require.Equal(t, stack, err.Callers())
	require.Equal(t, 10, len(stack))
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		arg  error
		want error
	}{
		{name: "nil", arg: nil, want: nil},
		{name: "Error", arg: &Error{}, want: &Error{}},
		{
			name: "default internal",
			arg:  errors.New("error"),
			want: &Error{
				Code:   CodeInternal,
				origin: errors.New("error"),
			},
		},
		{
			name: "validation error",
			arg:  validator.ValidationErrors{},
			want: &Error{
				Code:   CodeInvalidArgument,
				origin: validator.ValidationErrors{},
			},
		},
		{
			name: "grpc invalid argument",
			arg:  status.Error(codes.InvalidArgument, "error"),
			want: &Error{
				Code:   CodeInvalidArgument,
				origin: status.Error(codes.InvalidArgument, "error"),
			},
		},
		{
			name: "grpc not found",
			arg:  status.Error(codes.NotFound, "error"),
			want: &Error{
				Code:   CodeNotFound,
				origin: status.Error(codes.NotFound, "error"),
			},
		},
		{
			name: "grpc internal",
			arg:  status.Error(codes.Internal, "error"),
			want: &Error{
				Code:   CodeInternal,
				origin: status.Error(codes.Internal, "error"),
			},
		},
		{
			name: "grpc already exists",
			arg:  status.Error(codes.AlreadyExists, "error"),
			want: &Error{
				Code:   CodeAlreadyExists,
				origin: status.Error(codes.AlreadyExists, "error"),
			},
		},
		{
			name: "grpc unavailable",
			arg:  status.Error(codes.Unavailable, "error"),
			want: &Error{
				Code:   CodeUnavailable,
				origin: status.Error(codes.Unavailable, "error"),
			},
		},
		{
			name: "grpc aborted",
			arg:  status.Error(codes.Aborted, "error"),
			want: &Error{
				Code:   CodeAborted,
				origin: status.Error(codes.Aborted, "error"),
			},
		},
		{
			name: "grpc canceled",
			arg:  status.Error(codes.Canceled, "error"),
			want: &Error{
				Code:   CodeCancelled,
				origin: status.Error(codes.Canceled, "error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := New(tt.arg)
			if tt.want != nil {
				want, got := tt.want.(*Error), gotErr.(*Error)
				require.Equal(t, want.Code, got.Code)
				require.Equal(t, want.origin, got.origin)
			} else {
				require.Equal(t, tt.want, gotErr)
			}
		})
	}
}

func TestGrpcNew(t *testing.T) {
	tests := []struct {
		name string
		arg  error
		want error
	}{
		{name: "nil", arg: nil, want: nil},
		{name: "Unknown error", arg: errors.New("error"), want: status.Error(codes.Unknown, "error")},
		{
			name: "Known error",
			arg:  &Error{Code: CodeInvalidArgument},
			want: status.Error(codes.InvalidArgument, "Invalid argument: StackTrace:\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := NewGrpcError(tt.arg)
			require.Equal(t, tt.want, gotErr)
		})
	}
}

func TestNewError(t *testing.T) {
	tests := []struct {
		name    string
		newFunc func(format string, a ...interface{}) error
		want    Code
	}{
		{name: "not found", newFunc: NewNotFound, want: CodeNotFound},
		{name: "aborted", newFunc: NewAborted, want: CodeAborted},
		{name: "invalid argument", newFunc: NewInvalidArgument, want: CodeInvalidArgument},
		{name: "unimplemented", newFunc: NewUnimplemented, want: CodeUnimplemented},
		{name: "failed precondition", newFunc: NewFailedPrecondition, want: CodeFailedPrecondition},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.newFunc("foo")
			require.Equal(t, tt.want, NewCode(got))
		})
	}
}

func TestIsError(t *testing.T) {
	tests := []struct {
		name   string
		arg    error
		isFunc func(error) bool
		want   bool
	}{
		{name: "nil", arg: nil, isFunc: IsNotfound, want: false},
		{name: "not found true", arg: NewNotFound("foo"), isFunc: IsNotfound, want: true},
		{name: "not found false", arg: NewAborted("foo"), isFunc: IsNotfound, want: false},
		{name: "internal", arg: NewInternal("foo"), isFunc: IsInternal, want: true},
		{name: "unknown", arg: NewUnknown("foo"), isFunc: IsUnknown, want: true},
		{name: "invalid argument", arg: NewInvalidArgument("foo"), isFunc: IsInvalidArgument, want: true},
		{name: "server error: internal", arg: NewInternal("foo"), isFunc: IsServerError, want: true},
		{name: "server error: unknown", arg: NewUnknown("foo"), isFunc: IsServerError, want: true},
		{name: "aborted", arg: NewAborted("foo"), isFunc: IsAborted, want: true},
		{name: "cancelled", arg: &Error{Code: CodeCancelled}, isFunc: IsCancelled, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.isFunc(tt.arg)
			require.Equal(t, tt.want, got)
		})
	}
}
