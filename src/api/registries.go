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
	"fmt"
	"github.com/julienschmidt/httprouter"
	middleware "../http"
	"../service"
	"./dto"
	"../model"
	"encoding/json"
	"strconv"
)

func RegisterRegistriesHandlers(router *middleware.Middleware) {

	fmt.Printf("Registering registries handlers\n")

	//Registries
	router.POST("/api/registries", createRegistryHandler)
	router.GET("/api/registries", readRegistriesHandler)

	//Registry
	router.GET("/api/registries/:id", readRegistryHandler)
	router.PUT("/api/registries/:id", updateRegistryHandler)
	router.DELETE("/api/registries/:id", deleteRegistryHandler)

	//Registry Containers
	router.POST("/api/registries/:id/containers", createRegistryContainerHandler)
	router.GET("/api/registries/:id/containers", readRegistryContainersHandler)

	router.DELETE("/api/registries/:id/containers/:containerId", deleteRegistryContainerHandler)
}

func createRegistryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	decoder := json.NewDecoder(r.Body)

	var _body dto.NewRegistry
	err := decoder.Decode(&_body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	_service := &service.RegistryService{}

	if(r.URL.Query().Get("dryrun") == "true") {

		err := _service.TestNewRegistryCredentials(&_body)

		if(err != nil) {

			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {

			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte(""))
		}

	} else {

		res, err := _service.CreateRegistry(&_body)

		if (err != nil) {

			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {

			w.Header().Set("Location", fmt.Sprintf("/api/registries/%d", res.Id))
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(""))
		}
	}
}

func readRegistriesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	repo := &service.RegistryService{}

	model, err := repo.ReadRegistries()

	if (err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else if (model == nil) {

		http.NotFound(w, r)
	} else {

		_response := middleware.MakeSearchResult(len(model), 0, 0, model, make([]middleware.Link, 0))

		response, err := json.Marshal(_response)

		if (err != nil) {

			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {

			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	}
}

func readRegistryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	if(err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		repo := &service.RegistryService{}

		model, err := repo.ReadRegistry(id)

		if (err != nil) {

			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if (model == nil) {

			http.NotFound(w, r)
		} else {

			_response := middleware.MakeHateos(model, make([]middleware.Link, 0))

			response, err := json.Marshal(_response)

			if (err != nil) {

				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {

				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
		}
	}
}

func updateRegistryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	decoder := json.NewDecoder(r.Body)

	var _body model.Registry
	err := decoder.Decode(&_body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	_body.Id = id

	_service := &service.RegistryService{}

	if(r.URL.Query().Get("dryrun") == "true") {

		err := _service.TestRegistryCredentials(&_body)

		if(err != nil) {

			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {

			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte(""))
		}

	} else {

		_, err := _service.UpdateRegistry(&_body)

		if (err != nil) {

			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {

			_response := middleware.MakeHateos(_body, make([]middleware.Link, 0))

			response, err := json.Marshal(_response)

			if (err != nil) {

				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {

				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
		}
	}
}

func deleteRegistryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func createRegistryContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readRegistryContainersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func deleteRegistryContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}


