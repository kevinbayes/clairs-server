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

func (r *ContainerRepository) Save(container *model.ContainerImage) (error) {

	return inTransaction(func(tx *sql.Tx) (error) {

		// insert
		var lastInsertId int64 = 0
		err := db.QueryRow("INSERT INTO container_image(image, registry_id, state, created_on, updated_on, version) " +
			"VALUES ($1, $2, $3, $4, $4, 0) RETURNING id", container.Image, container.Registry,
			container.State, time.Now()).Scan(&lastInsertId)
		if (err != nil) {

			return err
		}

		container.Id = lastInsertId
		fmt.Printf("Created new id %d\n", lastInsertId)

		err = db.QueryRow("INSERT INTO container_image_tag(image_id, image_tag, state) " +
			"VALUES ($1, $2, $3) RETURNING id", lastInsertId, container.Tags[0].Tag, model.STATE_NOT_EVALUATED).Scan(&lastInsertId)
		if (err != nil) {

			return err
		}

		return nil
	});
}

func (r *ContainerRepository) Count() (int, error) {

	return count("container_image")
}

func (r *ContainerRepository) UpdateState(container *model.ContainerImage, newState string) (error) {

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

func (r *ContainerRepository) FindOne(_id int64) (*model.ContainerImage, error) {

	var (
		id int64
		registryId int64
		image string
		state string
		updated time.Time
		version int
		tag string
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

	_return := &model.ContainerImage{}

	if rows.Next() {

		err := rows.Scan(&id, &image, &registryId, &state, &updated, &version)
		if err != nil {

			log.Fatal(err)
			return nil, err
		}

		_return.Id = id
		_return.Image = image
		_return.Registry = registryId
		_return.State = state
		_return.DateUpdated = updated

		rows2, err := db.Query("select image_tag, state from container_image_tag where image_id = $1", _return.Id)
		defer rows2.Close()
		if(err != nil) {

			log.Fatal(err)
			return nil, err
		}

		for rows2.Next() {

			err := rows2.Scan(&tag, &state)
			if err != nil {

				log.Fatal(err)
				return nil, err
			}

			_return.Tags = append(_return.Tags, model.ContainerImageTag{
				Tag: tag,
				State: state,
			})
		}
	}

	return _return, nil
}


func (r *ContainerRepository) Find(pagination *Pagination) ([]*model.ContainerImage, error) {

	return r.find(func(db *sql.DB) (*sql.Rows, error) {

		return db.Query("select id, image, registry_id, state, updated_on, version from container_image " +
			"limit $1 offset $2", pagination.Size, pagination.Offset)
	})
}

func (r *ContainerRepository) FindByRegistry(pagination *Pagination, registryId int64) ([]*model.ContainerImage, error) {

	return r.find(func(db *sql.DB) (*sql.Rows, error) {

		return db.Query("select id, image, registry_id, state, " +
			"updated_on, version from container_image " +
			"where registry_id = $1 " +
			"limit $2 offset $3", registryId, pagination.Size, pagination.Offset)
	})
}

type sqlQuery func(*sql.DB) (*sql.Rows, error)

func (r *ContainerRepository) find(_func sqlQuery) ([]*model.ContainerImage, error) {

	var (
		id int64
		registryId int64
		image string
		state string
		updated time.Time
		version int
		tag string
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

	var result []*model.ContainerImage

	for rows.Next() {

		err := rows.Scan(&id, &image, &registryId, &state, &updated, &version)
		if err != nil {

			log.Fatal(err)
			return nil, err
		}

		_row := &model.ContainerImage{
			Id: id,
			Image: image,
			Registry: registryId,
			State: state,
			DateUpdated: updated,
		}

		rows2, err := db.Query("select image_tag, state from container_image_tag where image_id = $1", _row.Id)
		defer rows2.Close()
		if(err != nil) {

			log.Fatal(err)
			return nil, err
		}

		for rows2.Next() {

			err := rows2.Scan(&tag, &state)
			if err != nil {

				log.Fatal(err)
				return nil, err
			}

			_row.Tags = append(_row.Tags, model.ContainerImageTag{
				Tag: tag,
				State: state,
			})
		}

		result = append(result, _row)
	}

	return result, nil
}
