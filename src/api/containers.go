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
)

func RegisterContainersHandlers(router *middleware.Middleware) {

	fmt.Printf("Registering containers handlers")

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

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readContainersHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
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


