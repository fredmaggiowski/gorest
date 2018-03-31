package gorest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SimpleResponse is a standard response message useful for simple ACK or NAK.
type SimpleResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

const (
	// ACK message.
	ACK = "ACK"
	// NAK message.
	NAK = "NAK"
)

// NewSimpleResponse creates a new SimpleResponse with provided status
// and message.
func NewSimpleResponse(status, message string) SimpleResponse {
	return SimpleResponse{Status: status, Message: message}
}

// NewSimpleResponsef creates a new SimpleResponse with provided status and
// composing the message using provided format and variadic arguments.
func NewSimpleResponsef(status, format string, args ...interface{}) SimpleResponse {
	return SimpleResponse{Status: status, Message: fmt.Sprintf(format, args...)}
}

// NewFailResponse creates a new NAK SimpleResponse with provided message.
func NewFailResponse(message string) SimpleResponse {
	return NewSimpleResponse(NAK, message)
}

// NewFailResponsef creates a new NAK SimpleResponse composing the message
// using provided format and variadic arguments.
func NewFailResponsef(format string, args ...interface{}) SimpleResponse {
	return NewSimpleResponsef(NAK, format, args...)
}

// GetBody returns the JSON encoding of the FailResponse.
func (s SimpleResponse) GetBody() ([]byte, error) {
	content, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// GetCookie returns nil since no cookie is needed for this response.
func (s SimpleResponse) GetCookie() *http.Cookie {
	return nil
}

// GetHeaders returns nil since no header is needed for this response.
func (s SimpleResponse) GetHeaders() http.Header {
	return nil
}
