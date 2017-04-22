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
	"strconv"
)

var containerService = service.ContainerServiceSingleton()

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
	router.PUT("/api/containers/:id/_evaluate", evaluateContainerReportsHandler)
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

		created(fmt.Sprintf("/api/containers/%d", res.Id), w)
	}
}

func readContainersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	model, err := containerService.ReadContainers()

	listRespond(model, len(model), 0, 0, err, w, r)
}

func readContainerHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	if(err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		model, err := containerService.ReadContainer(id)

		respond(model, err, w, r)
	}
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

func evaluateContainerReportsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	if(err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		model, err := containerService.EvaluateContainers(id)

		respond(model, err, w, r)
	}
}


