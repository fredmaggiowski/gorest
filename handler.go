package gorest

import "net/http"

// Handler is the actual function that must be implemented by each resource.
type Handler func(*http.Request) (int, Response)
