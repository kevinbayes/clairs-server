package api

import (
	"net/http"
	"encoding/json"
	middleware "../http"
)

type EntityApi interface {

	resolveLocation(entity interface{}) string
}

func created(location string, w http.ResponseWriter) {

	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(""))
}

func listRespond(entity interface{}, size int, pages int, page int, err error, w http.ResponseWriter, r *http.Request) {

	if (err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if (entity == nil) {

		http.NotFound(w, r)
	} else {

		listOk(entity, size, pages, page, w)
	}
}

func respond(entity interface{}, err error, w http.ResponseWriter, r *http.Request) {

	if (err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if (entity == nil) {

		http.NotFound(w, r)
	} else {

		ok(entity, w)
	}
}

func listOk(entity interface{}, size int, pages int, page int, w http.ResponseWriter) {

	_response := middleware.MakeSearchResult(size, pages, page, entity, make([]middleware.Link, 0))

	ok(_response, w)
}

func ok(entity interface{}, w http.ResponseWriter) {

	response, err := json.Marshal(entity)

	if (err != nil) {

		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func noContent(w http.ResponseWriter) {

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(""))
}

