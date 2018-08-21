package gorest

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// RestHandler defines a utility structure that provides REST handling functions.
type RestHandler struct {
	routes []*Route // List of all the available routes.
}

// NewHandler creates a new Handler instance.
func NewHandler() *RestHandler {
	return &RestHandler{}
}

// GetRoutes defines and returns all the handled Resource routes.
func (h *RestHandler) GetRoutes() []*Route {
	return h.routes
}

// RegisterRoute register provided route into the internal routes list.
func (h *RestHandler) RegisterRoute(route *Route) {
	h.routes = append(h.routes, route)
}

// SetRoutes returns all the handled Resource routes.
func (h *RestHandler) SetRoutes(routes []*Route) {
	h.routes = routes
}

// HandleRoute returns the handler function for a specific h
func (h *RestHandler) HandleRoute(route *Route) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		// Allow CORS on all APIs.
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Try to parse the request form data.
		if request.ParseForm() != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get handler function for specified resource for the route.
		handler := h.getHandlerFunction(request.Method, route.GetResource())
		if handler == nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Invoke the proper handler and retrieve the response and status code.
		code, response := handler(request)

		// TODO: Log a warning for invalid requests (40X - 50X)
		if code != http.StatusOK && code != http.StatusPermanentRedirect && code != http.StatusTemporaryRedirect {
		}

		var responseBody []byte
		var err error
		if response != nil {
			// Retrieve the body to be transmitted
			responseBody, err = response.GetBody()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// cache successful GET request via ETAG
			// with forced revalidation on each request.
			if request.Method == http.MethodGet && code == http.StatusOK {

				// Generate new ETAG, force cache revalidation and set the etag.
				etagBytes := sha256.Sum256(responseBody)
				etag := base64.StdEncoding.EncodeToString(etagBytes[:])
				w.Header().Set("Cache-Control", "private, max-age=0, must-revalidate")
				w.Header().Set("ETag", etag)

				// Check if request has an etag set and compare, return status code
				// 304 NOT MODIFIED if they match.
				if requestEtag := request.Header.Get("If-None-Match"); requestEtag == etag {
					w.WriteHeader(http.StatusNotModified)
					return
				}
			}

			// Verify if a cookie is needed and set it.
			if cookie := response.GetCookie(); cookie != nil {
				http.SetCookie(w, cookie)
			}
			// Verfiy if a set of headers is needed by the response
			// and if so set them all.
			if headers := response.GetHeaders(); headers != nil {
				for key, val := range headers {
					w.Header().Set(key, strings.Join(val, ", "))
				}
			}
		}
		// Write status code and data.
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(responseBody)
	}
}

// getHandlerFunction returns the actual http handler implementation based on
// resource type and request method.
func (h *RestHandler) getHandlerFunction(requestMethod string, r Resource) Handler {
	// TODO: Consider logging.
	// if h.verbose {
	// 	h.log.VerboseLog("[%s] Request for resource: %s.\n", requestMethod, reflect.TypeOf(resource))
	// }
	var handler Handler
	switch requestMethod {
	case http.MethodGet:
		if res, ok := r.(GetSupported); ok {
			handler = res.Get
		}
	case http.MethodPost:
		if res, ok := r.(PostSupported); ok {
			handler = res.Post
		}
	case http.MethodPut:
		if res, ok := r.(PutSupported); ok {
			handler = res.Put
		}
	case http.MethodDelete:
		if res, ok := r.(DeleteSupported); ok {
			handler = res.Delete
		}
	case http.MethodHead:
		if res, ok := r.(HeadSupported); ok {
			handler = res.Head
		}
	case http.MethodPatch:
		if res, ok := r.(PatchSupported); ok {
			handler = res.Patch
		}
	}

	return handler
}

// GetMuxRouter returns a Gorilla Mux router which implements all
// defined Routes.
func (h *RestHandler) GetMuxRouter(router *mux.Router) *mux.Router {
	if router == nil {
		router = mux.NewRouter().StrictSlash(true)
	}
	for _, route := range h.GetRoutes() {
		router.HandleFunc(route.GetPattern(), h.HandleRoute(route))
	}
	return router
}
