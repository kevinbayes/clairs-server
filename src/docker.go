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

	gateway.DockerClientInstance().SaveImage(container)

	err2 := gateway.ClairClientInstance().PostLayer("1f7dc3c631cf050d003954694badabe318cc5bbd86956d872e7da14e54039a10", parentId)

	layer, err3 := gateway.ClairClientInstance().GetLayer("1f7dc3c631cf050d003954694badabe318cc5bbd86956d872e7da14e54039a10");

	log.Print(err)
	log.Print(imageId)
	log.Print(parentId)
	log.Print(err2)
	log.Print(err3)
	log.Print(layer)
}
