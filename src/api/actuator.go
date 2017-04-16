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
	"time"
	"encoding/json"
	"log"
)

type Metrics struct {

	StartTime time.Time
}

var metrics = &Metrics{

	StartTime: time.Now(),
}

func RegisterActuatorsHandlers(router *middleware.Middleware) {

	log.Printf("Registering actuator handlers")

	router.GET("/health", readHealthHandler)
	router.GET("/metrics", readMetricsHandler)
}

func readHealthHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\":\"UP\"}"))
}


func readMetricsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	response, err := json.Marshal(metrics)

	if (err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}