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
package api

import (
	"net/http"
	"encoding/json"
	"../repository"
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

func listRespond(entity interface{}, total int, pagination *repository.Pagination, err error, w http.ResponseWriter, r *http.Request) {

	if (err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if (entity == nil) {

		http.NotFound(w, r)
	} else {

		listOk(entity, pagination.Size, total / pagination.Size, pagination.Page, w)
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


func parseJson(_body *interface{}, r *http.Request) {

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(_body)
	if err != nil {
		panic(err)
	}
}
