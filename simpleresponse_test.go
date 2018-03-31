package gorest

import (
	"bytes"
	"testing"
)

// TestNewSimpleResponse creates a SimpleResponse (with and without format) and
// verifies internal fields
func TestNewSimpleResponse(t *testing.T) {
	// Create without format.
	r := NewSimpleResponse("TAG", "XXX")
	if r.Status != "TAG" {
		t.Fatalf("Unexpected status. Expected: %s - Found: %s.", "TAG", r.Status)
	}
	if r.Message != "XXX" {
		t.Fatalf("Unexpected message. Expected: %s - Found: %s.", "XXX", r.Message)
	}
	// Create with format.
	r = NewSimpleResponsef("TAG2", "the%s", "message")
	if r.Status != "TAG2" {
		t.Fatalf("Unexpected status. Expected: %s - Found: %s.", "TAG2", r.Status)
	}
	if r.Message != "themessage" {
		t.Fatalf("Unexpected message. Expected: %s - Found: %s.", "themessage", r.Message)
	}
}

// TestNewFailResponse creates a SimpleResponse with NAK status by using
// NewFailResponse(f) functions and verifies internal fields
func TestNewFailResponse(t *testing.T) {
	// Create without format.
	r := NewFailResponse("XXX")
	if r.Status != NAK {
		t.Fatalf("Unexpected status. Expected: %s - Found: %s.", NAK, r.Status)
	}
	if r.Message != "XXX" {
		t.Fatalf("Unexpected message. Expected: %s - Found: %s.", "XXX", r.Message)
	}
	// Create with format.
	r = NewFailResponsef("the%s", "message")
	if r.Status != NAK {
		t.Fatalf("Unexpected status. Expected: %s - Found: %s.", NAK, r.Status)
	}
	if r.Message != "themessage" {
		t.Fatalf("Unexpected message. Expected: %s - Found: %s.", "themessage", r.Message)
	}
}

// TestSimpleResponseGetBody verifies the body of the SimpleResponse is
// properly created.
func TestSimpleResponseGetBody(t *testing.T) {
	r := NewSimpleResponse("X", "Y")
	body, err := r.GetBody()
	if err != nil {
		t.Fatalf("Unexpected error: %s.", err.Error())
	}

	expected := []byte("{\"status\":\"X\",\"message\":\"Y\"}")
	if bytes.Compare(body, expected) != 0 {
		t.Fatalf("Unexpected body. Expected: %s - Found: %s.", string(expected), string(body))
	}
}

// TestSimpleResponseGetCookieAndHeaders verifies that nil is returned by the
// GetCookie and GetHeaders functions.
func TestSimpleResponseGetCookieAndHeaders(t *testing.T) {
	r := NewSimpleResponse("A", "B")
	if cookie := r.GetCookie(); cookie != nil {
		t.Fatalf("Unexpected cookie found. Should be nil, found: %+v.", cookie)
	}
	if headers := r.GetHeaders(); headers != nil {
		t.Fatalf("Unexpected headers found. Should be nil, found: %+v.", headers)
	}
}
