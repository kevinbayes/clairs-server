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
	"log"
)

var registryService = service.RegistryServiceSingleton()

func RegisterRegistriesHandlers(router *middleware.Middleware) {

	log.Printf("Registering registries handlers")

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

	if(r.URL.Query().Get("dryrun") == "true") {

		err := registryService.TestNewRegistryCredentials(&_body)

		if(err != nil) {

			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {

			noContent(w)
		}

	} else {

		res, err := registryService.CreateRegistry(&_body)

		if (err != nil) {

			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {

			created(fmt.Sprintf("/api/registries/%d", res.Id), w)
		}
	}
}

func readRegistriesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	pagination := middleware.MakePagination(r)

	result, err := registryService.ReadRegistries(pagination)

	listRespond(result.Result, result.Total, pagination, err, w, r)
}

func readRegistryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	if(err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		model, err := registryService.ReadRegistry(id)

		respond(model, err, w, r)
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

	if(r.URL.Query().Get("dryrun") == "true") {

		err := registryService.TestRegistryCredentials(&_body)

		if(err != nil) {

			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {

			noContent(w)
		}

	} else {

		_, err := registryService.UpdateRegistry(&_body)

		respond(&_body, err, w, r)
	}
}

func deleteRegistryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	res, err := registryService.DeleteRegistry(id)

	respond(res, err, w, r)
}

func createRegistryContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(r.Body)

	var _body dto.NewContainer
	err = decoder.Decode(&_body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	_body.Registry = id

	res, err := containerService.CreateNewContainer(&_body)

	if (err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		created(fmt.Sprintf("/api/containers/%d", res.Id), w)
	}
}

func readRegistryContainersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	pagination := middleware.MakePagination(r)

	model, err := containerService.ReadContainersByRegistry(pagination, id)

	listRespond(model, len(model), pagination, err, w, r)

}

func deleteRegistryContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}


