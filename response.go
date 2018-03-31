package gorest

import (
	"net/http"
)

// Response is the response structure that must be implemented and returned to
// the GoRest handler in order to let it configure the actual response.
type Response interface {
	GetBody() ([]byte, error)
	GetCookie() *http.Cookie
	GetHeaders() http.Header
}
