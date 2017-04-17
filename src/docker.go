package main

import (
	"./gateway"
	"./model"
	"log"
)

func main() {

	container := &model.Container{
		Image: "kibana:latest",
	}

	//gateway.DockerClientInstance().PullImage(registry, container)
	imageId, err := gateway.DockerClientInstance().ImageLayerId(container)

	log.Print(err)
	log.Print(imageId)
}
