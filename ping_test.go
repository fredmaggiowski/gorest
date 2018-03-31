package gorest

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"
)

// TestPingGet verifies that the GET method for Ping resource works.
func TestPingGet(t *testing.T) {
	p := Ping{}
	code, resp := p.Get(nil)
	if code != http.StatusOK {
		t.Fatalf("Unexpected status code. Expected: %d - Found: %d.", http.StatusOK, code)
	}
	r, ok := resp.(SimpleResponse)
	if !ok {
		t.Fatalf("Unexected response type. Should be SimpleResponse but is: %s", reflect.TypeOf(resp).String())
	}
	pingBody, err := r.GetBody()
	if err != nil {
		t.Fatalf("Unexpected error extracting ping body: %s.", err.Error())
	}
	if bytes.Compare(pingBody, []byte("{\"status\":\"ACK\"}")) != 0 {
		t.Fatalf("Unexpected response body. Expected: %s - Found: %s.", []byte("{\"status\":\"ACK\"}"), string(pingBody))
	}
}
