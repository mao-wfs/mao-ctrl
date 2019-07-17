package presenter

import (
	"context"
	"net/http"

	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/output"
)

// ErrorPresenter is the presenter for the application errors.
type ErrorPresenter struct{}

// NewErrorPresenter returns the new presenter fot the application error.
func NewErrorPresenter() output.ErrorOutputPort {
	return &ErrorPresenter{}
}

// InternalServerError returns an 'Internal Server Error'.
func (p *ErrorPresenter) InternalServerError(ctx context.Context, err error) output.Error {
	return &apiError{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}

// BadRequest returns an error that the request is not valid.
func (p *ErrorPresenter) BadRequest(ctx context.Context, err error) output.Error {
	return &apiError{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *apiError) StatusCode() int { return e.Code }

func (e *apiError) Error() string { return e.Message }
