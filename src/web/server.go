package web

import (
	"fmt"
	middleware "../http"
	"../config"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterWebsiteHandlers(router *middleware.Middleware) {

	fmt.Printf("Registering containers handlers")

	_config := config.GetConfig()

	//Containers
	router.Router().ServeFiles("/ui/*filepath", http.Dir(_config.Server.Filepath))
	router.GET("/", redirectHandler)
}

func redirectHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Location", "/ui")
	w.WriteHeader(http.StatusPermanentRedirect)
	w.Write([]byte(""))
}