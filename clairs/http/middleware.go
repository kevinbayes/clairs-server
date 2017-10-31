// Copyright 2017 Kevin Bayes
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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

		if(h.IsApplicable("PUT", path)) {
			handler = h.Handler(handler)
		}
	}

	m.router.PUT(path, handler)
	return m
}


func (m *Middleware) DELETE(path string, handler httprouter.Handle) *Middleware {

	for _, h := range m.filterChain.filters {

		if(h.IsApplicable("DELETE", path)) {
			handler = h.Handler(handler)
		}
	}

	m.router.DELETE(path, handler)
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