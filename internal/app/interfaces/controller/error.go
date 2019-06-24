package controller

// Error is the error of this application.
type Error interface {
	Code() int
	Error() string
}

type apiError struct {
	code    int
	message string
}

// NewError returns a new error of this application.
func NewError(code int, msg string) Error {
	return &apiError{
		code:    code,
		message: msg,
	}
}

func (e *apiError) Code() int { return e.code }

func (e *apiError) Error() string { return e.message }
