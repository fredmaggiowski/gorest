package gorest

import (
	"net/http"
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

// TestGetHandler verifies that the internal getHandler method works by
// properly converting and retrieving the correct handler.
func TestGetHandler(t *testing.T) {
	type invalidResource struct{}

	GetFunctionName := func(i interface{}) string {
		return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	}

	h := NewHandler()
	// Get
	handler := h.getHandler(http.MethodGet, testResourceWithGet{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Get") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Get", fn)
	}
	handler = h.getHandler(http.MethodGet, invalidResource{})
	if handler != nil {
		t.Fatalf("Unexpected non-nil handler: %+v.", handler)
	}

	handler = h.getHandler(http.MethodGet, testResourceWithPost{})
	if handler != nil {
		t.Fatalf("Unexpected non-nil handler: %+v.", handler)
	}
	// Post
	handler = h.getHandler(http.MethodPost, testResourceWithPost{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Post") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Post", fn)
	}
	// PUT
	handler = h.getHandler(http.MethodPut, testResourceWithPut{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Put") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Put", fn)
	}
	// DELETE
	handler = h.getHandler(http.MethodDelete, testResourceWithDelete{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Delete") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Delete", fn)
	}
	// HEAD
	handler = h.getHandler(http.MethodHead, testResourceWithHead{})
	if handler == nil {
		t.Fatalf("Unexpected nil handler.")
	}
	if fn := GetFunctionName(handler); !strings.ContainsAny(fn, "Head") {
		t.Fatalf("Unexpected function name. Expected: %s - Found: %s.", "Head", fn)
	}
	// PATCH
	handler = h.getHandler(http.MethodPatch, testResourceWithPatch{})
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
