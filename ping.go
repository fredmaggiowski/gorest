package gorest

import "net/http"

// Ping resource is defined here to avoid redefining it everywhere.
type Ping struct{}

// Get a ping ack simply returning http.StatusOK
func (p Ping) Get(r *http.Request) (int, Response) {
	return http.StatusOK, NewSimpleResponse(ACK, "")
}
