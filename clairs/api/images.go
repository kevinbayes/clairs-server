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

var imageService = service.ImageServiceSingleton()

func RegisterImagesHandlers(router *middleware.Middleware) {

	log.Printf("Registering container image handlers")

	//Containers
	router.POST("/api/images", createImageHandler)
	router.GET("/api/images", readImagesHandler)

	//Container
	router.GET("/api/images/:id", readImageHandler)
	router.PUT("/api/images/:id", updateImageHandler)
	router.DELETE("/api/images/:id", deleteImageHandler)

	//Container Reports
	router.GET("/api/images/:id/reports", readImageReportsHandler)

	//Container Report
	router.GET("/api/images/:id/reports/:reportId", readImageReportHandler)

	//Container Re-evaluate
	router.PUT("/api/images/:id/_evaluate", evaluateImageReportsHandler)
}

func createImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	decoder := json.NewDecoder(r.Body)

	var _body dto.NewContainer
	err := decoder.Decode(&_body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	res, err := imageService.CreateNewImage(&_body)

	if (err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		created(fmt.Sprintf("/api/containers/%d", res.Id), w)
	}
}

func readImagesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	pagination := middleware.MakePagination(r)

	model, err := imageService.ReadImages(pagination)

	listRespond(model, len(model), pagination, err, w, r)
}

func readImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	if(err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		model, err := imageService.ReadContainer(id)

		respond(model, err, w, r)
	}
}

func updateImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func deleteImageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readImageReportsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func readImageReportHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}

func evaluateImageReportsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	if(err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		model, err := imageService.EvaluateImages(id)

		respond(model, err, w, r)
	}
}


