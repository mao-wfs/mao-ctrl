package presenter

import (
	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/output"
)

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewError returns a new error in this application.
func NewError(code int, msg string) output.Error {
	return &apiError{
		Code:    code,
		Message: msg,
	}
}

// StatusCode returns the status code.
func (e *apiError) StatusCode() int { return e.Code }

// Error returns the error message.
func (e *apiError) Error() string { return e.Message }
