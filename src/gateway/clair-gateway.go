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
)

var clairHttp = &http.Client{
	Timeout: time.Minute * 10,
}

type ClairClient struct { }

func ClairClientInstance() *ClairClient {

	return &ClairClient{}
}



func (c * ClairClient) AnalyzeImage(container *model.Container, image string, layerIds []string) {

	path := fmt.Sprintf("%s/%d", config.GetConfig().TmpDir(), container.Id)
	size := len(layerIds)

	reportChannel := make(chan error)
	defer close(reportChannel)

	log.Printf("Analyzing %d layers... \n", size)

	for i := 0; i < size; i++ {

		log.Printf("Analyzing %s\n", layerIds[i])

		if i > 0 {
			go analyzeLayer(layerIds[i], layerIds[i-1], path, reportChannel)
		} else {
			go analyzeLayer(layerIds[i], "", path, reportChannel)
		}
	}

	for i := 0; i < size; i++ {

		err := <- reportChannel
		if(err != nil) {
			println(err.Error())
		}
	}

}

func analyzeLayer(layerId string, parentId string, path string, reportChannel chan<- error) {

	_config := config.GetConfig()

	dto := clairDto.LayerEnvelope{
		Layer: &clairDto.Layer{
			Name: layerId,
			ParentName: parentId,
			Format: "Docker",
			Path: fmt.Sprintf("file://%s/%s/layer.tar", path, layerId),
		},
	}

	_req, err := json.Marshal(dto)

	if(err != nil) {
		reportChannel <- err
		return
	}

	buf := bytes.NewBuffer(_req)

	_, err = clairHttp.Post(
		fmt.Sprintf("%s://%s:%s/%s",
			_config.Clair.Protocol,
			_config.Clair.Host,
			_config.Clair.Port,
			"v1/layers"),
		"application/json",
		buf)

	log.Printf("Completed analyzing %s", layerId)

	reportChannel <- err
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