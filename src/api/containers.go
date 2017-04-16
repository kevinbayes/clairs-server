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
	"github.com/julienschmidt/httprouter"
	middleware "../http"
	"log"
	"../service"
	"./dto"
	"encoding/json"
	"fmt"
)

var containerService = &service.ContainerService{}

func RegisterContainersHandlers(router *middleware.Middleware) {

	log.Printf("Registering containers handlers")

	//Containers
	router.POST("/api/containers", createContainerHandler)
	router.GET("/api/containers", readContainersHandler)

	//Container
	router.GET("/api/containers/:id", readContainerHandler)
	router.PUT("/api/containers/:id", updateContainerHandler)
	router.DELETE("/api/containers/:id", deleteContainerHandler)

	//Container Reports
	router.GET("/api/containers/:id/reports", readContainerReportsHandler)

	//Container Report
	router.GET("/api/containers/:id/reports/:reportId", readContainerReportHandler)

	//Container Re-evaluate
	router.PUT("/api/containers/:id/_evaluate", reevaluateContainerReportsHandler)
}

func createContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	decoder := json.NewDecoder(r.Body)

	var _body dto.NewContainer
	err := decoder.Decode(&_body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	res, err := containerService.CreateNewContainer(&_body)

	if (err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		w.Header().Set("Location", fmt.Sprintf("/api/containers/%d", res.Id))
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(""))
	}
}

func readContainersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	model, err := containerService.ReadContainers()

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

func readContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func updateContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func deleteContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readContainerReportsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readContainerReportHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func reevaluateContainerReportsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}


