package service

import (
	"../gateway"
	"../model"
	"../repository"
	"../api/dto"
)

type RegistryService struct { }

func (s *RegistryService) TestNewRegistryCredentials(body *dto.NewRegistry) (error) {

	_registry := s.convertRequest(body)

	return s.testRegistryCredentials(_registry)
}

func (s *RegistryService) TestRegistryCredentials(_registry *model.Registry) (error) {

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

func (s *RegistryService) UpdateRegistry(_registry *model.Registry) (*model.Registry, error) {

	validationError := s.testRegistryCredentials(_registry)

	if(validationError != nil) {

		return nil, validationError
	}

	err := repository.InstanceRegistryRepository().Update(_registry)

	return _registry, err
}

func (s *RegistryService) ReadRegistries() ([]*model.Registry, error) {

	return repository.InstanceRegistryRepository().Find()
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