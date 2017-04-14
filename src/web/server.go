package web

import (
	"fmt"
	middleware "../http"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterWebsiteHandlers(router *middleware.Middleware) {

	fmt.Printf("Registering containers handlers")

	//Containers
	router.Router().ServeFiles("/ui/*filepath", http.Dir("src/public"))
	router.GET("/", redirectHandler)
}

func redirectHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Location", "/ui")
	w.WriteHeader(http.StatusPermanentRedirect)
	w.Write([]byte(""))
}