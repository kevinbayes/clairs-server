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
	"../config"
	"log"
	"fmt"
	"encoding/base64"
	"bytes"
	"strings"
	"os/exec"
	"errors"
	"os"
	"io/ioutil"
	"encoding/json"
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
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(creds))

	auth := types.ImagePullOptions{
		RegistryAuth: encodedCredentials,
	}

	_, err2 := client.ImagePull(context.Background(), container.Image, auth)

	if(err2 != nil) {

		return err2
	}

	return nil
}


func (d *DockerClient) SaveImage(container *model.Container) (error) {

	path := fmt.Sprintf("%s/%d", config.GetConfig().TmpDir(), container.Id)

	mkerr := os.MkdirAll(path, os.ModePerm)
	if(mkerr != nil) {

		log.Panicln(mkerr)
		return mkerr
	}

	var stderr bytes.Buffer
	save := exec.Command("docker", "save", container.Image)

	save.Stderr = &stderr
	extract := exec.Command("tar", "xf", "-", "-C"+path)
	extract.Stderr = &stderr
	pipe, err := extract.StdinPipe()
	if err != nil {
		log.Printf("1. %s", err)
		return err
	}
	save.Stdout = pipe

	err = extract.Start()
	if err != nil {
		log.Printf("2. %s", err)
		return errors.New(stderr.String())
	}
	err = save.Run()
	if err != nil {
		log.Printf("3. %s", err)
		return errors.New(stderr.String())
	}
	err = pipe.Close()
	if err != nil {
		log.Printf("4. %s", err)
		return err
	}
	err = extract.Wait()
	if err != nil {
		log.Printf("5. %s", err)
		return errors.New(stderr.String())
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}


func (d *DockerClient) ImageLayers(container *model.Container) ([]string, error) {

	path := fmt.Sprintf("%s/%d", config.GetConfig().TmpDir(), container.Id)

	mf, err := os.Open(path + "/manifest.json")
	if err != nil {
		return nil, err
	}
	defer mf.Close()

	// https://github.com/docker/docker/blob/master/image/tarexport/tarexport.go#L17
	type manifestItem struct {
		Config   string
		RepoTags []string
		Layers   []string
	}

	var manifest []manifestItem
	if err = json.NewDecoder(mf).Decode(&manifest); err != nil {
		return nil, err
	} else if len(manifest) != 1 {
		return nil, err
	}
	var layers []string
	for _, layer := range manifest[0].Layers {
		layers = append(layers, strings.TrimSuffix(layer, "/layer.tar"))
	}
	return layers, nil
}



func (d *DockerClient) ImageId(container *model.Container) (string, error) {

	client, err := client.NewEnvClient()

	if (err != nil) {

		return "",err
	}

	inspect, _, err2 :=  client.ImageInspectWithRaw(context.Background(), container.Image)

	if(err2 != nil) {

		return "",err2
	}

	return inspect.ID[(strings.Index(inspect.ID, ":") + 1):len(inspect.ID)], err2
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