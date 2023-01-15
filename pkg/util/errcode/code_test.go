package errcode

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestCode_String(t *testing.T) {
	tests := []struct {
		name string
		c    Code
		want string
	}{
		{name: "unknown", c: CodeUnknown, want: "Unknown"},
		{name: "invalid argument", c: CodeInvalidArgument, want: "Invalid argument"},
		{name: "not found", c: CodeNotFound, want: "Not found"},
		{name: "forbidden", c: CodeForbidden, want: "Forbidden"},
		{name: "already exists", c: CodeAlreadyExists, want: "Already exists"},
		{name: "aborted", c: CodeAborted, want: "Aborted"},
		{name: "internal", c: CodeInternal, want: "Internal"},
		{name: "unavailable", c: CodeUnavailable, want: "Unavailable"},
		{name: "unimplemented", c: CodeUnimplemented, want: "Unimplemented"},
		{name: "cancelled", c: CodeCancelled, want: "Cancelled"},
		{name: "failed precondition", c: CodeFailedPrecondition, want: "Failed precondition"},
		{name: "error", c: -1, want: "Unknown: -1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCode_grpcCode(t *testing.T) {
	tests := []struct {
		name string
		c    Code
		want codes.Code
	}{
		{name: "unknown", c: CodeUnknown, want: codes.Unknown},
		{name: "invalid argument", c: CodeInvalidArgument, want: codes.InvalidArgument},
		{name: "not found", c: CodeNotFound, want: codes.NotFound},
		{name: "forbidden", c: CodeForbidden, want: codes.PermissionDenied},
		{name: "already exists", c: CodeAlreadyExists, want: codes.AlreadyExists},
		{name: "aborted", c: CodeAborted, want: codes.Aborted},
		{name: "internal", c: CodeInternal, want: codes.Internal},
		{name: "unavailable", c: CodeUnavailable, want: codes.Unavailable},
		{name: "unimplemented", c: CodeUnimplemented, want: codes.Unimplemented},
		{name: "cancelled", c: CodeCancelled, want: codes.Canceled},
		{name: "failed precondition", c: CodeFailedPrecondition, want: codes.FailedPrecondition},
		{name: "unknown", c: -1, want: codes.Unknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.grpcCode(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCode(t *testing.T) {
	tests := []struct {
		name string
		arg  error
		want Code
	}{
		{name: "nil", arg: nil, want: CodeUnknown},
		{name: "not found", arg: NewNotFound("not found"), want: CodeNotFound},
		{name: "unknown", arg: errors.New("error"), want: CodeUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCode(tt.arg)
			require.Equal(t, got, tt.want)
		})
	}
}
