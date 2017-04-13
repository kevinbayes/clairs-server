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
	"log"
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

func (r *RegistryRepository) Save(registry *model.Registry) (error) {

	db, err := Connect()
	if(err != nil) {
		return err
	}

	// insert
	stmt, err := db.Prepare("INSERT registries SET name=?,description=?,uri=?,username=?,password=?,version=0")
	if(err != nil) {
		return err
	}

	res, err := stmt.Exec(registry.Name, registry.Description, registry.Uri, registry.Credentials.Username, registry.Credentials.Password)
	if(err != nil) {
		return err
	}

	id, err := res.LastInsertId()
	if(err != nil) {
		return err
	}

	registry.Id = id
	fmt.Printf("Created new id %d\n", id)

	return nil
}


func (r *RegistryRepository) FindOne(_id int64) (*model.Registry, error) {

	var (
		id int64
		name string
		description string
		uri string
		username string
		password string
		version int
	)

	db, err := Connect()
	if(err != nil) {

		return nil, err
	}

	// read one
	rows, err := db.Query("select id, name, description, uri, username, password, version from registries where id = $1", _id)
	if(err != nil) {

		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	if(rows.Next()) {

		err := rows.Scan(&id, &name, &description, &uri, &username, &password, &version)
		if err != nil {

			log.Fatal(err)
			return nil, err
		}

		_result := &model.Registry{
			Id: id,
			Name: name,
			Description: description,
			Uri: uri,
			Credentials: model.Credentials{
				Username: username,
				Password: password,
			},
		}

		return _result, nil
	}

	return nil, nil
}