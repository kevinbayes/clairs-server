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
	"../model"
	"log"
	"time"
	"database/sql"
)

type RegistryRepository struct {
}


var registryRepositoryInstance *RegistryRepository = nil;

func InstanceRegistryRepository() *RegistryRepository {

	if(registryRepositoryInstance == nil) {

		registryRepositoryInstance = &RegistryRepository{}
	}

	return registryRepositoryInstance
}

func (r *RegistryRepository) Save(registry *model.Registry) (error) {

	return notTransaction(func(db *sql.DB) (error) {
		// insert
		var lastInsertId int64 = 0
		err := db.QueryRow("INSERT INTO registries(name, description, uri, username, password, created_on, version) VALUES ($1, $2, $3, $4, $5, $6, 0) RETURNING id", registry.Name, registry.Description, registry.Uri, registry.Credentials.Username, registry.Credentials.Password, time.Now()).Scan(&lastInsertId)
		if (err != nil) {
			return err
		}

		registry.Id = lastInsertId
		log.Printf("Created new id %d\n", lastInsertId)
		return err
	});
}

func (r *RegistryRepository) Update(registry *model.Registry) (error) {

	return notTransaction(func(db *sql.DB) (error) {

		stmt, err := db.Prepare("UPDATE registries SET name=$1, description=$2, uri=$3, username=$4, password=$5, version= version + 1 WHERE id = $6")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(registry.Name, registry.Description, registry.Uri, registry.Credentials.Username, registry.Credentials.Password, registry.Id)
		if err != nil {
			return err
		}

		return nil
	});
}

func (r *RegistryRepository) Delete(registry *model.Registry) (error) {

	return notTransaction(func(db *sql.DB) (error) {

		stmt, err := db.Prepare("DELETE FROM registries WHERE id = $1")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(registry.Id)
		if err != nil {
			return err
		}

		return nil
	});
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


func (r *RegistryRepository) Find() ([]*model.Registry, error) {

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
	rows, err := db.Query("select id, name, description, uri, username, password, version from registries")
	if(err != nil) {

		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var result []*model.Registry

	for rows.Next() {

		err := rows.Scan(&id, &name, &description, &uri, &username, &password, &version)
		if err != nil {

			log.Fatal(err)
			return nil, err
		}

		_row := &model.Registry{
			Id: id,
			Name: name,
			Description: description,
			Uri: uri,
			Credentials: model.Credentials{
				Username: username,
				Password: password,
			},
		}

		result = append(result, _row)
	}

	return result, nil
}