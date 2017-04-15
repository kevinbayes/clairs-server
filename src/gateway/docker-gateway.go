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
package gateway

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"context"
	"../model"
	"log"
)

type DockerClient struct { }

func DockerClientInstance() *DockerClient {

	return &DockerClient{}
}

/**
 * Validate that with the given credentials that the remote repository can be accessed.
 */
func (d *DockerClient) ValidateLogin(registry *model.Registry) (error) {

	client, err := client.NewEnvClient()

	if(err != nil) {

		return err
	}

	auth := types.AuthConfig{
		Username: registry.Credentials.Username,
		ServerAddress: registry.Uri,
		Password: registry.Credentials.Password,
	}


	res, err2 := client.RegistryLogin(context.Background(), auth)

	log.Printf("Received auth response status [%s].", res.Status)

	if(err2 != nil) {

		return err2
	}

	return nil
}