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
	"../gateway"
	"../model"
	"../repository"
	"../api/dto"
)

type RegistryService struct { }

var _registryService *RegistryService

func RegistryServiceSingleton() *RegistryService {

	if(_registryService == nil) {

		_registryService = &RegistryService{}
	}

	return _registryService;
}

func (s *RegistryService) TestNewRegistryCredentials(body *dto.NewRegistry) (error) {

	_registry := s.convertRequest(body)

	return s.testRegistryCredentials(_registry)
}

func (s *RegistryService) TestRegistryCredentials(_registry *model.Registry) (error) {

	return s.testRegistryCredentials(_registry)
}

func (s *RegistryService) CreateRegistry(body *dto.NewRegistry) (*model.Registry, error) {

	_registry := s.convertRequest(body)

	var validationError error = nil

	if(len(_registry.Credentials.Username) > 0) {

		validationError = s.testRegistryCredentials(_registry)
	}

	if(validationError != nil) {

		return nil, validationError
	}

	err := repository.InstanceRegistryRepository().Save(_registry)

	return _registry, err
}




func (s *RegistryService) convertRequest(body *dto.NewRegistry) (*model.Registry) {

	return &model.Registry{
		Name: body.Name,
		Description: body.Description,
		Uri: body.Uri,
		Credentials: model.Credentials{
			Username: body.Credentials.Username,
			Password: body.Credentials.Password,
		},
	}
}

func (s *RegistryService) testRegistryCredentials(registry *model.Registry) (error) {

	validationError := gateway.DockerClientInstance().ValidateLogin(registry)

	return validationError
}

func (s *RegistryService) ReadRegistry(id int64) (*model.Registry, error) {

	registry, err := repository.InstanceRegistryRepository().FindOne(id)

	if(err == nil) {

		registry.Credentials.Password = ""
	}

	return registry, err
}

func (s *RegistryService) CompileRegistryDashboard(id int64) (*model.RegistryDashboard, error) {

	registry, err := repository.InstanceRegistryRepository().FindOne(id)

	if(err != nil) {

		return nil, err
	}

	registry.Credentials.Password = "";

	containerCount, _ := repository.InstanceContainerRepository().RegistryCount(registry.Id)
	reportsCount, _ := repository.ImageReportRepositoryInstance().RegistryCount(registry.Id)

	dashboard := &model.RegistryDashboard{
		Registry: *registry,
		Counts: model.Summary{
			Containers: model.ContainersSummary{
				Total: containerCount,
			},
			Reports: model.ReportsSummary{
				Total: reportsCount,
			},
		},
	}

	return dashboard, nil
}

func (s *RegistryService) UpdateRegistry(_registry *model.Registry) (*model.Registry, error) {

	validationError := s.testRegistryCredentials(_registry)

	if(validationError != nil) {

		return nil, validationError
	}

	err := repository.InstanceRegistryRepository().Update(_registry)

	return _registry, err
}

func (s *RegistryService) ReadRegistries(pagination *repository.Pagination) (*repository.PaginationResult, error) {

	return repository.InstanceRegistryRepository().Find(pagination)
}

func (s *RegistryService) ReadRegistriesWithContainerCount(pagination *repository.Pagination) (*repository.PaginationResult, error) {

	return repository.InstanceRegistryRepository().FindSummary(pagination)
}

func (s *RegistryService) DeleteRegistry(id int64) (*model.Registry, error) {

	reg, err := s.ReadRegistry(id)

	if(err != nil) {

		return nil, err
	}

	err = repository.InstanceRegistryRepository().Delete(reg)

	if(err != nil) {

		return nil, err
	}

	return reg, nil
}