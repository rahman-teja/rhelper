package rhelper

import (
	"fmt"
	"net/http"
)

var (
	ErrMsgConflict       string = ErrMsgConflictFn("data")
	ErrMsgNotFound       string = "err_data_not_found"
	ErrMsgInternal       string = "err_internal_server"
	ErrMsgInvalid        string = "err_code_invalid"
	ErrMsgRequired       string = "err_data_required"
	ErrMsgExpired        string = "err_data_expired"
	ErrMsgInvalidPayload string = "err_code_invalid_payload"
	ErrMsgUnauthorized   string = "err_code_unauthorized"
)

var (
	ErrMsgConflictFn = func(s string) string {
		return fmt.Sprintf("err_%s_conflict", s)
	}
)

type ErrorBuild struct {
	HttpCode int
	Code     string
	Message  string
}

func (e ErrorBuild) Error() string {
	return fmt.Sprintf("[%d] %s | %s", e.HttpCode, e.Code, e.Message)
}

func NewErrorBuild(httpcode int, code, msg string) *ErrorBuild {
	return &ErrorBuild{
		HttpCode: httpcode,
		Code:     code,
		Message:  msg,
	}
}

var (
	ErrInternalServer error = ErrInternalServerFn("data")
	ErrNotFound       error = ErrNotFoundFn("data")
	ErrConflict       error = ErrConflictFn("data")
)

var (
	ErrDBNotPermissionDenied error = NewErrorBuild(http.StatusUnauthorized, ErrMsgUnauthorized, "DB Permission denied")
)

var (
	ErrConflictFn = func(s string) error {
		return NewErrorBuild(http.StatusConflict, ErrMsgConflict, s+" is duplicate")
	}
	ErrInternalServerFn = func(s string) error {
		return NewErrorBuild(http.StatusInternalServerError, ErrMsgInternal, fmt.Sprintf("an occured when processing %s", s))
	}
	ErrNotFoundFn = func(s string) error {
		return NewErrorBuild(http.StatusNotFound, ErrMsgNotFound, fmt.Sprintf("%s is not found", s))
	}
	ErrInvalidFn = func(s string) error {
		return NewErrorBuild(http.StatusUnprocessableEntity, ErrMsgInvalid, fmt.Sprintf("%s is invalid", s))
	}
	ErrRequiredFn = func(s string) error {
		return NewErrorBuild(http.StatusUnprocessableEntity, ErrMsgRequired, fmt.Sprintf("%s is required", s))
	}
	ErrExpiredFn = func(s string) error {
		return NewErrorBuild(http.StatusUnauthorized, ErrMsgExpired, fmt.Sprintf("%s is expired", s))
	}
)
