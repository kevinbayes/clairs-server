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

import (
	"time"
)

type Container struct {
	Id int64
	Registry int64
	Image string
	State string
	Shield string

	DateUpdated time.Time
}

type ContainerImageReport struct {
	Id int64
	ImageId int64
	Layer string
	Shield string
	CreatedOn time.Time

	Counts []ContainerImageVulnerabilityCount
}


type ContainerImageVulnerabilityCount struct {
	Level string
	Count int
}