package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrNotFound is returned when the resource is not found
	ErrNotFound = New("not found").WithCode(http.StatusNotFound)
	// ErrInvalidContentType is returned when the content type is invalid
	ErrInvalidContentType = New("invalid content type").WithCode(http.StatusUnsupportedMediaType)
	// ErrMaxRequestSizeExceeded is returned when the request size exceeds the max request size
	ErrMaxRequestSizeExceeded = New("request size exceeds max request size").WithCode(http.StatusRequestEntityTooLarge)
	// ErrInvalidRequestMethod is returned when the request method is invalid
	ErrInvalidRequestMethod = New("invalid request method").WithCode(http.StatusMethodNotAllowed)
	// ErrFailedToParseMultipartForm is returned when the multipart form cannot be parsed
	ErrFailedToParseMultipartForm = New("failed to parse multipart form").WithCode(http.StatusBadRequest)
	// ErrFailedToParseForm is returned when the url-encoded form cannot be parsed
	ErrFailedToParseForm = New("failed to parse url-encoded form").WithCode(http.StatusBadRequest)
	// ErrFailedToParseJSON is returned when the json cannot be parsed
	ErrFailedToParseJSON = New("failed to parse json").WithCode(http.StatusBadRequest)
	// ErrLimitsExceeded is returned when the limits are exceeded
	ErrLimitsExceeded = New("limits exceeded").WithCode(http.StatusRequestEntityTooLarge)
	// ErrDuplicateFilename is returned when the filename is duplicated
	ErrDuplicateFilename = New("duplicate filename").WithCode(http.StatusConflict)
	// ErrFormIDNotProvided is returned when the form id is not provided
	ErrFormIDNotProvided = New("form id not provided").WithCode(http.StatusBadRequest)
	// ErrFormDisabled is returned when the form is disabled
	ErrFormDisabled = New("form is disabled").WithCode(http.StatusForbidden)
	// ErrFormDeleted is returned when the form is deleted
	ErrFormDeleted = New("form is deleted").WithCode(http.StatusForbidden)
	// ErrFormNotFound is returned when the form is not found
	ErrFormNotFound = New("form not found").WithCode(http.StatusNotFound)
)

// Wrap wraps an error with a message
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

// Error represents an error that can be returned from a handler
type Error struct {
	Message string
	Code    int
	Err     error
	Cause   error
}

// New creates a new error
func New(err string) *Error {
	return &Error{
		Message: err,
		Code:    http.StatusInternalServerError,
		Err:     errors.New(err),
	}
}

func FromErr(err error) *Error {
	return &Error{
		Message: err.Error(),
		Code:    http.StatusInternalServerError,
		Err:     err,
		Cause:   err,
	}
}

// WithCode sets the error code
func (e *Error) WithCode(code int) *Error {
	e.Code = code
	return e
}

// WithMessage sets the error code
func (e *Error) WithMessage(message string) *Error {
	e.Message = message
	return e
}

// WithCause sets the error cause
func (e *Error) WithCause(cause error) *Error {
	e.Cause = cause
	return e
}

// Error returns the error message
func (e Error) Error() string {
	return errors.Wrap(e.Err, e.Message).Error()
}

// Response returns the error as a response
func (e Error) Response() ErrorResponse {
	return ErrorResponse{
		Message: e.Message,
		Code:    e.Code,
	}
}
