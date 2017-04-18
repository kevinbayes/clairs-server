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
	imageId, parentId, err := gateway.DockerClientInstance().ImageId(container)

	err2 := gateway.ClairClientInstance().PostLayer("4b8f83b7783857502bb3139a4c3ec9cbaf6eb2b34e0ca7360c9745b9aa5dae0a", parentId)

	layer, err3 := gateway.ClairClientInstance().GetLayer("4b8f83b7783857502bb3139a4c3ec9cbaf6eb2b34e0ca7360c9745b9aa5dae0a");

	log.Print(err)
	log.Print(imageId)
	log.Print(parentId)
	log.Print(err2)
	log.Print(err3)
	log.Print(layer)
}
