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
package service

import (
	"log"
	"../model"
	"../repository"
)

var newContainerChannel = make(chan *model.ContainerImage)
var analyzeContainerChannel = make(chan *model.ContainerImage)

func init() {

	log.Print("Initiating Services")
	ContainerServiceSingleton().Init()
}


type GeneralService struct { }

var _generalService *GeneralService

func GeneralServiceSingleton() *GeneralService {

	if(_containerService == nil) {

		_generalService = &GeneralService{}
	}

	return _generalService;
}


func (g *GeneralService) GenerateSummary() (*model.Summary, error) {

	containerCount, _ := repository.InstanceContainerRepository().Count()
	registriesCount, _ := repository.InstanceRegistryRepository().Count()
	reportsCount, _ := repository.ImageReportRepositoryInstance().Count()


	return &model.Summary{
		Registries: model.RegistriesSummary {
			Total: registriesCount,
		},
		Containers: model.ContainersSummary {
			Total: containerCount,
		},
		Reports: model.ReportsSummary {
			Total: reportsCount,
		},
	}, nil
}