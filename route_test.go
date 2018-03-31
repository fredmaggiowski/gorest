package gorest

import (
	"reflect"
	"testing"
)

// TestRouteGetPattern verifies the Route GetPattern behaviour.
func TestRouteGetPattern(t *testing.T) {
	r := NewRoute(nil, "/the/pattern")
	if r.GetPattern() != "/the/pattern" {
		t.Fatalf("Unexpected patter found. Expected: %s - Found: %s.", "/the/pattern", r.GetPattern())
	}
}

// TestRouteGetResource verifies the Route GetResource behaviour.
func TestRouteGetResource(t *testing.T) {
	type testResource struct {
		A string
	}
	resource := testResource{"a_value"}
	r := NewRoute(resource, "")

	got := r.GetResource()
	if reflect.TypeOf(got) != reflect.TypeOf(resource) {
		t.Fatalf("Unexpected resource type found. Expected: %s - Found: %s.", reflect.TypeOf(resource), reflect.TypeOf(got))
	}

	convertedResource, ok := got.(testResource)
	if !ok {
		t.Fatalf("Failed type conversion to testResource. Type is: %s.", reflect.TypeOf(got).String())
	}
	if convertedResource.A != "a_value" {
		t.Fatalf("Unexpected resource content. Expected: %s - Found: %s.", "a_value", convertedResource.A)
	}
}
