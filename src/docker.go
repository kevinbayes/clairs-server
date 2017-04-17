package main

import (
	"./gateway"
	"./model"
)

func main() {

	registry := &model.Registry{
		Credentials: model.Credentials{
			Username:"",
			Password:"",
		},
	}

	container := &model.Container{
		Image: "kibana:latest",
	}

	gateway.DockerClientInstance().PullImage(registry, container)
	gateway.DockerClientInstance().ListImages()

	//log.Print(err)
}
