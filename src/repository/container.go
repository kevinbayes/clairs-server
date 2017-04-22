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
	"database/sql"
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

	return notTransaction(func(db *sql.DB) (error) {

		// insert
		var lastInsertId int64 = 0
		err := db.QueryRow("INSERT INTO container_image(image, registry_id, state, created_on, updated_on, version) " +
			"VALUES ($1, $2, $3, $5, 0) RETURNING id", container.Image, container.Registry,
			container.State, time.Now(), time.Now()).Scan(&lastInsertId)
		if (err != nil) {

			return err
		}

		container.Id = lastInsertId
		fmt.Printf("Created new id %d\n", lastInsertId)

		return nil
	});
}

func (r *ContainerRepository) UpdateState(container *model.Container, newState string) (error) {

	return notTransaction(func(db *sql.DB) (error) {

		stmt, err := db.Prepare("UPDATE container_image SET state = $1, updated_on = $2 WHERE id = $3")
		if err != nil {
			log.Print(err)
		}

		res, err := stmt.Exec(newState, time.Now(), container.Id)
		if err != nil {
			log.Print(err)
			return err
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Print(err)
			return err
		}
		log.Printf("affected = %d\n", rowCnt)

		container.State = newState

		return nil
	});
}

func (r *ContainerRepository) FindOne(_id int64) (*model.Container, error) {

	var (
		id int64
		registryId int64
		image string
		state string
		updated time.Time
		version int
	)

	db, err := Connect()
	if(err != nil) {

		return nil, err
	}

	// read one
	rows, err := db.Query("select id, image, registry_id, state, updated_on, version from container_image where id = $1", _id)
	if(err != nil) {

		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&id, &image, &registryId, &state, &updated, &version)
		if err != nil {

			log.Fatal(err)
			return nil, err
		}

		return &model.Container{
			Id: id,
			Image: image,
			Registry: registryId,
			State: state,
			DateUpdated: updated,
		}, nil
	}

	return nil, nil
}


func (r *ContainerRepository) Find() ([]*model.Container, error) {

	return r.find(func(db *sql.DB) (*sql.Rows, error) {

		return db.Query("select id, image, registry_id, state, updated_on, version from container_image")
	})
}

func (r *ContainerRepository) FindByRegistry(registryId int64) ([]*model.Container, error) {

	return r.find(func(db *sql.DB) (*sql.Rows, error) {

		return db.Query("select id, image, registry_id, state, updated_on, version from container_image where registry_id = $1", registryId)
	})
}

type sqlQuery func(*sql.DB) (*sql.Rows, error)

func (r *ContainerRepository) find(_func sqlQuery) ([]*model.Container, error) {

	var (
		id int64
		registryId int64
		image string
		state string
		updated time.Time
		version int
	)

	db, err := Connect()
	if(err != nil) {

		return nil, err
	}

	// read one
	rows, err := _func(db)
	if(err != nil) {

		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var result []*model.Container

	for rows.Next() {

		err := rows.Scan(&id, &image, &registryId, &state, &updated, &version)
		if err != nil {

			log.Fatal(err)
			return nil, err
		}

		_row := &model.Container{
			Id: id,
			Image: image,
			Registry: registryId,
			State: state,
			DateUpdated: updated,
		}

		result = append(result, _row)
	}

	return result, nil
}
