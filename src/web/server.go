package web

import (
	middleware "../http"
	"../config"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
)

func RegisterWebsiteHandlers(router *middleware.Middleware) {

	log.Printf("Registering ui handlers")

	_config := config.GetConfig()

	//Containers
	router.Router().ServeFiles("/ui/*filepath", http.Dir(_config.Server.Filepath + "/public"))
	router.GET("/", redirectHandler)
}

func redirectHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Location", "/ui")
	w.WriteHeader(http.StatusPermanentRedirect)
	w.Write([]byte(""))
}