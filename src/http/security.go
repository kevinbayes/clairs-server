package http

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func NewSecurityFilter() Filter {

	filter := Filter{
		Name: "Security",
		Path: "/**",
		Order: 0,
		Handler: securityHandler,
	}

	return filter;
}


func securityHandler(next httprouter.Handle) httprouter.Handle {

	return func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

		//TODO: Security
		//response.Header().Set("Content-Type", "application/json")
		//response.Write([]byte("{ \"results\" : [\"secure\"]}"))

		next(response, request, params)
	};
}
