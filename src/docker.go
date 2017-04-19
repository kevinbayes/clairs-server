package main

import (
	"./gateway"
	"./model"
	"log"
)

func main() {

	container := &model.Container{
		Id: 1,
		Image: "kibana:latest",
	}

	//gateway.DockerClientInstance().PullImage(registry, container)


	//imageId, parentId, err := gateway.DockerClientInstance().ImageId(container)

	err4 := gateway.DockerClientInstance().SaveImage(container)


	if(err4 != nil) {

		log.Print(err4)
		panic(err4.Error())
	}

	layers, err5 := gateway.DockerClientInstance().ImageLayers(container)

	if(err5 != nil) {

		log.Print(err5)
		panic(err5.Error())
	}

	log.Print(layers)

	//log.Printf(path)

	//gateway.ClairClientInstance().PostLayer("1f7dc3c631cf050d003954694badabe318cc5bbd86956d872e7da14e54039a10")

	layer, err3 := gateway.ClairClientInstance().GetLayer("1f7dc3c631cf050d003954694badabe318cc5bbd86956d872e7da14e54039a10");

	/*log.Print(err)
	log.Print(imageId)
	log.Print(parentId)*/
	//log.Print(err2)
	log.Print(err3)
	log.Print(layer)
}
