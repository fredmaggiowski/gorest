package gorest

// Route defines a route pattern for a Resource.
type Route struct {
	resource Resource
	pattern  string
}

// NewRoute defines a New route object.
func NewRoute(resource Resource, pattern string) *Route {
	return &Route{
		resource: resource,
		pattern:  pattern,
	}
}

// GetPattern returns the Route URI pattern.
func (r *Route) GetPattern() string {
	return r.pattern
}

// GetResource returns the Resource to be used with the pattern.
func (r *Route) GetResource() Resource {
	return r.resource
}
