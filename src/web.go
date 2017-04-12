package main

import (
	"net/http"
	middleware "./http"
	"./api"
)


func main() {

	_middleware := middleware.New();

	_middleware.RegisterFilter(middleware.NewSecurityFilter());
	_middleware.RegisterFilter(middleware.NewCorsFilter());

	api.RegisterActuatorsHandlers(_middleware)
	api.RegisterContainersHandlers(_middleware)

	http.ListenAndServe(":18080", _middleware.Router())
}
