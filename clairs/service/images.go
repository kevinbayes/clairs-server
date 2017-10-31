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
	"strings"
)

const DEFAULT_SHIELD = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"150\" height=\"20\"><g shape-rendering=\"crispEdges\"><rect width=\"37\" height=\"20\" fill=\"#555\"/><rect x=\"37\" width=\"113\" height=\"20\" fill=\"#f00\"/></g><g fill=\"#fff\" text-anchor=\"middle\" font-family=\"DejaVu Sans,Verdana,Geneva,sans-serif\" font-size=\"11\"><text x=\"18\" y=\"14\">clair</text><text x=\"92\" y=\"14\">not implemented</text></g></svg>"

type ImageService struct { }

var _imageService *ImageService

func ImageServiceSingleton() *ImageService {

	if(_imageService == nil) {

		_imageService = &ImageService{}
		_imageService.Init()
	}

	return _imageService;
}

func (s *ImageService) Init() {

	log.Print("Initializing ImageService")

	registryService := RegistryServiceSingleton();
	clairClient := gateway.ClairClientInstance();
	dockerClient := gateway.DockerClientInstance();

	go prepareImage(registryService, dockerClient);
	go runAnalysis(registryService, clairClient, dockerClient);
}

func prepareImage(_registryService *RegistryService, dockerClient *gateway.DockerClient) {

	for {
		_container := <-newContainerChannel // read from a channel

		err := prepareImageSync(_container, _registryService, dockerClient)

		if(err != nil) {

			log.Printf("Error reading registry: %s", err.Error())
		} else {

			analyzeContainerChannel <- _container
		}
	}
}

func prepareImageSync(_container *model.ContainerImage, _registryService *RegistryService, dockerClient *gateway.DockerClient) (error) {

	log.Printf("Created container request %s.", _container.Image)

	registry, err := _registryService.ReadRegistry(_container.Registry)

	if(err != nil) {

		log.Printf("Error reading registry: %s", err.Error())
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

		repository.InstanceImageRepository().UpdateState(container, newState)
	}
}

func runAnalysisSync(_container *model.ContainerImage, _registryService *RegistryService, clairClient *gateway.ClairClient, dockerClient *gateway.DockerClient) (*model.ContainerImageReport, error) {

	log.Printf("Analyse container image %s.", _container.Image)

	_, err := _registryService.ReadRegistry(_container.Registry)

	if(err != nil) {

		log.Printf("Error reading registry: %s", err.Error())
		return nil, err;
	}

	_, err = dockerClient.ImageId(_container)

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

	clairClient.AnalyzeImage(_container, layers)
	return saveAnalysisResults(_container, clairClient)
}

func saveAnalysisResults(image *model.ContainerImage, clairClient *gateway.ClairClient) (*model.ContainerImageReport, error) {

	/*layer*/_, err := clairClient.GetAnalysis(image)

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
				ImageId: image.Id,
				Shield: buf.String(),
			}

			repository.ImageReportRepositoryInstance().Save(report)
			return report, nil;
		}
	} else {

		total := 0
		_counts := make(map[string]int)

		/*for _, feature := range layer.Features {

			total += len(feature.Vulnerabilities)

			for _, vulnerability := range feature.Vulnerabilities {

				_counts[vulnerability.Severity]++
			}
		}*/

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
				ImageId: image.Id,
				Shield: buf.String(),
				Counts: _summary,
			}

			repository.ImageReportRepositoryInstance().Save(report)
			return report, nil;
		}
	}

}

func (s *ImageService) CreateNewImage(req *dto.NewContainer) (*model.ContainerImage, error) {

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

	err = repository.InstanceImageRepository().Save(_container)

	return _container, err
}

func (s *ImageService) ReadImages(pagination *repository.Pagination) ([]*model.ContainerImage, error) {

	return repository.InstanceImageRepository().Find(pagination)
}

func (s *ImageService) ReadContainersByRegistry(pagination *repository.Pagination, id int64) ([]*model.ContainerImage, error) {

	return repository.InstanceImageRepository().FindByRegistry(pagination, id)
}

func (s *ImageService) ReadContainer(id int64) (*model.ContainerImage, error) {

	return repository.InstanceImageRepository().FindOne(id)
}

func (s *ImageService) convertRequest(req *dto.NewContainer) (*model.ContainerImage) {

	image := req.Image
	tag := "latest"

	if(strings.Contains(image, ":")) {
		parts := strings.Split(image, ":")

		image = parts[0]
		tag = parts[1]
	}

	tags := []model.ContainerImageTag{
		{
			Tag: tag,
			State: model.STATE_NOT_EVALUATED,
		},
	}

	return &model.ContainerImage{
		Registry: req.Registry,
		Image: image,
		State: model.STATE_REQUESTED,
		Tags: tags,
	}
}


func (s *ImageService) EvaluateImages(id int64) (*model.ContainerImageReport, error) {

	container, err := repository.InstanceImageRepository().FindOne(id)
	if(container != nil) {

		registryService := RegistryServiceSingleton();
		clairClient := gateway.ClairClientInstance();
		dockerClient := gateway.DockerClientInstance();

		prepareImageSync(container, registryService, dockerClient)
		report, err := runAnalysisSync(container, registryService, clairClient, dockerClient)

		if(err != nil) {

			log.Printf("Error with analysis")
		} else if(container.State != "analyzed") {

			repository.InstanceImageRepository().UpdateState(container, "analyzed")
		}

		return report, err

	} else if(err != nil) {

		return nil, err;
	} else {

		return nil, errors.New("No container found.")
	}
}
