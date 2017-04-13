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
package repository

import (
	"fmt"
	"../model"
)

type RegistryRepository struct {
}


var registryRepositoryInstance *RegistryRepository = nil;

func InstanceRegistryRepository() *RegistryRepository {

	if(db == nil) {

		registryRepositoryInstance = &RegistryRepository{}
	}

	return registryRepositoryInstance
}

func (r *RegistryRepository) CreateRepository(repository *model.Registry) (error) {

	db, err := Connect()
	if(err != nil) {
		return err
	}

	// insert
	stmt, err := db.Prepare("INSERT registries SET name=?,description=?,uri=?,username=?,password=?,version=0")
	if(err != nil) {
		return err
	}

	res, err := stmt.Exec(repository.Name, repository.Description, repository.Uri, repository.Credentials.Username, repository.Credentials.Password)
	if(err != nil) {
		return err
	}

	id, err := res.LastInsertId()
	if(err != nil) {
		return err
	}

	repository.Id = id
	fmt.Printf("Created new id %d\n", id)

	return nil
}