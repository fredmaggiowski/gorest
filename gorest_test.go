package gorest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// TestRegisterRoute verifies that registering a single route works.
func TestRegisterRoute(t *testing.T) {
	h := NewHandler()

	type testResource struct{ A string }
	r := NewRoute(testResource{"X"}, "/the/pattern")

	h.RegisterRoute(r)
	if len(h.routes) != 1 {
		t.Fatalf("Unexpected routes len. Expected: %d - Found: %d.", 1, len(h.routes))
	}
	if h.routes[0].GetPattern() != "/the/pattern" {
		t.Fatalf("Unexpected route[0] pattern. Expected: %s - Found: %s.", "/the/pattern", h.routes[0].GetPattern())
	}
	resource := h.routes[0].GetResource()
	converted, ok := resource.(testResource)
	if !ok {
		t.Fatalf("Failed route[0] resource type conversion. Resource type is: %s.", reflect.TypeOf(resource).String())
	}
	if converted.A != "X" {
		t.Fatalf("Unexpected route[0] resource content. Expected: %s - Found: %s.", "X", converted.A)
	}

	// Register another route.
	r2 := NewRoute(testResource{"Y"}, "/the/second/pattern")

	h.RegisterRoute(r2)
	// Route[0]
	if len(h.routes) != 2 {
		t.Fatalf("Unexpected routes len. Expected: %d - Found: %d.", 2, len(h.routes))
	}
	if h.routes[0].GetPattern() != "/the/pattern" {
		t.Fatalf("Unexpected route[0] pattern. Expected: %s - Found: %s.", "/the/pattern", h.routes[0].GetPattern())
	}
	resource = h.routes[0].GetResource()
	converted, ok = resource.(testResource)
	if !ok {
		t.Fatalf("Failed route[0] resource type conversion. Resource type is: %s.", reflect.TypeOf(resource).String())
	}
	if converted.A != "X" {
		t.Fatalf("Unexpected route[0] resource content. Expected: %s - Found: %s.", "X", converted.A)
	}
	// Route[1]
	if h.routes[1].GetPattern() != "/the/second/pattern" {
		t.Fatalf("Unexpected route[1] pattern. Expected: %s - Found: %s.", "/the/pattern", h.routes[1].GetPattern())
	}
	resource = h.routes[1].GetResource()
	converted, ok = resource.(testResource)
	if !ok {
		t.Fatalf("Failed route[1] resource type conversion. Resource type is: %s.", reflect.TypeOf(resource).String())
	}
	if converted.A != "Y" {
		t.Fatalf("Unexpected route[1] resource content. Expected: %s - Found: %s.", "X", converted.A)
	}
}

// TestSetRoutes verifies that registering a slice of routes works.
func TestSetRoutes(t *testing.T) {
	h := NewHandler()

	type testResource struct{ A string }
	routes := []*Route{
		NewRoute(testResource{"A"}, "/the/1"),
		NewRoute(testResource{"B"}, "/the/2"),
		NewRoute(testResource{"C"}, "/the/3"),
	}

	h.SetRoutes(routes)
	if len(h.routes) != 3 {
		t.Fatalf("Unexpected routes len. Expected: %d - Found: %d.", 3, len(h.routes))
	}
	// Route[0].
	if h.routes[0].GetPattern() != "/the/1" {
		t.Fatalf("Unexpected route[0] pattern. Expected: %s - Found: %s.", "/the/1", h.routes[0].GetPattern())
	}
	resource := h.routes[0].GetResource()
	converted, ok := resource.(testResource)
	if !ok {
		t.Fatalf("Failed route[0] resource type conversion. Resource type is: %s.", reflect.TypeOf(resource).String())
	}
	if converted.A != "A" {
		t.Fatalf("Unexpected route[0] resource content. Expected: %s - Found: %s.", "A", converted.A)
	}
	// Route[1].
	if h.routes[1].GetPattern() != "/the/2" {
		t.Fatalf("Unexpected route[1] pattern. Expected: %s - Found: %s.", "/the/2", h.routes[1].GetPattern())
	}
	resource = h.routes[1].GetResource()
	converted, ok = resource.(testResource)
	if !ok {
		t.Fatalf("Failed route[1] resource type conversion. Resource type is: %s.", reflect.TypeOf(resource).String())
	}
	if converted.A != "B" {
		t.Fatalf("Unexpected route[1] resource content. Expected: %s - Found: %s.", "B", converted.A)
	}
	// Route[2].
	if h.routes[2].GetPattern() != "/the/3" {
		t.Fatalf("Unexpected route[2] pattern. Expected: %s - Found: %s.", "/the/3", h.routes[2].GetPattern())
	}
	resource = h.routes[2].GetResource()
	converted, ok = resource.(testResource)
	if !ok {
		t.Fatalf("Failed route[2] resource type conversion. Resource type is: %s.", reflect.TypeOf(resource).String())
	}
	if converted.A != "C" {
		t.Fatalf("Unexpected route[2] resource content. Expected: %s - Found: %s.", "C", converted.A)
	}
}

// TestGetRoutes verifies that the routes are properly returned when using the
// GetRoutes method.
func TestGetRoutes(t *testing.T) {
	h := NewHandler()

	type testResource struct{ A string }
	r := NewRoute(testResource{"X"}, "/the/pattern")
	r2 := NewRoute(testResource{"Y"}, "/the/second/pattern")

	h.RegisterRoute(r)
	h.RegisterRoute(r2)

	routes := h.GetRoutes()
	if len(routes) != 2 {
		t.Fatalf("Unexpected routes len. Expected: %d - Found: %d.", 2, len(routes))
	}
	if routes[0].GetPattern() != "/the/pattern" {
		t.Fatalf("Unexpected route[0] pattern. Expected: %s - Found: %s.", "/the/pattern", routes[0].GetPattern())
	}
	resource := routes[0].GetResource()
	converted, ok := resource.(testResource)
	if !ok {
		t.Fatalf("Failed route[0] resource type conversion. Resource type is: %s.", reflect.TypeOf(resource).String())
	}
	if converted.A != "X" {
		t.Fatalf("Unexpected route[0] resource content. Expected: %s - Found: %s.", "X", converted.A)
	}
	// Route[1]
	if routes[1].GetPattern() != "/the/second/pattern" {
		t.Fatalf("Unexpected route[1] pattern. Expected: %s - Found: %s.", "/the/pattern", routes[1].GetPattern())
	}
	resource = routes[1].GetResource()
	converted, ok = resource.(testResource)
	if !ok {
		t.Fatalf("Failed route[1] resource type conversion. Resource type is: %s.", reflect.TypeOf(resource).String())
	}
	if converted.A != "Y" {
		t.Fatalf("Unexpected route[1] resource content. Expected: %s - Found: %s.", "X", converted.A)
	}
}

// TestGetHandlerFunction verifies that the internal getHandlerFunction method
//  works by properly converting and retrieving the correct handler.
func TestGetHandlerFunction(t *testing.T) {
	type invalidResource struct{}

	GetFunctionName := func(i interface{}) string {
		return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	}

	h := NewHandler()
	// Get
	handler := h.getHandlerFunction(http.MethodGet, testResourceWithGet{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Get") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Get", fn)
	}
	handler = h.getHandlerFunction(http.MethodGet, invalidResource{})
	if handler != nil {
		t.Fatalf("Unexpected non-nil handler: %+v.", handler)
	}

	handler = h.getHandlerFunction(http.MethodGet, testResourceWithPost{})
	if handler != nil {
		t.Fatalf("Unexpected non-nil handler: %+v.", handler)
	}
	// Post
	handler = h.getHandlerFunction(http.MethodPost, testResourceWithPost{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Post") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Post", fn)
	}
	// PUT
	handler = h.getHandlerFunction(http.MethodPut, testResourceWithPut{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Put") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Put", fn)
	}
	// DELETE
	handler = h.getHandlerFunction(http.MethodDelete, testResourceWithDelete{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Delete") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Delete", fn)
	}
	// HEAD
	handler = h.getHandlerFunction(http.MethodHead, testResourceWithHead{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Head") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Head", fn)
	}
	// PATCH
	handler = h.getHandlerFunction(http.MethodPatch, testResourceWithPatch{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Patch") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Patch", fn)
	}
}

// TestGetMuxRouter verifies that a filled-in mux router is returned.
func TestGetMuxRouter(t *testing.T) {
	h := NewHandler()
	r := mux.NewRouter()

	router := h.GetMuxRouter(r)
	if router == nil {
		t.Fatalf("Unexpected nil router.")
	}

	router = h.GetMuxRouter(nil)
	if router == nil {
		t.Fatalf("Unexpected nil router.")
	}

	h.routes = []*Route{
		NewRoute(testResourceWithGet{}, "/hej"),
	}
	router = h.GetMuxRouter(nil)
	if router == nil {
		t.Fatalf("Unexpected nil router.")
	}
}

// TestHandleRouteFormParsingFail verifies that no unexpected panic happens
// when parsing the request form body.
func TestHandleRouteFormParsingFail(t *testing.T) {
	h := NewHandler()
	route := NewRoute(testResourceWithPost{}, "/")
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Body = nil // This will make the ParseForm() function fail.
	w := httptest.NewRecorder()

	handler := h.handleRoute(route)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("Unexpected status code. Expected: %d - Found. %d.", http.StatusBadRequest, w.Code)
	}
}

// TestHandleRouteFailsForUnsupportedMethod verifies the right status code is
// returned when an invalid method is required on a URI.
func TestHandleRouteFailsForUnsupportedMethod(t *testing.T) {
	h := NewHandler()
	route := NewRoute(testResourceWithGet{}, "/")
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	w := httptest.NewRecorder()

	handler := h.handleRoute(route)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("Unexpected status code. Expected: %d - Found. %d.", http.StatusMethodNotAllowed, w.Code)
	}
}

// TestHandleRouteVerifyFlowWithNilResponse the HandleRoute function has two
// flows, one of which is basically a shortpath running when the Resource
// returns a nil response; this tes validates that flow.
func TestHandleRouteVerifyFlowWithNilResponse(t *testing.T) {
	h := NewHandler()
	route := NewRoute(testResourceWithGet{}, "/")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler := h.handleRoute(route)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Unexpected status code. Expected: %d - Found. %d.", http.StatusOK, w.Code)
	}

	if w.Body.String() != "" {
		t.Fatalf("The body should be an empty string, found: `%s` instead", w.Body.String())
	}
}

// TestHandleRouteOKWithResponseError will invoke HandleRoute function and validate
// the processing flow when handle Response is not nil but has returned an error.
func TestHandleRouteOKWithResponseError(t *testing.T) {
	h := NewHandler()
	route := NewRoute(testResourceWithGetAndResponse{
		testResponse{
			body:    "",
			bodyErr: fmt.Errorf("errorbody"),
			cookie:  nil,
			headers: nil,
		},
	}, "/")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler := h.handleRoute(route)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("Unexpected status code. Expected: %d - Found. %d.", http.StatusInternalServerError, w.Code)
	}
}

// TestHandleRouteOKWithResponse will invoke HandleRoute function and validate
// the processing flow when handle Response is not nil.
func TestHandleRouteOKWithResponse(t *testing.T) {
	h := NewHandler()
	route := NewRoute(testResourceWithGetAndResponse{
		testResponse{
			body:    "testbody",
			bodyErr: nil,
			cookie:  nil,
			headers: nil,
		},
	}, "/")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler := h.handleRoute(route)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Unexpected status code. Expected: %d - Found. %d.", http.StatusOK, w.Code)
	}

	if w.Body.String() != "testbody" {
		t.Fatalf("Unexpected body. Expected: %s - Found: %s.", "testbody", w.Body.String())
	}
}

// TestHandlerRouteOKWithCachableResponse validates the If-None-Match header
// usage by returning StatusNotModified when Etag is present and is unchanged.
func TestHandlerRouteOKWithCachableResponse(t *testing.T) {
	h := NewHandler()
	route := NewRoute(testResourceWithGetAndResponse{
		testResponse{
			body:    "testbody",
			bodyErr: nil,
			cookie:  nil,
			headers: nil,
		},
	}, "/")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("If-None-Match", getETag([]byte("testbody")))
	w := httptest.NewRecorder()

	handler := h.handleRoute(route)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusNotModified {
		t.Fatalf("Unexpected status code. Expected: %d - Found. %d.", http.StatusNotModified, w.Code)
	}

	if w.Body.String() != "" {
		t.Fatalf("Unexpected body. Expected: '' - Found: %s.", w.Body.String())
	}
}

// TestHandleRouteOKWithResponseAndCookie will invoke HandleRoute function and validates
// that Response cookies are properly set.
func TestHandleRouteOKWithResponseAndCookie(t *testing.T) {
	h := NewHandler()
	route := NewRoute(testResourceWithGetAndResponse{
		testResponse{
			body:    "testbody",
			bodyErr: nil,
			cookie:  &http.Cookie{Domain: "monster.cookie.net", Name: "CookieMonster"},
			headers: nil,
		},
	}, "/")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler := h.handleRoute(route)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Unexpected status code. Expected: %d - Found. %d.", http.StatusOK, w.Code)
	}

	if w.Body.String() != "testbody" {
		t.Fatalf("Unexpected body. Expected: %s - Found: %s.", "testbody", w.Body.String())
	}
	cookies := w.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("Unexpected number of cookies. Expected: %d - Found: %d.", 1, len(cookies))
	}

	if cookies[0].Domain != "monster.cookie.net" {
		t.Fatalf("Unexpected cookie domain. Expected: %s - Found: %s.", "monster.cookie.net", cookies[0].Domain)
	}
	if cookies[0].Name != "CookieMonster" {
		t.Fatalf("Unexpected cookie name. Expected: %s - Found: %s.", "CookieMonster", cookies[0].Name)
	}
}

// TestHandleRouteOKWithResponseAndHeaders will invoke HandleRoute function and validates
// that Response headers are properly set.
func TestHandleRouteOKWithResponseAndHeaders(t *testing.T) {
	h := NewHandler()
	route := NewRoute(testResourceWithGetAndResponse{
		testResponse{
			body:    "testbody",
			bodyErr: nil,
			cookie:  nil,
			headers: map[string][]string{
				"Test-Header": []string{"my-value"},
			},
		},
	}, "/")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler := h.handleRoute(route)
	handler.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Unexpected status code. Expected: %d - Found. %d.", http.StatusOK, w.Code)
	}

	if w.Body.String() != "testbody" {
		t.Fatalf("Unexpected body. Expected: %s - Found: %s.", "testbody", w.Body.String())
	}

	headers := w.Result().Header
	if len(headers["Test-Header"]) != 1 {
		t.Fatalf("Unexpected Test-Header value. Expected slice with length 1. Found: %d.", len(headers["Test-Header"]))
	}
	if (headers["Test-Header"])[0] != "my-value" {
		t.Fatalf("Unexpected Test-Header value. Expected: %s - Found: %s.", "my-value", headers["Test-Header"][0])
	}
}
