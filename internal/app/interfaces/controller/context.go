package controller

import (
	"net/http"
)

// Context represents the context of the current HTTP request.
type Context interface {
	Request() *http.Request
	Param(name string) string
	Bind(i interface{}) error
	JSON(code int, i interface{}) error
}
