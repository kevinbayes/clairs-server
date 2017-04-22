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
package gateway

import (
	"net/http"
	"../config"
	"../model"
	clairDto "github.com/coreos/clair/api/v1"
	"fmt"
	"bytes"
	"time"
	"log"
	"encoding/json"
	"io/ioutil"
)

var clairHttp = &http.Client{
	Timeout: time.Minute * 10,
}

type ClairClient struct { }

func ClairClientInstance() *ClairClient {

	return &ClairClient{}
}



func (c * ClairClient) AnalyzeImage(path string, container *model.Container, image string, layerIds []string) {

	size := len(layerIds)

	log.Printf("Analyzing %d layers... \n", size)

	for i := 0; i < size; i++ {

		log.Printf("Analyzing %s\n", layerIds[i])

		if i > 0 {
			analyzeLayer(layerIds[i], layerIds[i-1], path)
		} else {
			analyzeLayer(layerIds[i], "", path)
		}
	}
}

func analyzeLayer(layerId string, parentId string, path string) {

	_config := config.GetConfig()

	dto := clairDto.LayerEnvelope{
		Layer: &clairDto.Layer{
			Name: layerId,
			ParentName: parentId,
			Format: "Docker",
			Path: fmt.Sprintf("%s/%s/layer.tar", path, layerId),
		},
	}

	_req, _ := json.Marshal(dto)

	buf := bytes.NewBuffer(_req)

	response, err := clairHttp.Post(
		fmt.Sprintf("%s://%s:%s/%s",
			_config.Clair.Protocol,
			_config.Clair.Host,
			_config.Clair.Port,
			"v1/layers"),
		"application/json",
		buf)

	if err != nil {

		log.Printf("Error calling clair: %s", err.Error())
		return;
	}

	defer response.Body.Close()

	if response.StatusCode != 201 {
		body, _ := ioutil.ReadAll(response.Body)
		log.Printf("Got invalid response: %s", body)
	}

	log.Printf("Completed analyzing %s", layerId)
}


func (c * ClairClient) GetLayer(layerId string) (*clairDto.Layer, error) {

	_config := config.GetConfig().Clair

	_res, err := clairHttp.Get(
		fmt.Sprintf("%s://%s:%s/%s/%s?vulnerabilities",
			_config.Protocol,
			_config.Host,
			_config.Port,
			"v1/layers",
			layerId,
		))
	defer _res.Body.Close()

	if(err != nil) {
		return &clairDto.Layer{}, err;
	}

	var apiResponse clairDto.LayerEnvelope
	if err = json.NewDecoder(_res.Body).Decode(&apiResponse); err != nil {
		return &clairDto.Layer{}, err
	}

	return apiResponse.Layer, nil
}