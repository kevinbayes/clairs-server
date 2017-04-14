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
package main

import (
	"net/http"
	middleware "./http"
	"./api"
	"./web"
	"./config"
	"./repository"
	"fmt"
)


func main() {

	_config := config.GetConfig()

	defer repository.Connect()

	_middleware := middleware.New();

	_middleware.RegisterFilter(middleware.NewSecurityFilter());
	_middleware.RegisterFilter(middleware.NewCorsFilter());

	api.RegisterActuatorsHandlers(_middleware)
	api.RegisterRegistriesHandlers(_middleware)
	api.RegisterContainersHandlers(_middleware)

	web.RegisterWebsiteHandlers(_middleware)


	http.ListenAndServe(fmt.Sprintf("%s:%s", _config.Server.Host, _config.Server.Port), _middleware.Router())
}
