package gorest

import (
	"net/http"
	"testing"
)

// TestStandardResponseSetJSONBody verifies the SetJSONBody works as expected.
func TestStandardResponseSetJSONBody(t *testing.T) {
	type A struct {
		B string `json:"theB"`
	}

	r := NewStandardResponse()
	if err := r.SetJSONBody(A{"hej"}); err != nil {
		t.Fatalf("Unexpected error: %s.", err.Error())
	}

	expected := `{"theB":"hej"}`
	if string(r.body) != expected {
		t.Fatalf("Unexecpted body found. Expected: %s - Found: %s.", expected, string(r.body))
	}

	if err := r.SetJSONBody(make(chan int)); err == nil {
		t.Fatalf("An error was expected. Found nil.")
	}
}

// TestStandardResponseSetBody verifies the SetBody works as expected.
func TestStandardResponseSetBody(t *testing.T) {
	r := NewStandardResponse()
	body := []byte("hey monika")
	r.SetBody(body)

	if string(r.body) != string(body) {
		t.Fatalf("Unexecpted body found. Expected: %s - Found: %s.", string(body), string(r.body))
	}
}

// TestStandardResponseGetBody sets manually the body and verifies the GetBody
// method returns it properly.
func TestStandardResponseGetBody(t *testing.T) {
	r := NewStandardResponse()
	if r.GetBody() != nil {
		t.Fatalf("Body should be nil, found: %+v.", r.GetBody())
	}

	body := []byte("hey monika")
	r.body = body
	if string(r.GetBody()) != string(body) {
		t.Fatalf("Unexecpted body found. Expected: %s - Found: %s.", string(body), string(r.GetBody()))
	}
}

// TestStandardResponseSetCookie verifies that SetCookie works as expected.
func TestStandardResponseSetCookie(t *testing.T) {
	r := NewStandardResponse()
	cookie := http.Cookie{
		Domain: "text.example.com",
	}
	r.SetCookie(&cookie)

	if r.cookie == nil {
		t.Fatalf("Unexpected nil cookie.")
	}

	if r.cookie.Domain != cookie.Domain {
		t.Fatalf("Unexpected cookie domain. Expected: %s - Found: %s.", cookie.Domain, r.cookie.Domain)
	}
}

// TestStandardResponseGetCookie sets manually the cookie and verifies the
// GetCookie method returns it properly.
func TestStandardResponseGetCookie(t *testing.T) {
	r := NewStandardResponse()
	if r.GetCookie() != nil {
		t.Fatalf("Cookie should be nil, found: %+v.", r.GetCookie())
	}

	cookie := http.Cookie{
		Domain: "text.example.com",
	}
	r.cookie = &cookie
	if r.GetCookie() == nil {
		t.Fatalf("Unexpected nil cookie.")
	}

	if r.GetCookie().Domain != cookie.Domain {
		t.Fatalf("Unexpected cookie domain. Expected: %s - Found: %s.", cookie.Domain, r.GetCookie().Domain)
	}
}

// TestStandardResponseSetHeaders verifies that SetHeaders works as expected.
func TestStandardResponseSetHeaders(t *testing.T) {
	r := NewStandardResponse()
	headers := http.Header{}
	headers.Set("A", "B")
	r.SetHeaders(headers)

	if r.headers == nil {
		t.Fatalf("Unexpected nil headers.")
	}

	if r.headers.Get("A") != "B" {
		t.Fatalf("Unexpected header \"A\" value. Expected: %s - Found: %s.", "B", r.headers.Get("A"))
	}
}

// TestStandardResponseGetHeaders sets manually the headers and verifies the
// GetHeaders method returns it properly.
func TestStandardResponseGetHeaders(t *testing.T) {
	r := NewStandardResponse()
	if r.GetHeaders() != nil {
		t.Fatalf("Headers should be nil, found: %+v.", r.GetHeaders())
	}

	headers := http.Header{}
	headers.Set("A", "B")
	r.headers = headers

	if r.GetHeaders() == nil {
		t.Fatalf("Unexpected nil headers.")
	}

	if r.GetHeaders().Get("A") != "B" {
		t.Fatalf("Unexpected header \"A\" value. Expected: %s - Found: %s.", "B", r.GetHeaders().Get("A"))
	}
}
