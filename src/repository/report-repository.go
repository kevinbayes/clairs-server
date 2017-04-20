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
	r.saveVulnerabilitySummary(report)

	return nil
}

func (r *ImageReportRepository) saveVulnerabilitySummary(report *model.ContainerImageReport) (error) {

	if( report.Counts != nil ) {

		return nil;
	}

	db, err := Connect()
	if(err != nil) {
		return err
	}

	for _, summary := range report.Counts {

		var lastInsertId int64 = 0
		err = db.QueryRow("INSERT INTO container_image_vulnerability_counts (vulnerability_level, count, image_report_id) " +
			"VALUES ($1, $2, $3) RETURNING id", summary.Level, summary.Count,
			report.Id).Scan(&lastInsertId)

		if (err != nil) {
			return err
		}
	}

	return nil
}