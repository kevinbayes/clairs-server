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
)

const DEFAULT_SHIELD = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"150\" height=\"20\"><g shape-rendering=\"crispEdges\"><rect width=\"37\" height=\"20\" fill=\"#555\"/><rect x=\"37\" width=\"113\" height=\"20\" fill=\"#4c1\"/></g><g fill=\"#fff\" text-anchor=\"middle\" font-family=\"DejaVu Sans,Verdana,Geneva,sans-serif\" font-size=\"11\"><text x=\"18\" y=\"14\">clair</text><text x=\"92\" y=\"14\">not implemented</text></g></svg>"

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

	go func() {

		_registryService := &RegistryService{}

		for {

			_container := <-newContainerChannel // read from a channel

			log.Printf("Created container request %s.", _container.Image)

			registry, err := _registryService.ReadRegistry(_container.Registry)

			if(err != nil) {

				log.Panicf("Error reading registry: %s", err.Error())
			} else {

				gateway.DockerClientInstance().PullImage(registry, _container);
				analyzeContainerChannel <- _container
			}
		}
	}()

	go func() {

		_registryService := &RegistryService{}

		for {

			container := <-analyzeContainerChannel // read from a channel

			log.Printf("Created container request %s.", container.Image)

			_, err := _registryService.ReadRegistry(container.Registry)

			if(err != nil) {

				log.Panicf("Error reading registry: %s", err.Error())
			} else {

				imageId, parentId, err := gateway.DockerClientInstance().ImageId(container)

				if(err != nil) {

					log.Panicln(err)
				} else {

					err = gateway.ClairClientInstance().PostLayer(imageId, parentId)

					if(err != nil) {

						log.Panic(err)
					}
				}
			}
		}
	}()
}

func (s *ContainerService) CreateNewContainer(req *dto.NewContainer) (*model.Container, error) {

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

func (s *ContainerService) convertRequest(req *dto.NewContainer) (*model.Container) {

	return &model.Container{
		Registry: req.Registry,
		Image: req.Image,
		Shield: DEFAULT_SHIELD,
		State: model.STATE_REQUESTED,
	}
}

