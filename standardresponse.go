package gorest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// StandardResponse is the response that usually will be used for response
// preparation, it features methods to set custom body, cookie and headers
// and implements the Response interface, if you do not need specific
// functionalities in the responses then this is the object you want to use.
type StandardResponse struct {
	body    []byte
	cookie  *http.Cookie
	headers http.Header
}

// NewStandardResponse creates a new empty Response.
func NewStandardResponse() *StandardResponse {
	return &StandardResponse{}
}

// SetJSONBody sets the body performing JSON encoding of the provided object.
func (r *StandardResponse) SetJSONBody(body interface{}) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed JSON conversion: %s", err.Error())
	}

	r.SetBody(bodyBytes)
	return nil
}

// SetBody can be used to set the the internal body providing the bytes
// allowing you to perform the preferred encoding.
func (r *StandardResponse) SetBody(bodyBytes []byte) {
	r.body = bodyBytes
}

// GetBody will be used by gorest core to retrieve the body to be written in
// the HTTP response writer.
func (r *StandardResponse) GetBody() ([]byte, error) {
	return r.body, nil
}

// SetCookie can be used to set a custom cookie that will be set in the
// HTTP response writer.
func (r *StandardResponse) SetCookie(cookie *http.Cookie) {
	r.cookie = cookie
}

// GetCookie will be used by gorest core to retrieve the cookie to be sent in
// the HTTP response writer.
func (r *StandardResponse) GetCookie() *http.Cookie {
	return r.cookie
}

// SetHeaders can be used to set a custom set of headers that will be set in
// the HTTP response writer.
func (r *StandardResponse) SetHeaders(headers http.Header) {
	r.headers = headers
}

// GetHeaders will be used by gorest core to retrieve the headers to be sent in
// the HTTP response writer.
func (r *StandardResponse) GetHeaders() http.Header {
	return r.headers
}
