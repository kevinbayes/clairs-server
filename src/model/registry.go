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
package model

type Registry struct {

	Id int64 `Id of the repository`
	Name string `Name of the repository`
	Description string `Description of the repository`
	Uri string `Uri to the repostory`
	ContainerCount int `Count of the containers registered.`

	Credentials Credentials
}

type RegistryDashboard struct {

	Registry Registry
	Counts Summary
}