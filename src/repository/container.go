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
	"time"
)

type ContainerRepository struct {
}


var containerRepositoryInstance *ContainerRepository = nil;

func InstanceContainerRepository() *ContainerRepository {

	if(containerRepositoryInstance == nil) {

		containerRepositoryInstance = &ContainerRepository{}
	}

	return containerRepositoryInstance
}

func (r *ContainerRepository) Save(container *model.Container) (error) {

	db, err := Connect()
	if(err != nil) {
		return err
	}

	// insert
	var lastInsertId int64 = 0
	err = db.QueryRow("INSERT INTO container_image(image, registry_id, state, shield, created_on, version) " +
		"VALUES ($1, $2, $3, $4, $5, 0) RETURNING id", container.Image, container.Registry,
		container.State, container.Shield, time.Now()).Scan(&lastInsertId)
	if(err != nil) {
		return err
	}

	container.Id = lastInsertId
	fmt.Printf("Created new id %d\n", lastInsertId)

	return nil
}

func (r *ContainerRepository) FindOne(_id int64) (*model.Container, error) {

	var (
		id int64
		registryId int64
		image string
		state string
		shield string
		version int
	)

	db, err := Connect()
	if(err != nil) {

		return nil, err
	}

	// read one
	rows, err := db.Query("select id, image, registry_id, state, shield, version from container_image where id = $1", _id)
	if(err != nil) {

		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&id, &image, &registryId, &state, &shield, &version)
		if err != nil {

			log.Fatal(err)
			return nil, err
		}

		return &model.Container{
			Id: id,
			Image: image,
			Registry: registryId,
			State: state,
		}, nil
	}

	return nil, nil
}


func (r *ContainerRepository) Find() ([]*model.Container, error) {

	var (
		id int64
		registryId int64
		image string
		state string
		shield string
		version int
	)

	db, err := Connect()
	if(err != nil) {

		return nil, err
	}

	// read one
	rows, err := db.Query("select id, image, registry_id, state, shield, version from container_image")
	if(err != nil) {

		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var result []*model.Container

	for rows.Next() {

		err := rows.Scan(&id, &image, &registryId, &state, &shield, &version)
		if err != nil {

			log.Fatal(err)
			return nil, err
		}

		_row := &model.Container{
			Id: id,
			Image: image,
			Registry: registryId,
			State: state,
		}

		result = append(result, _row)
	}

	return result, nil
}