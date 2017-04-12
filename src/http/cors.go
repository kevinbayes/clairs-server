package http

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func NewCorsFilter() Filter {

	filter := Filter{
		Name: "Cors",
		Path: "/api/**",
		Methods: []string{"OPTIONS"},
		Order: 0,
		Handler: corsHandler,
	}

	return filter;
}


func corsHandler(next httprouter.Handle) httprouter.Handle {

	return func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

		if(request.Method == OPTIONS) {

			response.Header().Set("Content-Type", "application/json")
			response.Write([]byte{})
		} else {

			next(response, request, params)
		}
	};
}
