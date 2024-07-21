package model

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Code uint

const (
	CodeErrInternal        Code = 1000
	CodeErrTimeout         Code = 1001 // retryable
	CodeErrTooManyRequests Code = 1002 // retryable
	CodeErrUnavailable     Code = 1003 // retryable

	CodeErrValidation   Code = 2000
	CodeErrNotFound     Code = 2001 // retryable
	CodeErrUnauthorized Code = 2002
	CodeErrForbidden    Code = 2003
	CodeErrTooLarge     Code = 2004
)

//nolint:gochecknoglobals
var (
	codeToHTTPStatus = map[Code]int{
		CodeErrInternal:        fiber.StatusInternalServerError,
		CodeErrTimeout:         fiber.StatusRequestTimeout,
		CodeErrTooManyRequests: fiber.StatusTooManyRequests,
		CodeErrUnavailable:     fiber.StatusServiceUnavailable,

		CodeErrValidation:   fiber.StatusBadRequest,
		CodeErrNotFound:     fiber.StatusNotFound,
		CodeErrUnauthorized: fiber.StatusUnauthorized,
		CodeErrForbidden:    fiber.StatusForbidden,
		CodeErrTooLarge:     fiber.StatusRequestEntityTooLarge,
	}

	codeToMessage = map[Code]string{
		CodeErrInternal:        "Internal server error",
		CodeErrTimeout:         "Request timeout",
		CodeErrTooManyRequests: "Too many requests",
		CodeErrUnavailable:     "Service unavailable",

		CodeErrValidation:   "Validation error",
		CodeErrNotFound:     "Not found",
		CodeErrUnauthorized: "Unauthorized",
		CodeErrForbidden:    "Forbidden",
		CodeErrTooLarge:     "Request too large",
	}

	httpStatusToCode = map[int]Code{
		fiber.StatusInternalServerError: CodeErrInternal,
		fiber.StatusRequestTimeout:      CodeErrTimeout,
		fiber.StatusTooManyRequests:     CodeErrTooManyRequests,
		fiber.StatusServiceUnavailable:  CodeErrUnavailable,

		fiber.StatusBadRequest:            CodeErrValidation,
		fiber.StatusNotFound:              CodeErrNotFound,
		fiber.StatusUnauthorized:          CodeErrUnauthorized,
		fiber.StatusForbidden:             CodeErrForbidden,
		fiber.StatusRequestEntityTooLarge: CodeErrTooLarge,
	}

	protoStatusToCode = map[codes.Code]Code{
		codes.Canceled:           CodeErrTimeout,
		codes.InvalidArgument:    CodeErrValidation,
		codes.DeadlineExceeded:   CodeErrTimeout,
		codes.NotFound:           CodeErrNotFound,
		codes.AlreadyExists:      CodeErrValidation,
		codes.PermissionDenied:   CodeErrForbidden,
		codes.ResourceExhausted:  CodeErrTooManyRequests,
		codes.FailedPrecondition: CodeErrValidation,
		codes.Unimplemented:      CodeErrUnavailable,
		codes.Internal:           CodeErrInternal,
		codes.Unavailable:        CodeErrUnavailable,
		codes.Unauthenticated:    CodeErrUnauthorized,
	}
)

func FromGRPCError(status *status.Status) Code {
	if code, found := protoStatusToCode[status.Code()]; found {
		return code
	}

	return CodeErrInternal
}

func FromFiberError(err *fiber.Error) Code {
	if code, found := httpStatusToCode[err.Code]; found {
		return code
	}

	return CodeErrInternal
}

func ToHTTPStatus(code Code) int {
	if httpStatus, found := codeToHTTPStatus[code]; found {
		return httpStatus
	}

	return fiber.StatusInternalServerError
}

func IsRetryableCode(code Code) bool {
	switch code {
	case CodeErrTooManyRequests, CodeErrTimeout, CodeErrNotFound:
		return true
	case CodeErrInternal, CodeErrUnavailable, CodeErrValidation, CodeErrUnauthorized, CodeErrForbidden, CodeErrTooLarge:
		return false
	default:
		return false
	}
}

func IsClientError(code Code) bool {
	return code >= CodeErrValidation
}

func IsServerError(code Code) bool {
	return code < CodeErrValidation
}
