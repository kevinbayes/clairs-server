package http

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

type Middleware struct {
	router *httprouter.Router
	filterChain FilterChain
}

func New() *Middleware {

	_middleware := &Middleware{
		router: httprouter.New(),
		filterChain: FilterChain{
			filters: []Filter{},
		},
	}

	return _middleware
}

func (m *Middleware) RegisterFilter(filter Filter) *Middleware {

	m.filterChain.addFilter(filter)
	return m
}

func (m *Middleware) GET(path string, handler httprouter.Handle) *Middleware {

	for _, h := range m.filterChain.filters {

		if(h.IsApplicable("GET", path)) {
			handler = h.Handler(handler)
		}
	}

	m.router.GET(path, handler)
	return m
}

func (m *Middleware) POST(path string, handler httprouter.Handle) *Middleware {

	for _, h := range m.filterChain.filters {

		if(h.IsApplicable("POST", path)) {
			handler = h.Handler(handler)
		}
	}

	m.router.POST(path, handler)
	return m
}

func (m *Middleware) PUT(path string, handler httprouter.Handle) *Middleware {

	for _, h := range m.filterChain.filters {

		if(h.IsApplicable("POST", path)) {
			handler = h.Handler(handler)
		}
	}

	m.router.POST(path, handler)
	return m
}


func (m *Middleware) DELETE(path string, handler httprouter.Handle) *Middleware {

	for _, h := range m.filterChain.filters {

		if(h.IsApplicable("POST", path)) {
			handler = h.Handler(handler)
		}
	}

	m.router.POST(path, handler)
	return m
}

func (m *Middleware) Router() *httprouter.Router {

	return m.router;
}


func notFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{ message: \"%s\" }", "Not found.");
}