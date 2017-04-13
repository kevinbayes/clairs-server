package service

import (
	"../model"
	"../repository"
)

type RegistryService struct { }

func (s *RegistryService) ReadRegistry(id int64) (*model.Registry, error) {

	return repository.InstanceRegistryRepository().FindOne(id)
}
