package output

// Error is the error in this application.
type Error interface {
	StatusCode() int
	Error() string
}
