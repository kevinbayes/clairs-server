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
	"../service"
	"log"
	"strconv"
)

func RegisterShieldsHandlers(router *middleware.Middleware) {

	log.Printf("Registering shields handlers")

	router.GET("/api/containers/:id/shield", readShieldHandler)
}

func readShieldHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)

	if(err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return;
	}

	buf, err := service.ShieldsServiceSingleton().GetShield(id)

	if(err != nil) {

		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write(buf.Bytes())
}