package output

import (
	"context"
)

// ErrorOutputPort is the output port that returns the application error.
type ErrorOutputPort interface {
	// InternalServerError returns an 'Internal Server Error'.
	InternalServerError(ctx context.Context, err error) Error

	// BadRequest returns an error that the request is not valid.
	BadRequest(ctx context.Context, err error) Error
}

// Error is the application error.
type Error interface {
	// StatusCode returns the status code.
	StatusCode() int

	// Error returns the error message.
	Error() string
}
