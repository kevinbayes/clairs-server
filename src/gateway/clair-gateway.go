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
	"sync"
)

var clairHttp = &http.Client{
	Timeout: time.Minute * 10,
}

type ClairClient struct { }

func ClairClientInstance() *ClairClient {

	return &ClairClient{}
}

func (c * ClairClient) PostLayers(container *model.Container) (error) {

	//path, err := DockerClientInstance().SaveImage(container)

	dto := clairDto.LayerEnvelope{
		Layer: &clairDto.Layer{
			Name: container.Image,
			ParentName: container.Image,
			Format: "Docker",
			Path: "file:///Users/kevinbayes/software/clair/kibana.tar",
		},
	}

	_req, err := json.Marshal(dto)

	buf := bytes.NewBuffer(_req)
	_config := config.GetConfig().Clair

	log.Printf("%s://%s:%s/%s",
		_config.Protocol,
		_config.Host,
		_config.Port,
		"v1/layers")

	res, err := clairHttp.Post(
		fmt.Sprintf("%s://%s:%s/%s",
			_config.Protocol,
			_config.Host,
			_config.Port,
			"v1/layers"),
		"application/json",
		buf)

	if(err != nil) {

		return err
	}

	log.Print(res)

	return nil
}

func analyzeLayer(wg *sync.WaitGroup, layerId string, parentId string, path string, reportChannel chan<- error) {

	defer wg.Done()

	dto := clairDto.LayerEnvelope{
		Layer: &clairDto.Layer{
			Name: layerId,
			ParentName: parentId,
			Format: "Docker",
			Path: "file:///Users/kevinbayes/software/clair/kibana.tar",
		},
	}

	_req, err := json.Marshal(dto)

	buf := bytes.NewBuffer(_req)
	_config := config.GetConfig().Clair

	_, err = clairHttp.Post(
		fmt.Sprintf("%s://%s:%s/%s",
			_config.Protocol,
			_config.Host,
			_config.Port,
			"v1/layers"),
		"application/json",
		buf)

	reportChannel <- err
}

func (c * ClairClient) PostLayer(container *model.Container, image string, layerIds []string) {

	var wg sync.WaitGroup

	path := fmt.Sprintf("%s/%d", config.GetConfig().TmpDir(), container.Id)
	size := len(layerIds)

	reportChannel := make(chan error)

	log.Printf("Analyzing %d layers... \n", size)

	for i := 0; i < size; i++ {

		log.Printf("Analyzing %s\n", layerIds[i])
		wg.Add(1)

		if i > 0 {
			go analyzeLayer(&wg, layerIds[i], layerIds[i-1], path, reportChannel)
		} else {
			go analyzeLayer(&wg, layerIds[i], image, path, reportChannel)
		}
	}

	for i := 0; i < size; i++ {

		err := <- reportChannel
		println(err)
	}
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