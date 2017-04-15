package service

import (
	"../gateway"
	"../model"
	"../repository"
	"../api/dto"
)

type RegistryService struct { }

func (s *RegistryService) TestRegistryCredentials(body *dto.NewRegistry) (error) {

	_registry := s.convertRequest(body)

	return s.testRegistryCredentials(_registry)
}

func (s *RegistryService) CreateRegistry(body *dto.NewRegistry) (*model.Registry, error) {

	_registry := s.convertRequest(body)

	validationError := s.testRegistryCredentials(_registry)

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

	return repository.InstanceRegistryRepository().FindOne(id)
}
