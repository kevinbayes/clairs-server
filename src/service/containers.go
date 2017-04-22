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
package service

import (
	"../api/dto"
	"../model"
	"../repository"
	"../gateway"
	"errors"
	"log"
	"fmt"
	"os"
)

const DEFAULT_SHIELD = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"150\" height=\"20\"><g shape-rendering=\"crispEdges\"><rect width=\"37\" height=\"20\" fill=\"#555\"/><rect x=\"37\" width=\"113\" height=\"20\" fill=\"#f00\"/></g><g fill=\"#fff\" text-anchor=\"middle\" font-family=\"DejaVu Sans,Verdana,Geneva,sans-serif\" font-size=\"11\"><text x=\"18\" y=\"14\">clair</text><text x=\"92\" y=\"14\">not implemented</text></g></svg>"

type ContainerService struct { }

var _containerService *ContainerService

func ContainerServiceSingleton() *ContainerService {

	if(_containerService == nil) {

		_containerService = &ContainerService{}
		_containerService.Init()
	}

	return _containerService;
}

func (s *ContainerService) Init() {

	log.Print("Initializing ContainerService")

	registryService := RegistryServiceSingleton();
	clairClient := gateway.ClairClientInstance();
	dockerClient := gateway.DockerClientInstance();

	go prepareContainer(registryService, dockerClient);
	go runAnalysis(registryService, clairClient, dockerClient);
}

func prepareContainer(_registryService *RegistryService, dockerClient *gateway.DockerClient) {

	for {
		_container := <-newContainerChannel // read from a channel

		log.Printf("Created container request %s.", _container.Image)

		registry, err := _registryService.ReadRegistry(_container.Registry)

		if(err != nil) {

			log.Printf("Error reading registry: %s", err.Error())
		} else {

			err = dockerClient.PullImage(registry, _container);

			if(err != nil) {

				log.Printf("Error pulling container. Error: %s", err)
				return;
			}

			analyzeContainerChannel <- _container
		}
	}
}

func prepareContainerSync(_container *model.Container, _registryService *RegistryService, dockerClient *gateway.DockerClient) (error) {

	log.Printf("Created container request %s.", _container.Image)

	registry, err := _registryService.ReadRegistry(_container.Registry)

	if(err != nil) {

		log.Panicf("Error reading registry: %s", err.Error())
		return err;
	} else {

		return dockerClient.PullImage(registry, _container);
	}
}

func runAnalysis(_registryService *RegistryService, clairClient *gateway.ClairClient, dockerClient *gateway.DockerClient) {

	for {
		container := <-analyzeContainerChannel // read from a channel

		_, err := runAnalysisSync(container, _registryService, clairClient, dockerClient)

		newState := "analyzed";

		if(err != nil) {

			log.Printf("Error with analysis")
			newState = "invalid"
		}

		repository.InstanceContainerRepository().UpdateState(container, newState)
	}
}

func runAnalysisSync(_container *model.Container, _registryService *RegistryService, clairClient *gateway.ClairClient, dockerClient *gateway.DockerClient) (*model.ContainerImageReport, error) {

	log.Printf("Analyse container image %s.", _container.Image)

	_, err := _registryService.ReadRegistry(_container.Registry)

	if(err != nil) {

		log.Printf("Error reading registry: %s", err.Error())
		return nil, err;
	}

	imageId, err := dockerClient.ImageId(_container)

	if(err != nil) {

		log.Print(err)
		return nil, err;
	}

	path, err := dockerClient.SaveImage(_container)

	if(err != nil) {

		log.Print(err)
		return nil, err;
	}

	defer os.RemoveAll(path)

	layers, err := dockerClient.ImageLayers(_container)

	if(err != nil) {

		log.Print(err)
		return nil, err;
	}

	clairClient.AnalyzeImage(path, _container, imageId, layers)
	return saveAnalysisResults(_container, layers[len(layers)-1], clairClient)
}

func saveAnalysisResults(container *model.Container, layerId string, clairClient *gateway.ClairClient) (*model.ContainerImageReport, error) {

	layer, err := clairClient.GetLayer(layerId)

	if(err != nil) {

		shield := &model.Shield{
			Subject: model.Text{
				Value: "clair",
			},
			Status: model.Text{
				Value: "error",
			},
			Colour: "#f00",
			Template: "flat",
		}

		buf, err := ShieldsServiceSingleton().GenerateShieldSVG(shield)

		if( err != nil ) {

			log.Panic(err)
			return nil, err;
		} else {

			report := &model.ContainerImageReport{
				ImageId: container.Id,
				Layer: layerId,
				Shield: buf.String(),
			}

			repository.ImageReportRepositoryInstance().Save(report)
			return report, nil;
		}
	} else {

		total := 0
		_counts := make(map[string]int)

		for _, feature := range layer.Features {

			total += len(feature.Vulnerabilities)

			for _, vulnerability := range feature.Vulnerabilities {

				_counts[vulnerability.Severity]++
			}
		}

		shield := &model.Shield{
			Subject: model.Text{
				Value: "clair",
			},
			Status: model.Text{
				Value: fmt.Sprintf("%d vulnerabilities", total),
			},
			Colour: "#4c1",
			Template: "flat",
		}

		buf, err := ShieldsServiceSingleton().GenerateShieldSVG(shield)

		if( err != nil ) {

			log.Panic(err)
			return nil, err;
		} else {

			_summary := []model.ContainerImageVulnerabilityCount{}

			for key, value := range _counts {

				_summary = append(_summary, model.ContainerImageVulnerabilityCount{
					Level: key,
					Count: value,
				})
			}

			report := &model.ContainerImageReport{
				ImageId: container.Id,
				Layer: layerId,
				Shield: buf.String(),
				Counts: _summary,
			}

			repository.ImageReportRepositoryInstance().Save(report)
			return report, nil;
		}
	}

}

func (s *ContainerService) CreateNewContainer(req *dto.NewContainer) (*model.Container, error) {

	log.Print("Creatng new container.")

	_container := s.convertRequest(req)

	_registryService := &RegistryService{}

	registry, err := _registryService.ReadRegistry(req.Registry)
	if(err != nil) { return nil, err }

	if( registry == nil ) {

		return nil, errors.New("Not found")
	}

	log.Print("Sending request to pull container.")

	newContainerChannel <- _container

	log.Print("Sent request to pull container.")

	err = repository.InstanceContainerRepository().Save(_container)

	return _container, err
}

func (s *ContainerService) ReadContainers() ([]*model.Container, error) {

	return repository.InstanceContainerRepository().Find()
}

func (s *ContainerService) ReadContainer(id int64) (*model.Container, error) {

	return repository.InstanceContainerRepository().FindOne(id)
}

func (s *ContainerService) convertRequest(req *dto.NewContainer) (*model.Container) {

	return &model.Container{
		Registry: req.Registry,
		Image: req.Image,
		State: model.STATE_REQUESTED,
	}
}


func (s *ContainerService) EvaluateContainers(id int64) (*model.ContainerImageReport, error) {

	container, err := repository.InstanceContainerRepository().FindOne(id)
	if(container != nil) {

		registryService := RegistryServiceSingleton();
		clairClient := gateway.ClairClientInstance();
		dockerClient := gateway.DockerClientInstance();

		prepareContainerSync(container, registryService, dockerClient)
		report, err := runAnalysisSync(container, registryService, clairClient, dockerClient)

		if(err != nil) {

			log.Printf("Error with analysis")
		} else if(container.State != "analyzed") {

			repository.InstanceContainerRepository().UpdateState(container, "analyzed")
		}

		return report, err

	} else if(err != nil) {

		return nil, err;
	} else {

		return nil, errors.New("No container found.")
	}
}
