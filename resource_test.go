package gorest

import (
	"net/http"
	"reflect"
	"testing"
)

// TestResourceGetSupport verifies the type conversion for resource supporting
// GET method.
func TestResourceGetSupport(t *testing.T) {
	res, ok := (Resource(testResourceWithGet{})).(GetSupported)
	if !ok {
		t.Fatalf("Get should be supported by %s", reflect.TypeOf(testResourceWithGet{}).String())
	}
	code, resp := res.Get(nil)
	if code != 200 {
		t.Fatalf("Unexpected status code. Expected: %d - Found: %d.", 200, code)
	}
	if resp != nil {
		t.Fatalf("Unexpected resp. Shoud be `nil` - Found: %+v.", resp)
	}

	_, ok = (Resource(testResourceWithGet{})).(PostSupported)
	if ok {
		t.Fatalf("Post should not be supported by %s", reflect.TypeOf(testResourceWithGet{}).String())
	}
	_, ok = (Resource(testResourceWithGet{})).(PutSupported)
	if ok {
		t.Fatalf("Put should not be supported by %s", reflect.TypeOf(testResourceWithGet{}).String())
	}
	_, ok = (Resource(testResourceWithGet{})).(DeleteSupported)
	if ok {
		t.Fatalf("Delete should not be supported by %s", reflect.TypeOf(testResourceWithGet{}).String())
	}
	_, ok = (Resource(testResourceWithGet{})).(HeadSupported)
	if ok {
		t.Fatalf("Head should not be supported by %s", reflect.TypeOf(testResourceWithGet{}).String())
	}
	_, ok = (Resource(testResourceWithGet{})).(PatchSupported)
	if ok {
		t.Fatalf("Patch should not be supported by %s", reflect.TypeOf(testResourceWithGet{}).String())
	}
}

// TestResourcePostSupport verifies the type conversion for resource supporting
// POST method.
func TestResourcePostSupport(t *testing.T) {
	res, ok := (Resource(testResourceWithPost{})).(PostSupported)
	if !ok {
		t.Fatalf("Post should be supported by %s", reflect.TypeOf(testResourceWithPost{}).String())
	}
	code, resp := res.Post(nil)
	if code != 200 {
		t.Fatalf("Unexpected status code. Expected: %d - Found: %d.", 200, code)
	}
	if resp != nil {
		t.Fatalf("Unexpected resp. Shoud be `nil` - Found: %+v.", resp)
	}

	_, ok = (Resource(testResourceWithPost{})).(GetSupported)
	if ok {
		t.Fatalf("Get should not be supported by %s", reflect.TypeOf(testResourceWithPost{}).String())
	}
	_, ok = (Resource(testResourceWithPost{})).(PutSupported)
	if ok {
		t.Fatalf("Put should not be supported by %s", reflect.TypeOf(testResourceWithPost{}).String())
	}
	_, ok = (Resource(testResourceWithPost{})).(DeleteSupported)
	if ok {
		t.Fatalf("Delete should not be supported by %s", reflect.TypeOf(testResourceWithPost{}).String())
	}
	_, ok = (Resource(testResourceWithPost{})).(HeadSupported)
	if ok {
		t.Fatalf("Head should not be supported by %s", reflect.TypeOf(testResourceWithPost{}).String())
	}
	_, ok = (Resource(testResourceWithPost{})).(PatchSupported)
	if ok {
		t.Fatalf("Patch should not be supported by %s", reflect.TypeOf(testResourceWithPost{}).String())
	}
}

// TestResourcePutSupport verifies the type conversion for resource supporting
// PUT method.
func TestResourcePutSupport(t *testing.T) {
	res, ok := (Resource(testResourceWithPut{})).(PutSupported)
	if !ok {
		t.Fatalf("Put should be supported by %s", reflect.TypeOf(testResourceWithPut{}).String())
	}
	code, resp := res.Put(nil)
	if code != 200 {
		t.Fatalf("Unexpected status code. Expected: %d - Found: %d.", 200, code)
	}
	if resp != nil {
		t.Fatalf("Unexpected resp. Shoud be `nil` - Found: %+v.", resp)
	}

	_, ok = (Resource(testResourceWithPut{})).(GetSupported)
	if ok {
		t.Fatalf("Get should not be supported by %s", reflect.TypeOf(testResourceWithPut{}).String())
	}
	_, ok = (Resource(testResourceWithPut{})).(PostSupported)
	if ok {
		t.Fatalf("Post should not be supported by %s", reflect.TypeOf(testResourceWithPut{}).String())
	}
	_, ok = (Resource(testResourceWithPut{})).(DeleteSupported)
	if ok {
		t.Fatalf("Delete should not be supported by %s", reflect.TypeOf(testResourceWithPut{}).String())
	}
	_, ok = (Resource(testResourceWithPut{})).(HeadSupported)
	if ok {
		t.Fatalf("Head should not be supported by %s", reflect.TypeOf(testResourceWithPut{}).String())
	}
	_, ok = (Resource(testResourceWithPut{})).(PatchSupported)
	if ok {
		t.Fatalf("Patch should not be supported by %s", reflect.TypeOf(testResourceWithPut{}).String())
	}
}

// TestResourceDeleteSupport verifies the type conversion for resource supporting
// DELETE method.
func TestResourceDeleteSupport(t *testing.T) {
	res, ok := (Resource(testResourceWithDelete{})).(DeleteSupported)
	if !ok {
		t.Fatalf("Delete should be supported by %s", reflect.TypeOf(testResourceWithDelete{}).String())
	}
	code, resp := res.Delete(nil)
	if code != 200 {
		t.Fatalf("Unexpected status code. Expected: %d - Found: %d.", 200, code)
	}
	if resp != nil {
		t.Fatalf("Unexpected resp. Shoud be `nil` - Found: %+v.", resp)
	}

	_, ok = (Resource(testResourceWithDelete{})).(GetSupported)
	if ok {
		t.Fatalf("Get should not be supported by %s", reflect.TypeOf(testResourceWithDelete{}).String())
	}
	_, ok = (Resource(testResourceWithDelete{})).(PostSupported)
	if ok {
		t.Fatalf("Post should not be supported by %s", reflect.TypeOf(testResourceWithDelete{}).String())
	}
	_, ok = (Resource(testResourceWithDelete{})).(PutSupported)
	if ok {
		t.Fatalf("Put should not be supported by %s", reflect.TypeOf(testResourceWithDelete{}).String())
	}
	_, ok = (Resource(testResourceWithDelete{})).(HeadSupported)
	if ok {
		t.Fatalf("Head should not be supported by %s", reflect.TypeOf(testResourceWithDelete{}).String())
	}
	_, ok = (Resource(testResourceWithDelete{})).(PatchSupported)
	if ok {
		t.Fatalf("Patch should not be supported by %s", reflect.TypeOf(testResourceWithDelete{}).String())
	}
}

// TestResourceHeadSupport verifies the type conversion for resource supporting
// HEAD method.
func TestResourceHeadSupport(t *testing.T) {
	res, ok := (Resource(testResourceWithHead{})).(HeadSupported)
	if !ok {
		t.Fatalf("Head should be supported by %s", reflect.TypeOf(testResourceWithHead{}).String())
	}
	code, resp := res.Head(nil)
	if code != 200 {
		t.Fatalf("Unexpected status code. Expected: %d - Found: %d.", 200, code)
	}
	if resp != nil {
		t.Fatalf("Unexpected resp. Shoud be `nil` - Found: %+v.", resp)
	}

	_, ok = (Resource(testResourceWithHead{})).(GetSupported)
	if ok {
		t.Fatalf("Get should not be supported by %s", reflect.TypeOf(testResourceWithHead{}).String())
	}
	_, ok = (Resource(testResourceWithHead{})).(PostSupported)
	if ok {
		t.Fatalf("Post should not be supported by %s", reflect.TypeOf(testResourceWithHead{}).String())
	}
	_, ok = (Resource(testResourceWithHead{})).(PutSupported)
	if ok {
		t.Fatalf("Put should not be supported by %s", reflect.TypeOf(testResourceWithHead{}).String())
	}
	_, ok = (Resource(testResourceWithHead{})).(DeleteSupported)
	if ok {
		t.Fatalf("Delete should not be supported by %s", reflect.TypeOf(testResourceWithHead{}).String())
	}
	_, ok = (Resource(testResourceWithHead{})).(PatchSupported)
	if ok {
		t.Fatalf("Patch should not be supported by %s", reflect.TypeOf(testResourceWithHead{}).String())
	}
}

// TestResourcePatchSupport verifies the type conversion for resource supporting
// PATCH method.
func TestResourcePatchSupport(t *testing.T) {
	res, ok := (Resource(testResourceWithPatch{})).(PatchSupported)
	if !ok {
		t.Fatalf("Patch should be supported by %s", reflect.TypeOf(testResourceWithPatch{}).String())
	}
	code, resp := res.Patch(nil)
	if code != 200 {
		t.Fatalf("Unexpected status code. Expected: %d - Found: %d.", 200, code)
	}
	if resp != nil {
		t.Fatalf("Unexpected resp. Shoud be `nil` - Found: %+v.", resp)
	}

	_, ok = (Resource(testResourceWithPatch{})).(GetSupported)
	if ok {
		t.Fatalf("Get should not be supported by %s", reflect.TypeOf(testResourceWithPatch{}).String())
	}
	_, ok = (Resource(testResourceWithPatch{})).(PostSupported)
	if ok {
		t.Fatalf("Post should not be supported by %s", reflect.TypeOf(testResourceWithPatch{}).String())
	}
	_, ok = (Resource(testResourceWithPatch{})).(PutSupported)
	if ok {
		t.Fatalf("Put should not be supported by %s", reflect.TypeOf(testResourceWithPatch{}).String())
	}
	_, ok = (Resource(testResourceWithPatch{})).(DeleteSupported)
	if ok {
		t.Fatalf("Delete should not be supported by %s", reflect.TypeOf(testResourceWithPatch{}).String())
	}
	_, ok = (Resource(testResourceWithPatch{})).(HeadSupported)
	if ok {
		t.Fatalf("Head should not be supported by %s", reflect.TypeOf(testResourceWithPatch{}).String())
	}
}

//
// Structures for test
//

type testResourceWithGet struct{}

func (testResourceWithGet) Get(r *http.Request) (int, Response) {
	return 200, nil
}

type testResourceWithPost struct{}

func (testResourceWithPost) Post(r *http.Request) (int, Response) {
	return 200, nil
}

type testResourceWithPut struct{}

func (testResourceWithPut) Put(r *http.Request) (int, Response) {
	return 200, nil
}

type testResourceWithDelete struct{}

func (testResourceWithDelete) Delete(r *http.Request) (int, Response) {
	return 200, nil
}

type testResourceWithHead struct{}

func (testResourceWithHead) Head(r *http.Request) (int, Response) {
	return 200, nil
}

type testResourceWithPatch struct{}

func (testResourceWithPatch) Patch(r *http.Request) (int, Response) {
	return 200, nil
}
