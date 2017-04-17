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
	"fmt"
	"encoding/base64"
	"bytes"
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


/*
 * ValidateImage
 */
func (d *DockerClient) PullImage(registry *model.Registry, container *model.Container) (error) {

	client, err := client.NewEnvClient()

	if(err != nil) {

		return err
	}

	creds := fmt.Sprintf("%s:%s", registry.Credentials.Username, registry.Credentials.Password)

	sEnc := base64.StdEncoding.EncodeToString([]byte(creds))
	log.Print(sEnc)

	auth := types.ImagePullOptions{
	}

	log.Print(container.Image)

	read, err2 := client.ImagePull(context.Background(), container.Image, auth)

	if(err2 != nil) {

		return err2
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(read)
	newStr := buf.String()

	log.Print(newStr)

	return nil
}


func (d *DockerClient) ListImages() (error) {

	client, err := client.NewEnvClient()

	if(err != nil) {

		return err
	}

	auth := types.ImageListOptions{}

	read, err2 := client.ImageList(context.Background(), auth)

	if(err2 != nil) {

		return err2
	}

	for _, item := range read {

		log.Print("-------------------")
		log.Print(item.Labels)
		log.Print(item.RepoDigests)
		log.Print(item.RepoTags)
		log.Print(item.ID)
	}

	log.Print("---------------")
	log.Print(len(read))

	return nil
}

func (d *DockerClient) SearchImages() (error) {

	client, err := client.NewEnvClient()

	if(err != nil) {

		return err
	}

	auth := types.ImageListOptions{}

	read, err2 := client.ImageList(context.Background(), auth)

	if(err2 != nil) {

		return err2
	}

	for _, item := range read {

		log.Print("-------------------")
		log.Print(item.Labels)
		log.Print(item.RepoDigests)
		log.Print(item.RepoTags)
		log.Print(item.ID)
	}

	log.Print("---------------")
	log.Print(len(read))

	return nil
}