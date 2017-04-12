package api

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
	middleware "../http"
)

func RegisterActuatorsHandlers(router *middleware.Middleware) {

	fmt.Printf("Registering party handlers")

	router.GET("/health", readHealthHandler)
}

func readHealthHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}
