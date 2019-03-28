package gorest

import "net/http"

type testResponse struct {
	body    string
	bodyErr error
	cookie  *http.Cookie
	headers http.Header
}

func (t *testResponse) GetBody() ([]byte, error) {
	return []byte(t.body), t.bodyErr
}

func (t *testResponse) GetCookie() *http.Cookie {
	return t.cookie
}

func (t *testResponse) GetHeaders() http.Header {
	return t.headers
}
