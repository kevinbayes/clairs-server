package main

import (
	"net/http"
	middleware "./http"
	"github.com/julienschmidt/httprouter"
	"./api"
)


func main() {

	_middleware := middleware.New();

	_middleware.RegisterFilter(middleware.NewSecurityFilter());

	api.RegisterActuatorsHandlers(_middleware)

	_middleware.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{ \"results\" : [\"done\"]}"))
	})

	http.ListenAndServe(":18080", _middleware.Router())
}
