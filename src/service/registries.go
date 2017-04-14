package service

import (
	"../model"
	"../repository"
	"../api/dto"
)

type RegistryService struct { }

func (s *RegistryService) CreateRegistry(body *dto.NewRegistry) (*model.Registry, error) {

	_registry := &model.Registry{
		Name: body.Name,
		Description: body.Description,
		Uri: body.Uri,
		Credentials: model.Credentials{
			Username: body.Credentials.Username,
			Password: body.Credentials.Password,
		},
	}

	err := repository.InstanceRegistryRepository().Save(_registry)

	return _registry, err
}

func (s *RegistryService) ReadRegistry(id int64) (*model.Registry, error) {

	return repository.InstanceRegistryRepository().FindOne(id)
}
