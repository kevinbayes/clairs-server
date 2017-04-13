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

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readRegistriesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readRegistryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	if(err != nil) {

		http.Error(w, err.Error(), 401)
	} else {

		repo := &service.RegistryService{}

		model, err := repo.ReadRegistry(id)

		if (err != nil) {

			http.Error(w, err.Error(), 401)
		} else if (model == nil) {

			http.NotFound(w, r)
		} else {

			response, err := json.Marshal(model)

			if (err != nil) {

				http.Error(w, err.Error(), 401)
			} else {

				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			}
		}
	}
}

func updateRegistryHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
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


