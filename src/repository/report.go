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
	"time"
	"database/sql"
	"log"
)

type ImageReportRepository struct {
}

var imageReportRepositoryInstance *ImageReportRepository = nil;

func ImageReportRepositoryInstance() *ImageReportRepository {

	if(imageReportRepositoryInstance == nil) {

		imageReportRepositoryInstance = &ImageReportRepository{}
	}

	return imageReportRepositoryInstance
}

func (r *ImageReportRepository) Save(report *model.ContainerImageReport) (error) {

	db, err := Connect()
	if(err != nil) {
		return err
	}

	// insert
	var lastInsertId int64 = 0
	err = db.QueryRow("INSERT INTO container_image_report (image_id, layer_id, shield, created_on) " +
		"VALUES ($1, $2, $3, $4) RETURNING id", report.ImageId, report.Layer,
		report.Shield, time.Now()).Scan(&lastInsertId)

	if(err != nil) {
		return err
	}

	report.Id = lastInsertId

	fmt.Printf("Created new id %d\n", lastInsertId)
	r.saveVulnerabilitySummary(report, db)

	return nil
}

func (r *ImageReportRepository) saveVulnerabilitySummary(report *model.ContainerImageReport, db *sql.DB) (error) {

	if( report.Counts == nil ) {

		return nil;
	}

	for _, summary := range report.Counts {

		var lastInsertId int64 = 0
		err := db.QueryRow("INSERT INTO container_image_vulnerability_counts (vulnerability_level, count, image_report_id) " +
			"VALUES ($1, $2, $3) RETURNING id", summary.Level, summary.Count,
			report.Id).Scan(&lastInsertId)

		if (err != nil) {
			return err
		}
	}

	return nil
}

func (r *ImageReportRepository) Count() (int, error) {

	return count("container_image_report")
}

func (r *ImageReportRepository) RegistryCount(id int64) (int, error) {

	return count(fmt.Sprintf("container_image_report where image_id in ( select i.id from container_image i where i.registry_id = %d )", id))
}


func (r *ImageReportRepository) FindLatest(containerId int64, _tag string) (*model.ContainerImageReport, error) {

	var (
		id int64
		imageId int64
		tag string
		layerId string
		shield string
		createdOn time.Time
	)

	db, err := Connect()
	if(err != nil) {

		return nil, err
	}

	// read one
	rows, err := db.Query("select id, image_id, image_tag, layer_id, shield, created_on from container_image_report where image_id = $1 and image_tag = $2 order by created_on desc", containerId, _tag)
	if(err != nil) {

		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&id, &imageId, &tag, &layerId, &shield, &createdOn)
		if err != nil {

			log.Fatal(err)
			return nil, err
		}

		return &model.ContainerImageReport{
			Id: id,
			ImageId: imageId,
			Layer: layerId,
			Shield: shield,
			Tag: tag,
			CreatedOn: createdOn,
		}, nil
	}

	return nil, nil
}